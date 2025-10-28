package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	// 生成简单的SSH密钥对
	config := &ssh.ServerConfig{
		NoClientAuth: true, // 为了简化demo，不需要客户端认证
	}

	// 生成临时的host key
	privateKey, err := generateHostKey()
	if err != nil {
		log.Fatal("Failed to generate host key:", err)
	}
	config.AddHostKey(privateKey)

	// 监听SSH连接
	listener, err := net.Listen("tcp", "0.0.0.0:2222")
	if err != nil {
		log.Fatal("Failed to listen on 2222:", err)
	}
	defer listener.Close()

	log.Println("SSH server listening on :2222")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn, config)
	}
}

func handleConnection(conn net.Conn, config *ssh.ServerConfig) {
	defer conn.Close()

	// SSH握手
	sshConn, chans, reqs, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Println("Failed to handshake:", err)
		return
	}
	defer sshConn.Close()

	log.Printf("New SSH connection from %s (Android client)", sshConn.RemoteAddr())

	// 处理全局请求（包括端口转发请求）
	go handleRequests(reqs, sshConn)

	// 处理通道请求
	for newChannel := range chans {
		log.Printf("New channel request: %s", newChannel.ChannelType())

		// 接受forwarded-tcpip类型的通道
		if newChannel.ChannelType() == "forwarded-tcpip" {
			channel, requests, err := newChannel.Accept()
			if err != nil {
				log.Println("Could not accept channel:", err)
				continue
			}
			go ssh.DiscardRequests(requests)
			go handleForwardedChannel(channel)
		} else {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
		}
	}
}

func handleRequests(reqs <-chan *ssh.Request, sshConn *ssh.ServerConn) {
	for req := range reqs {
		log.Printf("Received request: %s", req.Type)

		switch req.Type {
		case "tcpip-forward":
			// 解析端口转发请求
			var payload struct {
				Address string
				Port    uint32
			}

			if err := ssh.Unmarshal(req.Payload, &payload); err != nil {
				log.Println("Failed to parse tcpip-forward request:", err)
				if req.WantReply {
					req.Reply(false, nil)
				}
				continue
			}

			log.Printf("Client requested reverse port forwarding: %s:%d", payload.Address, payload.Port)

			// 启动端口监听
			go startPortForwarding(payload.Port, sshConn)

			if req.WantReply {
				// 回复绑定的端口
				var response struct {
					Port uint32
				}
				response.Port = payload.Port
				req.Reply(true, ssh.Marshal(&response))
			}

		case "cancel-tcpip-forward":
			log.Println("Client requested to cancel port forwarding")
			if req.WantReply {
				req.Reply(true, nil)
			}

		default:
			if req.WantReply {
				req.Reply(false, nil)
			}
		}
	}
}

func startPortForwarding(port uint32, sshConn *ssh.ServerConn) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("Failed to listen on %s: %v", addr, err)
		return
	}
	defer listener.Close()

	log.Printf("Now listening on %s for ADB connections", addr)

	for {
		localConn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		log.Printf("New ADB connection from %s", localConn.RemoteAddr())
		go forwardToSSH(localConn, sshConn, port)
	}
}

func forwardToSSH(localConn net.Conn, sshConn *ssh.ServerConn, port uint32) {
	defer localConn.Close()

	// 打开一个forwarded-tcpip通道到客户端
	payload := ssh.Marshal(&struct {
		Address           string
		Port              uint32
		OriginatorAddress string
		OriginatorPort    uint32
	}{
		Address:           "127.0.0.1",
		Port:              port,
		OriginatorAddress: "127.0.0.1",
		OriginatorPort:    uint32(localConn.RemoteAddr().(*net.TCPAddr).Port),
	})

	channel, reqs, err := sshConn.OpenChannel("forwarded-tcpip", payload)
	if err != nil {
		log.Println("Failed to open forwarded-tcpip channel:", err)
		return
	}
	defer channel.Close()
	go ssh.DiscardRequests(reqs)

	log.Println("Forwarding connection through SSH tunnel")

	// 双向转发数据
	done := make(chan struct{}, 2)

	go func() {
		io.Copy(channel, localConn)
		done <- struct{}{}
	}()

	go func() {
		io.Copy(localConn, channel)
		done <- struct{}{}
	}()

	<-done
	log.Println("Connection closed")
}

func handleForwardedChannel(channel ssh.Channel) {
	defer channel.Close()
	// 这个函数现在不需要了，因为我们主动打开通道
}

func generateHostKey() (ssh.Signer, error) {
	// 使用预生成的测试密钥（简化demo）
	privateKeyPEM := []byte(`-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
NhAAAAAwEAAQAAAQEA+lJf4TKjdpNLuG1VMf7ujHbfku54HJRVypna5UFkrxpPCpJhrCfr
q0wf5ggJ8xTr9KhVbBXJtdu2f+OqbdlJqbBj0HpVQeXIUM6JSD2e4cKRnxBGNgEQm0B82G
jjQQ2+w6kKTzIh7Vzgz97pL2//VEkNlIbI9QVzhTEmlFWlmxqVB6sveTKXAEjQGRPLdJKJ
SixBISci+Xa0kSpBKM6/opsoAM/OmuiUmLIipvGd55TLcvUM9mcDKog2fDanK/OkjBg7oY
X7RRP2X7xzZona84w088sI1CXVurRFgiocl9TJvYSck8cwEgWAvM8jdlwPjqeknMXrcGuQ
z7ARQ0cYuwAAA+DOtmTQzrZk0AAAAAdzc2gtcnNhAAABAQD6Ul/hMqN2k0u4bVUx/u6Mdt
+S7ngclFXKmdrlQWSvGk8KkmGsJ+urTB/mCAnzFOv0qFVsFcm127Z/46pt2UmpsGPQelVB
5chQzolIPZ7hwpGfEEY2ARCbQHzYaONBDb7DqQpPMiHtXODP3ukvb/9USQ2Uhsj1BXOFMS
aUVaWbGpUHqy95MpcASNAZE8t0kolKLEEhJyL5drSRKkEozr+imygAz86a6JSYsiKm8Z3n
lMty9Qz2ZwMqiDZ8Nqcr86SMGDuhhftFE/ZfvHNmidrzjDTzywjUJdW6tEWCKhyX1Mm9hJ
yTxzASBYC8zyN2XA+Op6Scxetwa5DPsBFDRxi7AAAAAwEAAQAAAQAkxk1HrETPerw5D/bg
LW+mOyCFWXtT14bzCL9bxsuf3eGW0AKBZCv/MmPNS4kKqRwxCTnjcx/7E5gwGaZEFRP5Ve
yUCCTWVYekN3N7pXHeANOb5qwp6uYhqMuuj1oziG1qGT9mexr44Bg28ayYR3/fRgw0sch8
FpzuCFZ+nEaEDo1M9WheihAd2m0fzwHleKbS2Lz7Yos71FTZFQVfJ7dFdro5h6LJMBRmO0
ErgrZTKaa0ffUGmkwVdvcyu6gf1JDQwxYrfoDEnieOAZtPBTQnn6fTu2zzLpDcXMY+q9vA
i5AZws7xIqtHhoHICkqc6C8d4uVfyH/MvSXSPKchHasBAAAAgQCn3oG2tnRNOpnSb6B2Hy
nqBIdqQ9/NV+hNc9lJHZrTiTpDw686htF1prs//RxxJf/6QUoSBfZ0w60DRTASuznuIP6e
Az9SV8hDdnjwPKdRNL0UcQMWUErGtdnCOSlShlqmyF9WXy/Uuhh3OTXCKI1rfAGNz6dtpF
j5OLfRvQK16wAAAIEA/jNgEG+a7KtfvNjPRYwmY7xkOBGRrfP3l2QgMtjqyvbBgsUoMp06
zZrIv1jVPEGxgTwgxvDmRwHaifW5mCDR7C0MeahaVg4sRmNOLLa9YqlJhXFbg+k6REU3Ml
emRDrcckuG/KytivMKTOv7lOe5e/tAdYn2StiKEnasYOUrhiEAAACBAPwX+HJsHsYXy1KU
f5yny9wTKf6HP1hZNew24XgGlPmIRiMm+Nt2vh0ve6TagSiXUVG1YDLBjyDHR/rCL1ey1m
cUYuRdO+kzGtce9zpwZYnsWk+y2QG197ETRzBtyGFyJzpEtj2hgA4kkhZkxCvF+O/xy0j0
BjMDYcw3jX2WqwtbAAAAKGNoZW5qaWF5YW9AY2hlbmppYXlhb2RlTWFjQm9vay1Qcm8ubG
9jYWwBAg==
-----END OPENSSH PRIVATE KEY-----`)

	return ssh.ParsePrivateKey(privateKeyPEM)
}
