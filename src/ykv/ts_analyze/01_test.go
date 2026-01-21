package ts_analyze

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

const tsPacketSize = 188

type TsHeader struct {
	PID                uint16
	PayloadUnitStart   bool
	AdaptationFieldCtl uint8
	ContinuityCounter  uint8
}

func parseTsHeader(pkt []byte) TsHeader {
	return TsHeader{
		PID:                ((uint16(pkt[1]) & 0x1F) << 8) | uint16(pkt[2]),
		PayloadUnitStart:   (pkt[1] & 0x40) != 0,
		AdaptationFieldCtl: (pkt[3] >> 4) & 0x03,
		ContinuityCounter:  pkt[3] & 0x0F,
	}
}

func findStartCode(data []byte) bool {
	for i := 0; i+3 < len(data); i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 &&
			(data[i+2] == 0x01 ||
				(data[i+2] == 0x00 && data[i+3] == 0x01)) {
			return true
		}
	}
	return false
}

func TestA01(t *testing.T) {

	f, err := os.Open("1.ts")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, tsPacketSize)

	pidCount := make(map[uint16]int)
	videoPID := uint16(0xFFFF)
	audioPID := uint16(0xFFFF)

	packetIndex := 0

	for {
		_, err := io.ReadFull(f, buf)
		if err != nil {
			break
		}
		packetIndex++

		if buf[0] != 0x47 {
			fmt.Printf("❌ Sync byte error at packet %d\n", packetIndex)
			return
		}

		h := parseTsHeader(buf)
		pidCount[h.PID]++

		payloadStart := 4
		if h.AdaptationFieldCtl == 2 || h.AdaptationFieldCtl == 3 {
			afLen := int(buf[4])
			payloadStart += 1 + afLen
		}
		if payloadStart >= tsPacketSize {
			continue
		}
		payload := buf[payloadStart:]

		// PAT
		if h.PID == 0 && h.PayloadUnitStart {
			fmt.Println("📌 PAT found")
		}

		// PMT (通常 PID != 0)
		if h.PayloadUnitStart && h.PID != 0 {
			if len(payload) > 4 && payload[0] == 0x00 {
				fmt.Printf("📌 Possible PSI table on PID 0x%x\n", h.PID)
			}
		}

		// PES
		if h.PayloadUnitStart && len(payload) >= 6 &&
			payload[0] == 0x00 && payload[1] == 0x00 && payload[2] == 0x01 {

			streamID := payload[3]
			pesLen := binary.BigEndian.Uint16(payload[4:6])

			switch {
			case streamID >= 0xE0 && streamID <= 0xEF:
				videoPID = h.PID
				fmt.Printf("🎥 Video PES: PID=0x%x stream_id=0x%x pes_len=%d\n",
					h.PID, streamID, pesLen)

				if findStartCode(payload[6:]) {
					fmt.Println("   ✅ Found Annex-B start code")
				} else {
					fmt.Println("   ❌ No Annex-B start code")
				}

			case streamID >= 0xC0 && streamID <= 0xDF:
				audioPID = h.PID
				fmt.Printf("🔊 Audio PES: PID=0x%x stream_id=0x%x pes_len=%d\n",
					h.PID, streamID, pesLen)
			}
		}
	}

	fmt.Println("\n==== TS Summary ====")
	fmt.Printf("Total packets: %d\n", packetIndex)
	fmt.Printf("Video PID: 0x%x\n", videoPID)
	fmt.Printf("Audio PID: 0x%x\n", audioPID)
	fmt.Println("PID usage:")
	for pid, c := range pidCount {
		fmt.Printf("  PID 0x%x : %d packets\n", pid, c)
	}
}

func TestName(t *testing.T) {
	//clientRecord.LastConnectTime.Add(3 * time.Minute).After(t)

	now := time.Now()
	LastConnectTime := now.Add(-5 * time.Minute)
	fmt.Println("now: ", now)
	fmt.Println("LastConnectTime: ", LastConnectTime)

	fmt.Println(LastConnectTime.Add(3 * time.Minute).After(now))
	fmt.Println(LastConnectTime.Add(3 * time.Minute).Before(now))

}

/*
📌 Possible PSI table on PID 0x11
📌 PAT found
📌 Possible PSI table on PID 0x1000
📌 Possible PSI table on PID 0x100
🎥 Video PES: PID=0x100 stream_id=0xe0 pes_len=0

	❌ No Annex-B start code

📌 Possible PSI table on PID 0x100
... 都是【No Annex-B start code】
🎥 Video PES: PID=0x100 stream_id=0xe0 pes_len=0

	❌ No Annex-B start code

📌 Possible PSI table on PID 0x101
🔊 Audio PES: PID=0x101 stream_id=0xc0 pes_len=1235

==== TS Summary ====
Total packets: 1109
Video PID: 0x100
Audio PID: 0x101
PID usage:

	PID 0x11 : 1 packets
	PID 0x0 : 2 packets
	PID 0x1000 : 2 packets
	PID 0x100 : 696 packets
	PID 0x101 : 408 packets
*/
func parseTsHeader2(pkt []byte) (pid uint16, pusi bool, afc uint8) {
	pid = ((uint16(pkt[1]) & 0x1F) << 8) | uint16(pkt[2])
	pusi = (pkt[1] & 0x40) != 0
	afc = (pkt[3] >> 4) & 0x03
	return
}

const (
	videoPID = 0x100
)

func TestParseES(t *testing.T) {
	f, err := os.Open("1.ts")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var es []byte
	buf := make([]byte, tsPacketSize)

	for {
		_, err := io.ReadFull(f, buf)
		if err != nil {
			break
		}
		if buf[0] != 0x47 {
			panic("sync byte error")
		}

		pid, pusi, afc := parseTsHeader2(buf)
		if pid != videoPID {
			continue
		}

		offset := 4
		if afc == 2 || afc == 3 {
			afLen := int(buf[4])
			offset += 1 + afLen
		}
		if offset >= tsPacketSize {
			continue
		}
		payload := buf[offset:]

		// PES start
		if pusi {
			// 跳过 PES header
			if len(payload) < 9 {
				continue
			}
			pesHeaderLen := int(payload[8])
			start := 9 + pesHeaderLen
			if start < len(payload) {
				es = append(es, payload[start:]...)
			}
		} else {
			es = append(es, payload...)
		}
	}

	fmt.Printf("Collected ES size: %d bytes\n", len(es))

	fmt.Println("ES first 64 bytes:")
	for i := 0; i < 64 && i < len(es); i++ {
		fmt.Printf("%02x ", es[i])
		if (i+1)%16 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	// ==== 解析 HEVC NALU ====
	i := 0
	esLenByte := 4
	for i+6 < len(es) {
		naluLen := int(binary.BigEndian.Uint32(es[i : i+esLenByte]))
		if naluLen <= 0 || i+esLenByte+naluLen > len(es) {
			fmt.Printf("Invalid NALU length at offset %d\n", i)
			break
		}

		naluHeader := es[i+esLenByte]
		naluType := (naluHeader >> 1) & 0x3F

		fmt.Printf("NALU @%-8d len=%-6d type=%d",
			i, naluLen, naluType)

		switch naluType {
		case 32:
			fmt.Println(" (VPS)")
		case 33:
			fmt.Println(" (SPS)")
		case 34:
			fmt.Println(" (PPS)")
		case 19, 20:
			fmt.Println(" (IDR)")
		default:
			fmt.Println()
		}

		i += esLenByte + naluLen

		// 只打印前几十个，避免刷屏
		if i > 200000 {
			break
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	tryParse(es, 4)

}
func tryParse(es []byte, lenSize int) {
	fmt.Printf("\nTrying length size = %d bytes\n", lenSize)
	i := 0
	n := 0
	for i+lenSize+1 < len(es) {
		var naluLen int
		switch lenSize {
		case 4:
			naluLen = int(binary.BigEndian.Uint32(es[i : i+4]))
		case 2:
			naluLen = int(binary.BigEndian.Uint16(es[i : i+2]))
		case 1:
			naluLen = int(es[i])
		}

		if naluLen <= 0 || i+lenSize+naluLen > len(es) {
			fmt.Printf("  ❌ invalid len=%d at offset %d\n", naluLen, i)
			return
		}

		naluType := (es[i+lenSize] >> 1) & 0x3F
		fmt.Printf("  NALU %d: off=%-6d len=%-5d type=%d\n",
			n, i, naluLen, naluType)

		i += lenSize + naluLen
		n++
		if n >= 10 {
			fmt.Println("  ... looks consistent")
			return
		}
	}
}

/*
ES first 64 bytes:
e3 35 cb 7e ad a0 39 e0 8f 29 fc 08 d6 3f 13 5a
f2 3f d5 d5 8e e7 a0 d3 e8 bd de 85 c5 ba fb bd
4a b7 28 5b e7 44 f5 4b 31 43 ab bd ae fb c1 8c
9c 1e 20 99 b6 7b 1f 5d 6d de bf 68 d5 a7 0d fa
*/

func Test122(t *testing.T) {
	f, err := os.Open("1.ts")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, tsPacketSize)
	annexBList := make([]int, 0)
	for {
		_, err := io.ReadFull(f, buf)
		if err != nil {
			break
		}
		if buf[0] != 0x47 {
			panic("sync byte error")
		}
		annexBList = append(annexBList, findAnnexB(buf))
	}
	for _, annexB := range annexBList {
		fmt.Printf("%d,", annexB)
	}
	// -1,11,-1,11,-1,-1,4,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,4,-1,6,-1,-1,4,-1,4,-1,4,-1,-1,4,-1,4,-1,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,26,4,-1,4,-1,-1,-1,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,4,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,40,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,15,4,-1,50,55,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,65,50,4,-1,45,55,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,4,-1,25,48,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,35,17,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,41,4,-1,4,-1,-1,11,4,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,4,-1,4,-1,4,-1,11,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,44,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,4,-1,29,53,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,61,17,4,-1,-1,-1,-1,-1,11,-1,11,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,4,-1,-1,4,-1,4,-1,4,-1,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,51,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,17,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,72,64,45,60,74,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,4,81,82,16,71,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,67,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,4,-1,73,71,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,58,50,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,94,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,95,104,100,104,106,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,-1,-1,-1,-1,-1,-1,-1,-1,4,-1,94,106,106,104,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,106,106,4,-1,-1,-1,-1,-1,89,105,106,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,106,106,106,106,48,101,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,105,106,106,106,106,106,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,102,104,105,106,106,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,106,106,106,105,104,105,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,106,106,106,106,106,6,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,-1,105,105,6,-1,-1,-1,-1,-1,-1,
	fmt.Println()
}
func findAnnexB(data []byte) int {
	for i := 0; i+3 < len(data); i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 &&
			(data[i+2] == 0x01 ||
				(data[i+2] == 0x00 && data[i+3] == 0x01)) {
			return i
		}
	}
	return -1
}

func Test123(t *testing.T) {
	f, err := os.Open("1.ts")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, tsPacketSize)
	for {
		_, err := io.ReadFull(f, buf)
		if err != nil {
			break
		}
		if buf[0] != 0x47 {
			panic("sync byte error")
		}
		//fmt.Println(countEPB(buf))
		scanHEVCNALU(buf)
	}
}
func scanHEVCNALU(data []byte) {
	for i := 0; i < len(data)-1; i++ {
		naluType := (data[i] >> 1) & 0x3f
		if naluType <= 63 {
			fmt.Printf("possible nalu type %d at %d\n", naluType, i)
		}
	}
}

func countEPB(data []byte) int {
	cnt := 0
	for i := 0; i+2 < len(data); i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 && data[i+2] == 0x03 {
			cnt++
		}
	}
	return cnt
}
