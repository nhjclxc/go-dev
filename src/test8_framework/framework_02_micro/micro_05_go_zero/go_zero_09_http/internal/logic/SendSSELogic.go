package logic

import (
	"context"
	"fmt"
	"go_zero_09_http/internal/deepseek"
	"go_zero_09_http/internal/svc"
	"go_zero_09_http/internal/types"
	"log"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSSELogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	// 用于保存sse的客户端
	clients map[string]http.ResponseWriter
}

var staticData []string = []string{
	"人生若只如初见，何事秋风悲画扇。等闲变却故人心，却道故人心易变。骊山语罢清宵半，泪雨零铃终不怨。何如薄幸锦衣郎，比翼连枝当日愿。",
	"春水初生，春林初盛，春风十里，不如你。时光荏苒，岁月无声，唯有那一抹笑容，温暖了我整个冬天。每当想起你，心中便充满了柔情。",
	"远方的山川河流，诉说着千百年来的故事。落日余晖洒在大地上，金色的光芒映照出生命的坚韧与希望。无论风雨多大，心中的梦想永不熄灭。",
	"夜空中闪烁的星辰，如同无数双眼睛注视着这个世界。它们见证了无数的喜怒哀乐，也守护着每一个孤独的灵魂。漫步星河，感受宇宙的浩瀚与神秘。",
	"人生如逆旅，我亦是行人。每一段旅程，都是成长的印记。跌倒时擦干眼泪，继续前行，终会到达梦想的彼岸。无论多远，心之所向，即为归处。",
	"花开花落春常在，云卷云舒自无言。岁月的脚步轻轻踏过指尖，带走了青春，也留下了回忆。那些温暖的瞬间，如同烙印，永远铭刻在心底。",
	"晨曦初露，空气中弥漫着泥土的芬芳。鸟儿欢唱，唤醒沉睡的大地。新的一天开始了，充满无限可能与希望。愿我们都能勇敢追梦，不负此生。",
	"书页轻翻，墨香四溢。文字像涓涓细流，滋养着心灵的花朵。无论世界多么喧嚣，书籍总能带来片刻宁静，让思绪翱翔于无垠的天空。",
	"海浪拍打着岸边的礁石，发出阵阵悠扬的声音。那是大海的低语，诉说着无尽的深情与宽广。站在海边，感受风的拥抱，内心也随之澄明宁静。",
	"灯火阑珊处，旧城的巷弄里藏着多少故事。斑驳的墙壁，诉说着岁月的沧桑。那些曾经的人和事，化作时间的烟尘，轻轻飘散在风中，久久不散。",
}

// 获取订单信息
func NewSendSSELogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSSELogic {
	return &SendSSELogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		clients: make(map[string]http.ResponseWriter),
	}
}


// 模拟消息生成
// SimulateEvents 模拟周期性事件
func (h *SendSSELogic) SimulateEvents() {
	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	//
	//for range ticker.C {
	//	message := fmt.Sprintf("Server time: %s", time.Now().Format(time.RFC3339))
	//	// 广播给所有客户端
	//	for clientChan := range h.clients {
	//		select {
	//		case clientChan <- message:
	//		default:
	//			// 跳过阻塞的 channel
	//		}
	//	}
	//}
}

// Serve 处理 SSE 连接
func (sendSSELogic *SendSSELogic) SendSSE(req *types.SendSSEReq, w http.ResponseWriter, r *http.Request) {
	// todo: add your logic here and delete this line

	fmt.Printf(" SendSSE = %v \n", req)

	// 设置 SSE 必需的 HTTP 头
	// for versions > v1.8.1, no need to add 3 lines below
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")

	// CORS 设置访问源地址
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Allowed-Header", "*")
	w.Header().Add("Allowed-Method", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		panic("flusher not supported")
	}

	//// 先用当前时间作为种子，保证每次运行结果不同
	//rand.Seed(time.Now().UnixNano())
	//
	//// 生成 1 到 10 的随机数，rand.Intn(10) 返回 0~9，所以 +1
	//randIndex := rand.Intn(10) + 1
	//
	//fmt.Println("随机数:", randIndex)
	//
	//itemContent := staticData[randIndex]
	//for _, ch := range itemContent {
	//	_, err := fmt.Fprintf(w, "%c", ch) // 按字符写入
	//	if err != nil {
	//		log.Println("写入失败:", err)
	//		break
	//	}
	//	flusher.Flush()
	//	time.Sleep(100 * time.Millisecond)
	//}

	// 调用deepseek实现聊天


	//content := "给我作一首诗"

	//var msgChan chan []string = make(chan []string, 500)
	var msgChan chan string = make(chan string)
	var exitChan chan bool = make(chan bool)
	//defer close(msgChan)
	//defer close(exitChan)


	// 写数据
	go deepseek.SendDeepSeek(exitChan, msgChan, req.Content)

	// 读数据
	for {
		select {
		case <-exitChan:
			return
		case msg := <-msgChan:
			_, err := fmt.Fprintf(w, "%s", msg) // 按字符写入
			if err != nil {
				log.Println("写入失败:", err)
				break
			}
			flusher.Flush()
		}
	}


}
