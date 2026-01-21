package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
)

const tsSize = 188

type PIDCtx struct {
	offset uint64
}

func main() {
	inFile := flag.String("in", "", "input ts")
	outFile := flag.String("out", "", "output ts")
	keyB64 := flag.String("key", "", "base64 aes key")
	vpid := flag.Int("vpid", -1, "video pid (hex or dec)")
	apid := flag.Int("apid", -1, "audio pid (hex or dec)")
	flag.Parse()

	if *inFile == "" || *outFile == "" || *keyB64 == "" {
		fmt.Println("missing required args")
		return
	}

	key, err := base64.StdEncoding.DecodeString(*keyB64)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	in, err := os.Open(*inFile)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out, err := os.Create(*outFile)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	ctxMap := map[uint16]*PIDCtx{}

	buf := make([]byte, tsSize)
	packetCount := 0

	for {
		_, err := io.ReadFull(in, buf)
		if err != nil {
			break
		}
		packetCount++

		if buf[0] != 0x47 {
			// 非法 TS，原样写
			out.Write(buf)
			continue
		}

		pusi := buf[1]&0x40 != 0
		pid := (uint16(buf[1]&0x1F) << 8) | uint16(buf[2])
		afc := (buf[3] >> 4) & 0x3

		payloadStart := 4
		if afc == 2 || afc == 0 {
			// 无 payload
			out.Write(buf)
			continue
		}
		if afc == 3 {
			// adaptation + payload
			adaptLen := int(buf[4])
			payloadStart += 1 + adaptLen
		}
		if payloadStart >= tsSize {
			out.Write(buf)
			continue
		}

		isTargetPID := int(pid) == *vpid || int(pid) == *apid
		if !isTargetPID {
			out.Write(buf)
			continue
		}

		ctx := ctxMap[pid]
		if ctx == nil {
			ctx = &PIDCtx{}
			ctxMap[pid] = ctx
		}

		payload := buf[payloadStart:]

		// PES 起始：counter reset + 跳过 PES header
		if pusi && len(payload) >= 9 && payload[0] == 0x00 && payload[1] == 0x00 && payload[2] == 0x01 {
			ctx.offset = 0

			pesHeaderLen := int(payload[8]) + 9
			if pesHeaderLen < len(payload) {
				decryptAESCTR(block, payload[pesHeaderLen:], ctx)
			}
		} else {
			decryptAESCTR(block, payload, ctx)
		}

		out.Write(buf)
	}

	fmt.Printf("done, packets: %d\n", packetCount)
}

func decryptAESCTR(block cipher.Block, data []byte, ctx *PIDCtx) {
	if len(data) == 0 {
		return
	}

	iv := make([]byte, 16)
	binary.BigEndian.PutUint64(iv[8:], ctx.offset/16)

	stream := cipher.NewCTR(block, iv)

	// 对齐 block
	if rem := ctx.offset % 16; rem != 0 {
		dummy := make([]byte, rem)
		stream.XORKeyStream(dummy, dummy)
	}

	stream.XORKeyStream(data, data)
	ctx.offset += uint64(len(data))
}
