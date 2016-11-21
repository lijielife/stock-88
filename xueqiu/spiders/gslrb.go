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
	gslrb.Register()
}

var gslrb = &Spider{
	Name:        "公司利润表",
	Description: "获取雪球公司利润表数据",
	// Pausetime:    300,
	// Keyin:        KEYIN,
	// Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			for _, symbol := range stocks() {
				if symbol == "" {
					continue
				}

				ctx.AddQueue(&request.Request{
					Url:  fmt.Sprintf(gslrburl, symbol),
					Rule: "雪球数据",
					Header: http.Header{
						"Cookie": []string{cookies},
					},
					Temp: map[string]interface{}{
						"code":   symbol[2:8],
						"symbol": symbol,
					},
				})
			}
		},
		Trunk: map[string]*Rule{
			"雪球数据": {
				ItemFields: []string{
					"code",
					"symbol",
					"comptype",
					"enddate",           //: "报表日期",
					"biztotinco",        //: "一、营业总收入",
					"bizinco",           //: "营业收入",
					"inteinco",          //: "利息收入",
					"earnprem",          //: "已赚保费",
					"pouninco",          //: "手续费及佣金收入",
					"realsale",          //: "房地产销售收入",
					"otherbizinco",      //: "其他业务收入",
					"biztotincoitse",    //: "特殊处理本身不平营业总收入",
					"biztotincoform",    //: "特殊处理格式不同营业总收入",
					"biztotcost",        //: "二、营业总成本",
					"bizcost",           //: "营业成本",
					"inteexpe",          //: "利息支出",
					"pounexpe",          //: "手续费及佣金支出",
					"realsalecost",      //: "房地产销售成本",
					"deveexpe",          //: "研发费用",
					"surrgold",          //: "退保金",
					"compnetexpe",       //: "赔付支出净额",
					"contress",          //: "提取保险合同准备金净额",
					"polidiviexpe",      //: "保单红利支出",
					"reinexpe",          //: "分保费用",
					"otherbizcost",      //: "其他业务成本",
					"biztax",            //: "营业税金及附加",
					"salesexpe",         //: "销售费用",
					"manaexpe",          //: "管理费用",
					"finexpe",           //: "财务费用",
					"asseimpaloss",      //: "资产减值损失",
					"biztotcostitse",    //: "特殊处理本身不平营业总成本",
					"biztotcostform",    //: "特殊处理格式不同营业总成本",
					"valuechgloss",      //: "公允价值变动收益",
					"inveinco",          //: "投资收益",
					"assoinveprof",      //: "其中:对联营企业和合营企业的投资收益",
					"exchggain",         //: "汇兑收益",
					"futuloss",          //: "期货损益",
					"custinco",          //: "托管收益",
					"subsidyincome",     //: "补贴收入",
					"otherbizprof",      //: "其他业务利润",
					"bizprofitse",       //: "特殊处理本身不平营业利润",
					"operprofform",      //: "特殊处理格式不同营业利润",
					"perprofit",         //: "三、营业利润",
					"nonoreve",          //: "营业外收入",
					"nonoexpe",          //: "营业外支出",
					"noncassetsdisl",    //: "非流动资产处置损失",
					"proftotitse",       //: "特殊处理本身不平利润总额",
					"proftotform",       //: "特殊处理格式不同利润总额",
					"totprofit",         //: "利润总额",
					"incotaxexpe",       //: "所得税费用",
					"unreinveloss",      //: "未确认投资损失",
					"netpro1itse",       //: "特殊处理本身不平净利润",
					"netpro1form",       //: "特殊处理格式不同净利润",
					"netprofit",         //: "四、净利润",
					"parenetp",          //: "归属于母公司所有者的净利润",
					"mergeformnetprof",  //: "被合并方在合并前实现净利润",
					"minysharrigh",      //: "少数股东损益",
					"netpro2itse",       //: "特殊处理本身不平净利润",
					"netpro2form",       //: "特殊处理格式不同净利润",
					"basiceps",          //: "五、基本每股收益",
					"dilutedeps",        //: "稀释每股收益",
					"othercompinco",     //: "六、其他综合收益",
					"parecompinco",      //: "归属于母公司所有者的其他综合收益",
					"minysharinco",      //: "归属于少数股东的其他综合收益",
					"compincoamt",       //: "七、综合收益总额",
					"parecompincoamt",   //: "归属于母公司所有者的综合收益总额",
					"minysharincoamt",   //: "归属于少数股东的综合收益总额",
					"earlyundiprof",     //: "年初未分配利润",
					"undisprofredu",     //: "减少注册资本减少的未分配利润",
					"otherinto",         //: "其他转入",
					"otherdistprof",     //: "可分配利润",
					"legalsurp",         //: "提取法定盈余公积",
					"welfare",           //: "提取公益金",
					"capitalreserve",    //: "提取资本公积金",
					"staffaward",        //: "提取职工奖福基金",
					"reservefund",       //: "提取储备基金",
					"developfund",       //: "提取企业发展基金",
					"profreturninvest",  //: "利润归还投资",
					"supplycurcap",      //: "补充流动资本",
					"avaidistshareprof", //: "可供股东分配的利润",
					"preferredskdiv",    //: "应付优先股股利",
					"freesurplu",        //: "提取任意公积",
					"dealwithdivi",      //: "应付普通股股利",
					"capitalizeddivi",   //: "转作资本股本的普通股股利",
					"undisprofit",       //: "未分配利润",
					"selldepartgain",    //: "出售处置部门或被投资单位所得收益",
					"natudisasterloss",  //: "自然灾害发生的损失",
					"accpolicychg",      //: "会计政策变更增加减少利润总额",
					"accestimatechg",    //: "会计估计变更增加减少利润总额",
					"debtrestruloss",    //: "债务重组损失",
					"othersupply",       //: "其他补充资料"
				},
				ParseFunc: func(ctx *Context) {
					text := ctx.GetText()
					var data map[string]interface{}
					err := json.Unmarshal([]byte(text), &data)
					if err != nil {
						panic(err)
					}
					tp, code, symbol := data["comptype"], ctx.GetTemp("code", ""), ctx.GetTemp("symbol", "")
					tlist := data["list"].([]interface{})
					for _, i := range tlist {
						item := i.(map[string]interface{})
						if tp.(float64) != 4 {
							item = map[string]interface{}{}
						}
						item["code"] = code
						item["symbol"] = symbol
						item["comptype"] = tp
						ctx.Output(map[int]interface{}{
							0:  item["code"],
							1:  item["symbol"],
							2:  item["comptype"],
							3:  item["enddate"],
							4:  item["biztotinco"],
							5:  item["bizinco"],
							6:  item["inteinco"],
							7:  item["earnprem"],
							8:  item["pouninco"],
							9:  item["realsale"],
							10: item["otherbizinco"],
							11: item["biztotincoitse"],
							12: item["biztotincoform"],
							13: item["biztotcost"],
							14: item["bizcost"],
							15: item["inteexpe"],
							16: item["pounexpe"],
							17: item["realsalecost"],
							18: item["deveexpe"],
							19: item["surrgold"],
							20: item["compnetexpe"],
							21: item["contress"],
							22: item["polidiviexpe"],
							23: item["reinexpe"],
							24: item["otherbizcost"],
							25: item["biztax"],
							26: item["salesexpe"],
							27: item["manaexpe"],
							28: item["finexpe"],
							29: item["asseimpaloss"],
							30: item["biztotcostitse"],
							31: item["biztotcostform"],
							32: item["valuechgloss"],
							33: item["inveinco"],
							34: item["assoinveprof"],
							35: item["exchggain"],
							36: item["futuloss"],
							37: item["custinco"],
							38: item["subsidyincome"],
							39: item["otherbizprof"],
							40: item["bizprofitse"],
							41: item["operprofform"],
							42: item["perprofit"],
							43: item["nonoreve"],
							44: item["nonoexpe"],
							45: item["noncassetsdisl"],
							46: item["proftotitse"],
							47: item["proftotform"],
							48: item["totprofit"],
							49: item["incotaxexpe"],
							50: item["unreinveloss"],
							51: item["netpro1itse"],
							52: item["netpro1form"],
							53: item["netprofit"],
							54: item["parenetp"],
							55: item["mergeformnetprof"],
							56: item["minysharrigh"],
							57: item["netpro2itse"],
							58: item["netpro2form"],
							59: item["basiceps"],
							60: item["dilutedeps"],
							61: item["othercompinco"],
							62: item["parecompinco"],
							63: item["minysharinco"],
							64: item["compincoamt"],
							65: item["parecompincoamt"],
							66: item["minysharincoamt"],
							67: item["earlyundiprof"],
							68: item["undisprofredu"],
							69: item["otherinto"],
							70: item["otherdistprof"],
							71: item["legalsurp"],
							72: item["welfare"],
							73: item["capitalreserve"],
							74: item["staffaward"],
							75: item["reservefund"],
							76: item["developfund"],
							77: item["profreturninvest"],
							78: item["supplycurcap"],
							79: item["avaidistshareprof"],
							80: item["preferredskdiv"],
							81: item["freesurplu"],
							82: item["dealwithdivi"],
							83: item["capitalizeddivi"],
							84: item["undisprofit"],
							85: item["selldepartgain"],
							86: item["natudisasterloss"],
							87: item["accpolicychg"],
							88: item["accestimatechg"],
							89: item["debtrestruloss"],
							90: item["othersupply"],
						})
					}
				},
			},
		},
	},
}
