package main


import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/alexbrainman/odbc"
)


func main() {
	// 修改为你的实际路径
	//dbPath := `D:\code\go\go-dev\src\test7_db\test71_access\Database1.accdb`
	//dsn := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=%s;", dbPath)

	//path := `D:\code\go\go-dev\src\test7_db\test71_access\Database1.accdb`
	//pwd := ""
	//fixed := 5

	// go build -o MixingPlant22.exe MixingPlant.go
	// MixingPlant.exe  -path=D:\code\go\go-dev\src\test7_db\test71_access\Database1.accdb -pwd= -fixed=5 -uuid=abc123 -env=1
	// rsrc -ico favicon.ico -o rsrc.syso

	// 定义命令行参数
	uuid := flag.String("uuid", "aaa111", "全局唯一uuid")
	path := flag.String("path", "Database1.accdb", "Access 数据库路径")
	pwd := flag.String("pwd", "", "Access 数据库密码")
	fixed := flag.Int("fixed", 60, "固定参数（整数）")
	env := flag.Int("env", 0, "启动环境")

	// 解析命令行参数
	flag.Parse()

	// 打印读取结果
	fmt.Println("全局唯一uuid:", *uuid)
	fmt.Println("数据库路径:", *path)
	fmt.Println("数据库密码:", *pwd)
	fmt.Println("固定参数 fixed:", *fixed)
	fmt.Println("启动环境:", *env)


	//*path = `D:\code\go\go-dev\src\test7_db\test71_access\Database1.accdb`
	dsn := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=%s;PWD=%s;", *path, *pwd)


	db, err := sql.Open("odbc", dsn)
	if err != nil {
		log.Fatal("无法打开数据库:", err)
	}
	defer db.Close()

	// 定时每分钟执行
	//ticker := time.NewTicker(time.Second * 5)
	ticker := time.NewTicker(time.Second * time.Duration(*fixed))
	defer ticker.Stop()


	domain := "https://127.0.0.1:8080"
	if *env == 1 {
		domain = "https://api-sc-dev.hkznkj.com"
	} else if *env == 2 {
		domain = "https://api-sc.hkznkj.com"
	}

	for range ticker.C {

		//// 获取当前时间和一分钟前时间
		now := time.Now()
		oneMinuteAgo := now.Add(time.Second * -1 * time.Duration(*fixed))

		fmt.Printf("查询时间段：start: " + oneMinuteAgo.String() + ", end: " + now.String() + "\n")

		var resultsDosage []Dosage = queryDosage(db, now, oneMinuteAgo)

		if resultsDosage != nil && len(resultsDosage) > 0 {
			dosageBody, errdosage := json.Marshal(resultsDosage)
			if errdosage == nil {
				uploadToServer(string(dosageBody), domain + "/device/mining/plant/dosage/" + *uuid)
			}
		} else {
			fmt.Printf("Dosage 数据为空！\n")
		}

		var resultsPiece []Piece = queryPiece(db, now, oneMinuteAgo)
		if resultsPiece != nil && len(resultsPiece) > 0 {
			pieceBody, errpiece := json.Marshal(resultsPiece)
			if errpiece == nil {
				uploadToServer(string(pieceBody), domain + "/device/mining/plant/piece/" + *uuid)
			}
		} else {
			fmt.Printf("Piece 数据为空！\n")
		}

		var resultsProduce []Produce = queryProduce(db, now, oneMinuteAgo)
		if resultsProduce != nil && len(resultsProduce) > 0 {
			produceBody, errproduce := json.Marshal(resultsProduce)
			if errproduce == nil {
				uploadToServer(string(produceBody), domain + "/device/mining/plant/produce/" + *uuid)
			}
		} else {
			fmt.Printf("Produce 数据为空！\n")
		}
	}
}


func uploadToServer(body, url string) {

	fmt.Printf("url = %s, request = %s. \n", url, body)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 跳过证书校验
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Post(url, "application/json", strings.NewReader(string(body)))
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf(fmt.Sprintf("上传失败，HTTP 状态码: %d", resp.StatusCode))
		return
	}

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// 将响应体转换为字符串（如果需要）
	respStr := string(respBody)
	fmt.Printf("url = %s, response = %s. \n", url, respStr)

}



func queryProduce(db *sql.DB, now time.Time, oneMinuteAgo time.Time) ([]Produce) {


	query := fmt.Sprintf(`
		SELECT * FROM Produce  WHERE Stamp BETWEEN ? AND ?
	`)

	rows, err := db.Query(query, now, oneMinuteAgo)
	if err != nil {
		log.Println("查询失败:", err)
		return nil
	}
	defer rows.Close()


	var results []Produce
	for rows.Next() {
		var p Produce
		err = rows.Scan(
			&p.ID, &p.Code, &p.DatTim, &p.Attribute, &p.Contract,
			&p.Customer,
			&p.ProjName, &p.ProjType, &p.ProjGrade, &p.ProjArea, &p.ProjAdr,
			&p.Distance, &p.ConsPos, &p.Pour, &p.Variety, &p.BetLev, &p.Filter,
			&p.Freeze, &p.Lands, &p.Cement, &p.Stone, &p.EnSize, &p.AddLiq,
			&p.Request, &p.Recipe, &p.MixLast, &p.MorRec, &p.Mete, &p.BegTim,
			&p.EndTim, &p.Attamper, &p.Data, &p.Flag, &p.Notes, &p.Vehicle,
			&p.Driver, &p.ProdTimB, &p.ProdTimE, &p.ProdMete, &p.MorMete,
			&p.ProdErr, &p.ProdCnt, &p.MorCnt, &p.TotVehs, &p.TotMete, &p.Qualitor,
			&p.Operator, &p.LeftTim, &p.ArriveTim, &p.ChkLands, &p.ChkTemp,
			&p.UnloadTim, &p.OverTim, &p.Acceptor, &p.Mark, &p.MISID,
			&p.Stamp,
			&p.Task, &p.Contacts, &p.ContTel, &p.Bend,
		)

		if err != nil {
			log.Println("解析数据失败:", err)
			continue
		}
		results = append(results, p)
	}

	return results
}


func queryPiece(db *sql.DB, now time.Time, oneMinuteAgo time.Time) ([]Piece) {

	query := fmt.Sprintf(`
		SELECT * FROM Piece WHERE Stamp BETWEEN ? AND ?
	`)


	rows, err := db.Query(query, now, oneMinuteAgo)
	if err != nil {
		log.Println("查询失败:", err)
		return nil
	}
	defer rows.Close()


	var results []Piece
	for rows.Next() {
		var piece Piece
		err = rows.Scan(
			&piece.ID,&piece.Produce,&piece.RecID,&piece.Recipe,&piece.Serial,
			&piece.Blender, &piece.DatTim,
			&piece.BldTim, &piece.PieAmnt, &piece.Lands, &piece.Temper, &piece.PieErr, &piece.Data, &piece.Flag,
			&piece.Stamp,
		)

		if err != nil {
			log.Println("解析数据失败:", err)
			continue
		}
		results = append(results, piece)
	}

	return results
}


func queryDosage(db *sql.DB, now time.Time, oneMinuteAgo time.Time) ([]Dosage) {


	query := fmt.Sprintf(`
		SELECT * FROM Dosage WHERE Stamp BETWEEN ? AND ?
	`)

	rows, err := db.Query(query, now, oneMinuteAgo)
	if err != nil {
		log.Println("查询失败:", err)
		return nil
	}
	defer rows.Close()


	var results []Dosage
	for rows.Next() {
		var dosage Dosage
		err = rows.Scan(
			&dosage.ID,
			&dosage.Piece,
			&dosage.StorlD,
			&dosage.Storage,
			&dosage.MaterlD,
			&dosage.Material,
			&dosage.RecAmnt,
			&dosage.PlanAmnt,
			&dosage.FactAjnnt,
			&dosage.Fall,
			&dosage.FinTim,
			&dosage.Data,
			&dosage.Flag,
			&dosage.Stamp,
		)
		if err != nil {
			log.Println("解析数据失败:", err)
			continue
		}
		results = append(results, dosage)
	}

	return results
}






// 派车记录表
type Produce struct {
	ID        int     `json:"id"`         // 主键
	Code      string  `json:"code"`       // 任务单编号
	DatTim    string  `json:"datTim"`     // 创建日期
	Attribute string  `json:"attribute"`  // 任务性质
	Contract  string  `json:"contract"`   // 合同信息
	Customer  string  `json:"customer"`   // 客户名称
	ProjName  string  `json:"projName"`   // 工程名称
	ProjType  string  `json:"projType"`   // 工程类别
	ProjGrade string  `json:"projGrade"`  // 工程级别
	ProjArea  float64 `json:"projArea"`   // 开工面积
	ProjAdr   string  `json:"projAdr"`    // 施工地址
	Distance  float64 `json:"distance"`   // 运输距离
	ConsPos   string  `json:"consPos"`    // 施工部位
	Pour      string  `json:"pour"`       // 浇筑方式
	Variety   string  `json:"variety"`    // 产品种类
	BetLev    string  `json:"betLev"`     // 强度等级
	Filter    string  `json:"filter"`     // 抗渗等级
	Freeze    string  `json:"freeze"`     // 抗冻等级
	Lands     string  `json:"lands"`      // 坍落度
	Cement    string  `json:"cement"`     // 水泥品种
	Stone     string  `json:"stone"`      // 石子种类

	EnSize     string  `json:"enSize"`     // 骨科粒径
	AddLiq     string  `json:"addLiq"`     // 外加剂要求
	Request    string  `json:"request"`    // 技术要求
	Recipe     string  `json:"recipe"`     // 施工配合比
	MixLast    string  `json:"mixLast"`    // 搅拌时间
	MorRec     string  `json:"morRec"`     // 砂浆配比
	Mete       float64 `json:"mete"`       // 任务方量
	BegTim     string  `json:"begTim"`     // 浇筑日期
	EndTim     string  `json:"endTim"`     // 截止日期
	Attamper   string  `json:"attamper"`   // 任务调度
	Data       string  `json:"data"`       // 附加数据
	Flag       string  `json:"flag"`       // 标记
	Notes      string  `json:"notes"`      // 备注
	Vehicle    string  `json:"vehicle"`    // 车辆ID
	Driver     string  `json:"driver"`     // 驾驶员
	ProdTimB   string  `json:"prodTimB"`   // 开始生产时刻
	ProdTimE   string  `json:"prodTimE"`   // 结束生产时刻
	ProdMete   float64 `json:"prodMete"`   // 生产方量
	MorMete    float64 `json:"morMete"`    // 砂浆方量
	ProdErr    float64 `json:"prodErr"`    // 车误差
	ProdCnt    int     `json:"prodCnt"`    // 生产盘数
	MorCnt     int     `json:"morCnt"`     // 砂浆盘数

	TotVehs    int     `json:"totVehs"`    // 累计车次
	TotMete    float64 `json:"totMete"`    // 累计方量
	Qualitor   string  `json:"qualitor"`   // 质检员
	Operator   string  `json:"operator"`   // 操作员
	LeftTim    string  `json:"leftTim"`    // 出站时间
	ArriveTim  string  `json:"arriveTim"`  // 到达时间
	ChkLands   string  `json:"chkLands"`   // 检测坍落度
	ChkTemp    string  `json:"chkTemp"`    // 卸砼温度
	UnloadTim  string  `json:"unloadTim"`  // 卸料时间
	OverTim    string  `json:"overTim"`    // 卸完时间
	Acceptor   string  `json:"acceptor"`   // 现场验收
	Mark       string  `json:"mark"`       // 总第n车
	MISID      string  `json:"misId"`      // MIS系统ID
	Stamp      string  `json:"stamp"`      // 时间戳
	Task       string  `json:"task"`       // 任务标识
	Contacts   string  `json:"contacts"`   // 联系人
	ContTel    string  `json:"contTel"`    // 联系电话
	Bend       string  `json:"bend"`       // 弯沉
}

// 盘次记录表
type Piece struct {
	ID       string  `json:"id"`        // ID
	Produce  string  `json:"produce"`   // 一次配方ID 对应Produce.ID
	RecID    string  `json:"recId"`     // 配方ID
	Recipe   string  `json:"recipe"`    // 配方
	Serial   string  `json:"serial"`    // 序列号
	Blender  string  `json:"blender"`   // 搅拌机
	DatTim   string  `json:"datTim"`    // 生产时刻
	BldTim   string  `json:"bldTim"`    // 搅拌时间
	PieAmnt  float64 `json:"pieAmnt"`   // 盘方量
	Lands    string  `json:"lands"`     // 盘坍落度
	Temper   string  `json:"temper"`    // 盘温度
	PieErr   float64 `json:"pieErr"`    // 盘误差
	Data     string  `json:"data"`      // 附加数据
	Flag     string  `json:"flag"`      // 标识
	Stamp    string  `json:"stamp"`     // 更新时间
}


// 原料消耗表
type Dosage struct {
	ID        string  `json:"id"`         // 主键ID
	Piece     string  `json:"piece"`      // 盘次ID 对应PieceID.ID
	StorlD    string  `json:"storlId"`    // 原料ID
	Storage   string  `json:"storage"`    // 原料料仓
	MaterlD   string  `json:"materlId"`   // 原材料ID
	Material  string  `json:"material"`   // 原材料
	RecAmnt   float64 `json:"recAmnt"`    // 配方方量
	PlanAmnt  float64 `json:"planAmnt"`   // 理论用量
	FactAjnnt float64 `json:"factAjnnt"`  // 实际用量
	Fall      float64 `json:"fall"`       // 当前落差
	FinTim    string  `json:"finTim"`     // 完成时刻
	Data      string  `json:"data"`       // 附加数据
	Flag      string  `json:"flag"`       // 标识
	Stamp     string  `json:"stamp"`      // 更新时间
}

