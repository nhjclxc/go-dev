package main

import (
	"fmt"
	"io"
	"os"
)

// TS Packet size
const TS_PKT_SIZE = 188

type TsHeader struct {
	PID  uint16
	PUSI bool
	AFC  uint8
	CC   uint8
}

// parseTsHeader 解析 ts头信息
func parseTsHeader(pkt []byte) TsHeader {
	return TsHeader{
		PID:  ((uint16(pkt[1]) & 0x1F) << 8) | uint16(pkt[2]),
		PUSI: (pkt[1] & 0x40) != 0,
		AFC:  (pkt[3] >> 4) & 0x03,
		CC:   pkt[3] & 0x0F,
	}
}

// getTsFileBytesData 从文件开头跳过34字节
func getTsFileBytesData(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	_, err = file.Seek(34, io.SeekStart)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// scanPIDStats 扫描 TS，统计 PID 出现次数
func scanPIDStats(path string) (map[uint16]int, error) {
	f, err := getTsFileBytesData(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stats := make(map[uint16]int)
	buf := make([]byte, TS_PKT_SIZE)

	for {
		_, err := f.Read(buf)
		if err != nil {
			break
		}
		if buf[0] != 0x47 {
			continue
		}
		h := parseTsHeader(buf)
		stats[h.PID]++
	}
	return stats, nil
}

// extractPES 提取指定 PID 的 PES 数据
func extractPES(path string, targetPID uint16, streamStart byte, streamEnd byte) ([]byte, error) {
	f, err := getTsFileBytesData(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var pes []byte
	buf := make([]byte, TS_PKT_SIZE)

	for {
		_, err := f.Read(buf)
		if err != nil {
			break
		}
		if buf[0] != 0x47 {
			continue
		}
		h := parseTsHeader(buf)
		if h.PID != targetPID {
			continue
		}

		// Adaptation Field
		pos := 4
		if h.AFC == 2 || h.AFC == 3 {
			afLen := int(buf[4])
			pos += 1 + afLen
		}
		if pos >= TS_PKT_SIZE {
			continue
		}
		payload := buf[pos:]

		// PES start
		if h.PUSI && len(payload) > 3 &&
			payload[0] == 0x00 &&
			payload[1] == 0x00 &&
			payload[2] == 0x01 &&
			payload[3] >= streamStart && payload[3] <= streamEnd {
			pes = append(pes, payload...)
		} else if len(pes) > 0 {
			pes = append(pes, payload...)
		}
	}
	return pes, nil
}

// stripPESHeader 跳过 PES header
func stripPESHeader(pes []byte) []byte {
	if len(pes) < 9 {
		return nil
	}
	if pes[0] != 0x00 || pes[1] != 0x00 || pes[2] != 0x01 {
		return pes
	}
	headerLen := int(pes[8])
	start := 9 + headerLen
	if start >= len(pes) {
		return nil
	}
	return pes[start:]
}

// findH264ParamSets 找 SPS / PPS
func findH264ParamSets(es []byte) (sps, pps []byte) {
	i := 0
	for i+4 < len(es) {
		if es[i] == 0 && es[i+1] == 0 &&
			(es[i+2] == 1 || (es[i+2] == 0 && es[i+3] == 1)) {

			start := i
			if es[i+2] == 1 {
				i += 3
			} else {
				i += 4
			}
			nalType := es[i] & 0x1F

			// 找下一个 start code
			next := i
			for next+3 < len(es) &&
				!(es[next] == 0 && es[next+1] == 0 &&
					(es[next+2] == 1 || (es[next+2] == 0 && es[next+3] == 1))) {
				next++
			}

			nalu := es[start:next]

			if nalType == 7 {
				sps = nalu
			} else if nalType == 8 {
				pps = nalu
			}
			i = next
		} else {
			i++
		}
	}
	return
}

// 给 AAC 帧加 ADTS header
func addADTSHeader(aac []byte, profile, freqIdx, chanCfg int) []byte {
	frameLen := len(aac) + 7
	adts := []byte{
		0xFF, 0xF1,
		byte(((profile - 1) << 6) | ((freqIdx & 0x0F) << 2) | ((chanCfg >> 2) & 0x01)),
		byte(((chanCfg & 3) << 6) | ((frameLen >> 11) & 0x03)),
		byte((frameLen >> 3) & 0xFF),
		byte(((frameLen & 7) << 5) | 0x1F),
		0xFC,
	}
	return append(adts, aac...)
}

// checkNALU 检查 NALU
func checkNALU(data []byte) bool {
	for i := 0; i < len(data)-4; i++ {
		if data[i] == 0x00 && data[i+1] == 0x00 &&
			(data[i+2] == 0x01 || (data[i+2] == 0x00 && data[i+3] == 0x01)) {
			fmt.Println("Found NALU start code at", i)
			return true
		}
	}
	fmt.Println("No NALU found")
	return false
}

// checkAACType 检查 AAC
func checkAACType(data []byte) {
	for i := 0; i < len(data)-2; i++ {
		if data[i] == 0xFF && (data[i+1]&0xF0) == 0xF0 {
			fmt.Println("AAC ADTS detected")
			return
		}
	}
	fmt.Println("AAC LATM (no ADTS header)")
}

func main() {
	tsPath := "input.ts"
	videoPID := uint16(256)
	audioPID := uint16(257)

	// 1️⃣ PID stats
	stats, _ := scanPIDStats(tsPath)
	fmt.Println("PID statistics:")
	for pid, cnt := range stats {
		fmt.Printf("PID %-5d (0x%04x): %d packets\n", pid, pid, cnt)
	}

	// 2️⃣ 提取视频 PES
	videoPES, _ := extractPES(tsPath, videoPID, 0xE0, 0xE0)
	checkNALU(videoPES)
	videoES := stripPESHeader(videoPES)

	sps, pps := findH264ParamSets(videoES)
	outV, _ := os.Create("video.h264")
	if sps != nil {
		outV.Write(sps)
	}
	if pps != nil {
		outV.Write(pps)
	}
	outV.Write(videoES)
	outV.Close()

	// 3️⃣ 提取音频 PES
	audioPES, _ := extractPES(tsPath, audioPID, 0xC0, 0xDF)
	checkAACType(audioPES)
	audioES := stripPESHeader(audioPES)
	adts := addADTSHeader(audioES, 2, 4, 2) // profile 2 = AAC-LC, freqIdx 4 = 44100Hz, chanCfg 2
	outA, _ := os.Create("audio.aac")
	outA.Write(adts)
	outA.Close()

	fmt.Println("Extraction done: video.h264 + audio.aac")
}

// ffmpeg -y -fflags +genpts -r 25 -i video.h264 -i audio.aac -c:v copy -c:a copy -movflags +faststart out.mp4
// ffmpeg -y -f h264 -r 25 -i video.h264 -i audio.aac -c:v copy -c:a copy -movflags +faststart out.mp4
