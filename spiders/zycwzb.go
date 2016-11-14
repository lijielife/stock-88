package spiders

// 基础包
import (

	//必需
	//DOM解析
	//DOM解析

	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/henrylee2cn/pholcus/app/downloader/request"
	. "github.com/henrylee2cn/pholcus/app/spider"
	//必需
	//信息输出
	//DOM解析
	//DOM解析
	// . "github.com/henrylee2cn/pholcus/app/spider/common" //选用
	// net包
	// "net/http" //设置http.Header
	// "net/url"
	// 编码包
	// "encoding/xml"
	// 字符串处理包
	// "strconv"
	// 其他包
	// "fmt"
	// "math"
	// "time"
)

func init() {
	list.Register()
}

var (
	listurl = "https://xueqiu.com/stock/cata/stocklist.json?page=%d&size=90&type=%d,%d"
	cwzburl = "https://xueqiu.com/stock/f10/finmainindex.json?symbol=%s&page=1&size=100"
	cookie  = "s=7714unyv2f; xq_a_token=e8e54f45363e45c762c42cd926e470175c1123a1; xq_r_token=36c71dba39380a7c439676ea63f63f717faa677b; Hm_lvt_1db88642e346389874251b5a1eded6e3=1479109387,1479109582,1479117256; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1479117256"
	now     = time.Now()
)

var list = &Spider{
	Name:        "雪球财务指标",
	Description: "获取雪球主要财务指标数据",
	// Pausetime:    300,
	// Keyin:        KEYIN,
	// Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {

			ctx.AddQueue(&request.Request{
				Url:  fmt.Sprintf(listurl, 1, now.Month(), now.Day()),
				Rule: "列表",
				Header: http.Header{
					"Cookie": []string{cookie},
				},
			})
		},

		Trunk: map[string]*Rule{
			"列表": {
				ParseFunc: func(ctx *Context) {
					body := ctx.GetText()
					var result result
					err := json.Unmarshal([]byte(body), &result)
					if err != nil {
						fmt.Println(err)
						return
					}
					nextpage := ctx.GetTemp("page", 1).(int)
					if len(result.Stocks) > 0 {
						nextpage++
						ctx.AddQueue(&request.Request{
							Url:  fmt.Sprintf(listurl, nextpage, now.Month(), now.Day()),
							Rule: "列表",
							Header: http.Header{
								"Cookie": []string{cookie},
							},
							Temp: map[string]interface{}{
								"page": nextpage,
							},
						})
					} else {
						return
					}
					for _, i := range result.Stocks {
						ctx.AddQueue(&request.Request{
							Url:  fmt.Sprintf(cwzburl, i.Symbol),
							Rule: "财务指标",
							Temp: map[string]interface{}{
								"symbol": i.Symbol,
								"name":   i.Name,
								"code":   i.Code,
							},
							Header: http.Header{
								"Cookie": []string{cookie},
							},
						})
					}
				},
			},
			"财务指标": {
				ItemFields: []string{
					"compcode",
					"reportdate",
					"basiceps",
					"epsdiluted",
					"epsweighted",
					"naps",
					"opercashpershare",
					"peropecashpershare",
					"netassgrowrate",
					"dilutedroe",
					"weightedroe",
					"mainbusincgrowrate",
					"netincgrowrate",
					"totassgrowrate",
					"salegrossprofitrto",
					"mainbusiincome",
					"mainbusiprofit",
					"totprofit",
					"netprofit",
					"totalassets",
					"totalliab",
					"totsharequi",
					"operrevenue",
					"invnetcashflow",
					"finnetcflow",
					"chgexchgchgs",
					"cashnetr",
					"cashequfinbal",
					"symbol",
					"name",
					"code",
				},
				ParseFunc: func(ctx *Context) {
					body := ctx.GetText()
					var result map[string]interface{}
					json.Unmarshal([]byte(body), &result)
					for _, i := range result["list"].([]interface{}) {
						data := i.(map[string]interface{})
						ctx.Output(map[int]interface{}{
							0:  data["compcode"],
							1:  data["reportdate"],
							2:  data["basiceps"],
							3:  data["epsdiluted"],
							4:  data["epsweighted"],
							5:  data["naps"],
							6:  data["opercashpershare"],
							7:  data["peropecashpershare"],
							8:  data["netassgrowrate"],
							9:  data["dilutedroe"],
							10: data["weightedroe"],
							11: data["mainbusincgrowrate"],
							12: data["netincgrowrate"],
							13: data["totassgrowrate"],
							14: data["salegrossprofitrto"],
							15: data["mainbusiincome"],
							16: data["mainbusiprofit"],
							17: data["totprofit"],
							18: data["netprofit"],
							19: data["totalassets"],
							20: data["totalliab"],
							21: data["totsharequi"],
							22: data["operrevenue"],
							23: data["invnetcashflow"],
							24: data["finnetcflow"],
							25: data["chgexchgchgs"],
							26: data["cashnetr"],
							27: data["cashequfinbal"],
							28: ctx.GetTemp("symbol", ""),
							29: ctx.GetTemp("name", ""),
							30: ctx.GetTemp("code", ""),
						})
					}
				},
			},
		},
	},
}

type result struct {
	Count  count   `json:"count"`
	Stocks []stock `json:"stocks"`
}

type count struct {
	Count int `json:"count"`
}

type stock struct {
	Symbol string `json:"symbol"`
	Code   string `json:"code"`
	Name   string `json:"name"`
}
