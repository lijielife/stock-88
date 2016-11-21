package spiders

// 基础包
import (
	"fmt"
	"net/http"

	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	//DOM解析
	//DOM解析
	// "github.com/henrylee2cn/pholcus/logs"               //信息输出
	. "github.com/henrylee2cn/pholcus/app/spider" //必需
	// . "github.com/henrylee2cn/pholcus/app/spider/common" //选用
	// net包
	// "net/http" //设置http.Header
	// "net/url"
	// 编码包
	// "encoding/xml"
	// "encoding/json"
	// 字符串处理包
	// "strconv"
	// 其他包
	// "fmt"
	// "math"
	// "time"
	"encoding/json"
)

func init() {
	hslist.Register()
}

var hslist = &Spider{
	Name:        "沪深列表",
	Description: "获取雪球沪深列表数据",
	// Pausetime:    300,
	// Keyin:        KEYIN,
	// Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:  fmt.Sprintf(listurl, 1),
				Rule: "雪球数据",
				Header: http.Header{
					"Cookie": []string{cookies},
				},
			})
		},
		Trunk: map[string]*Rule{
			"雪球数据": {
				ItemFields: []string{
					"code",
					"name",
					"symbol",
				},
				ParseFunc: func(ctx *Context) {
					text := ctx.GetText()
					page := ctx.GetTemp("page", 1).(int)
					var data map[string]interface{}
					err := json.Unmarshal([]byte(text), &data)
					if err != nil {
						panic(err)
					}
					tlist := data["stocks"].([]interface{})
					if len(tlist) > 0 {
						//
						for _, i := range tlist {
							item := i.(map[string]interface{})
							ctx.Output(map[int]interface{}{
								0: item["code"],
								1: item["name"],
								2: item["symbol"],
							})
						}

						page++
						ctx.AddQueue(&request.Request{
							Url:  fmt.Sprintf(listurl, page),
							Rule: "雪球数据",
							Header: http.Header{
								"Cookie": []string{cookies},
							},
							Temp: map[string]interface{}{
								"page": page,
							},
						})
					}
				},
			},
		},
	},
}
