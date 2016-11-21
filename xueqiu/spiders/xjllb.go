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
	xjllb.Register()
}

var xjllb = &Spider{
	Name:        "公司现金流量表",
	Description: "获取雪球公司现金流量表表数据",
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
					Url:  fmt.Sprintf(xjllburl, symbol),
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
					"enddate",            //: "报表日期",
					"headerTitle1",       //: "一、经营活动产生的现金流量",
					"laborgetcash",       //: "销售商品、提供劳务收到的现金",
					"deponetr",           //: "客户存款和同业存放款项净增加额",
					"bankloannetincr",    //: "向中央银行借款净增加额",
					"fininstnetr",        //: "向其他金融机构拆入资金净增加额",
					"inspremcash",        //: "收到原保险合同保费取得的现金",
					"insnetc",            //: "收到再保险业务现金净额",
					"savinetr",           //: "保户储金及投资款净增加额",
					"disptradnetincr",    //: "处置交易性金融资产净增加额",
					"charintecash",       //: "收取利息、手续费及佣金的现金",
					"fdsborrnetr",        //: "拆入资金净增加额",
					"repnetincr",         //: "回购业务资金净增加额",
					"taxrefd",            //: "收到的税费返还",
					"receotherbizcash",   //: "收到的其他与经营活动有关的现金",
					"bizinflitse",        //: "特殊处理本身不平经营流入",
					"bizinflform",        //: "特殊处理格式不同经营流入",
					"bizcashinfl",        //: "经营活动现金流入小计",
					"labopayc",           //: "购买商品、接受劳务支付的现金",
					"loansnetr",          //: "客户贷款及垫款净增加额",
					"tradepaymnetr",      //: "存放中央银行和同业款项净增加额",
					"paycompgold",        //: "支付原保险合同赔付款项的现金",
					"payintecash",        //: "支付利息、手续费及佣金的现金",
					"paydivicash",        //: "支付保单红利的现金",
					"payworkcash",        //: "支付给职工以及为职工支付的现金",
					"paytax",             //: "支付的各项税费",
					"payacticash",        //: "支付的其他与经营活动有关的现金",
					"bizoutfitse",        //: "特殊处理本身不平经营流出",
					"bizoutfform",        //: "特殊处理格式不同经营流出",
					"bizcashoutf",        //: "经营活动现金流出小计",
					"biznetitse",         //: "特殊处理本身不平经营净额",
					"biznetform",         //: "特殊处理格式不同经营净额",
					"mananetr",           //: "经营活动产生的现金流量净额",
					"headerTitle2",       //: "二、投资活动产生的现金流量",
					"withinvgetcash",     //: "收回投资所收到的现金",
					"inveretugetcash",    //: "取得投资收益收到的现金",
					"fixedassetnetc",     //: "处置固定资产、无形资产和其他长期资产所回收的现金净额",
					"subsnetc",           //: "处置子公司及其他营业单位收到的现金净额",
					"receinvcash",        //: "收到的其他与投资活动有关的现金",
					"reducashpled",       //: "减少质押和定期存款所收到的现金",
					"invinflitse",        //: "特殊处理本身不平投资流入",
					"invinffrom",         //: "特殊处理格式不同投资流入",
					"invcashinfl",        //: "投资活动现金流入小计",
					"acquassetcash",      //: "购建固定资产、无形资产和其他长期资产所支付的现金",
					"invpayc",            //: "投资所支付的现金",
					"loannetr",           //: "质押贷款净增加额",
					"subspaynetcash",     //: "取得子公司及其他营业单位支付的现金净额",
					"payinvecash",        //: "支付的其他与投资活动有关的现金",
					"incrcashpled",       //: "增加质押和定期存款所支付的现金",
					"invoutfitse",        //: "特殊处理本身不平投资流出",
					"invoutfform",        //: "特殊处理格式不同投资流出",
					"invcashoutf",        //: "投资活动现金流出小计",
					"netinvitse",         //: "特殊处理本身不平投资净额",
					"netinvform",         //: "特殊处理格式不同投资净额",
					"invnetcashflow",     //: "投资活动产生的现金流量净额",
					"headerTitle3",       //: "三、筹资活动产生的现金流量",
					"invrececash",        //: "吸收投资收到的现金",
					"subsrececash",       //: "其中：子公司吸收少数股东投资收到的现金",
					"recefromloan",       //: "取得借款收到的现金",
					"issbdrececash",      //: "发行债券收到的现金",
					"recefincash",        //: "收到其他与筹资活动有关的现金",
					"fininflitse",        //: "特殊处理本身不平筹资流入",
					"fininflform",        //: "特殊处理格式不同筹资流入",
					"fincashinfl",        //: "筹资活动现金流入小计",
					"debtpaycash",        //: "偿还债务支付的现金",
					"diviprofpaycash",    //: "分配股利、利润或偿付利息所支付的现金",
					"subspaydivid",       //: "其中：子公司支付给少数股东的股利，利润",
					"finrelacash",        //: "支付其他与筹资活动有关的现金",
					"finoutfitse",        //: "特殊处理本身不平筹资流出",
					"finoutfform",        //: "特殊处理格式不同筹资流出",
					"fincashoutf",        //: "筹资活动现金流出小计",
					"finnetitse",         //: "特殊处理本身不平筹资净额",
					"finnetform",         //: "特殊处理格式不同筹资净额",
					"finnetcflow",        //: "筹资活动产生的现金流量净",
					"chgexchgchgs",       //: "四、汇率变动对现金及现金等价物的影响",
					"netcashitse",        //: "特殊处理本身不平现金净额",
					"netcashform",        //: "特殊处理格式不同现金净额",
					"cashnetr",           //: "五、现金及现金等价物净增加额",
					"inicashbala",        //: "期初现金及现金等价物余额",
					"cashfinalitse",      //: "特殊处理本身不平现金期末",
					"cashfinalform",      //: "特殊处理格式不同现金期末",
					"finalcashbala",      //: "六、期末现金及现金等价物余额",
					"netprofit",          //: "净利润",
					"minysharrigh",       //: "少数股东权益",
					"unreinveloss",       //: "未确认的投资损失",
					"asseimpa",           //: "资产减值准备",
					"assedepr",           //: "固定资产折旧、油气资产折耗、生产性物资折旧",
					"realestadep",        //: "投资性房地产折旧、摊销",
					"intaasseamor",       //: "无形资产摊销",
					"longdefeexpenamor",  //: "长期待摊费用摊销",
					"prepexpedecr",       //: "待摊费用的减少",
					"accrexpeincr",       //: "预提费用的增加",
					"dispfixedassetloss", //: "处置固定资产、无形资产和其他长期资产的损失",
					"fixedassescraloss",  //: "固定资产报废损失",
					"valuechgloss",       //: "公允价值变动损失",
					"defeincoincr",       //: "递延收益增加（减：减少）",
					"estidebts",          //: "预计负债",
					"finexpe",            //: "财务费用",
					"inveloss",           //: "投资损失",
					"defetaxassetdecr",   //: "递延所得税资产减少",
					"defetaxliabincr",    //: "递延所得税负债增加",
					"inveredu",           //: "存货的减少",
					"receredu",           //: "经营性应收项目的减少",
					"payaincr",           //: "经营性应付项目的增加",
					"unseparachg",        //: "已完工尚未结算款的减少(减:增加)",
					"unfiparachg",        //: "已结算尚未完工款的增加(减:减少)",
					"other",              //: "其他",
					"biznetscheitse",     //: "特殊处理本身不平经营净额附表",
					"biznetscheform",     //: "特殊处理格式不同经营净额附表",
					"biznetcflow",        //: "经营活动产生现金流量净额",
					"debtintocapi",       //: "债务转为资本",
					"expiconvbd",         //: "一年内到期的可转换公司债券",
					"finfixedasset",      //: "融资租入固定资产",
					"cashfinalbala",      //: "现金的期末余额",
					"cashopenbala",       //: "现金的期初余额",
					"equfinalbala",       //: "现金等价物的期末余额",
					"equopenbala",        //: "现金等价物的期初余额",
					"netcashscheitse",    //: "特殊处理本身不平现金净额附表",
					"netcashscheform",    //: "特殊处理格式不同现金净额附表",
					"cashneti",           //: "现金及现金等价物的净增加额 "
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
							0:   item["code"],
							1:   item["symbol"],
							2:   item["comptype"],
							3:   item["enddate"],            //: "报表日期",
							4:   item["headerTitle1"],       //: "一、经营活动产生的现金流量",
							5:   item["laborgetcash"],       //: "销售商品、提供劳务收到的现金",
							6:   item["deponetr"],           //: "客户存款和同业存放款项净增加额",
							7:   item["bankloannetincr"],    //: "向中央银行借款净增加额",
							8:   item["fininstnetr"],        //: "向其他金融机构拆入资金净增加额",
							9:   item["inspremcash"],        //: "收到原保险合同保费取得的现金",
							10:  item["insnetc"],            //: "收到再保险业务现金净额",
							11:  item["savinetr"],           //: "保户储金及投资款净增加额",
							12:  item["disptradnetincr"],    //: "处置交易性金融资产净增加额",
							13:  item["charintecash"],       //: "收取利息、手续费及佣金的现金",
							14:  item["fdsborrnetr"],        //: "拆入资金净增加额",
							15:  item["repnetincr"],         //: "回购业务资金净增加额",
							16:  item["taxrefd"],            //: "收到的税费返还",
							17:  item["receotherbizcash"],   //: "收到的其他与经营活动有关的现金",
							18:  item["bizinflitse"],        //: "特殊处理本身不平经营流入",
							19:  item["bizinflform"],        //: "特殊处理格式不同经营流入",
							20:  item["bizcashinfl"],        //: "经营活动现金流入小计",
							21:  item["labopayc"],           //: "购买商品、接受劳务支付的现金",
							22:  item["loansnetr"],          //: "客户贷款及垫款净增加额",
							23:  item["tradepaymnetr"],      //: "存放中央银行和同业款项净增加额",
							24:  item["paycompgold"],        //: "支付原保险合同赔付款项的现金",
							25:  item["payintecash"],        //: "支付利息、手续费及佣金的现金",
							26:  item["paydivicash"],        //: "支付保单红利的现金",
							27:  item["payworkcash"],        //: "支付给职工以及为职工支付的现金",
							28:  item["paytax"],             //: "支付的各项税费",
							29:  item["payacticash"],        //: "支付的其他与经营活动有关的现金",
							30:  item["bizoutfitse"],        //: "特殊处理本身不平经营流出",
							31:  item["bizoutfform"],        //: "特殊处理格式不同经营流出",
							32:  item["bizcashoutf"],        //: "经营活动现金流出小计",
							33:  item["biznetitse"],         //: "特殊处理本身不平经营净额",
							34:  item["biznetform"],         //: "特殊处理格式不同经营净额",
							35:  item["mananetr"],           //: "经营活动产生的现金流量净额",
							36:  item["headerTitle2"],       //: "二、投资活动产生的现金流量",
							37:  item["withinvgetcash"],     //: "收回投资所收到的现金",
							38:  item["inveretugetcash"],    //: "取得投资收益收到的现金",
							39:  item["fixedassetnetc"],     //: "处置固定资产、无形资产和其他长期资产所回收的现金净额",
							40:  item["subsnetc"],           //: "处置子公司及其他营业单位收到的现金净额",
							41:  item["receinvcash"],        //: "收到的其他与投资活动有关的现金",
							42:  item["reducashpled"],       //: "减少质押和定期存款所收到的现金",
							43:  item["invinflitse"],        //: "特殊处理本身不平投资流入",
							44:  item["invinffrom"],         //: "特殊处理格式不同投资流入",
							45:  item["invcashinfl"],        //: "投资活动现金流入小计",
							46:  item["acquassetcash"],      //: "购建固定资产、无形资产和其他长期资产所支付的现金",
							47:  item["invpayc"],            //: "投资所支付的现金",
							48:  item["loannetr"],           //: "质押贷款净增加额",
							49:  item["subspaynetcash"],     //: "取得子公司及其他营业单位支付的现金净额",
							50:  item["payinvecash"],        //: "支付的其他与投资活动有关的现金",
							51:  item["incrcashpled"],       //: "增加质押和定期存款所支付的现金",
							52:  item["invoutfitse"],        //: "特殊处理本身不平投资流出",
							53:  item["invoutfform"],        //: "特殊处理格式不同投资流出",
							54:  item["invcashoutf"],        //: "投资活动现金流出小计",
							55:  item["netinvitse"],         //: "特殊处理本身不平投资净额",
							56:  item["netinvform"],         //: "特殊处理格式不同投资净额",
							57:  item["invnetcashflow"],     //: "投资活动产生的现金流量净额",
							58:  item["headerTitle3"],       //: "三、筹资活动产生的现金流量",
							59:  item["invrececash"],        //: "吸收投资收到的现金",
							60:  item["subsrececash"],       //: "其中：子公司吸收少数股东投资收到的现金",
							61:  item["recefromloan"],       //: "取得借款收到的现金",
							62:  item["issbdrececash"],      //: "发行债券收到的现金",
							63:  item["recefincash"],        //: "收到其他与筹资活动有关的现金",
							64:  item["fininflitse"],        //: "特殊处理本身不平筹资流入",
							65:  item["fininflform"],        //: "特殊处理格式不同筹资流入",
							66:  item["fincashinfl"],        //: "筹资活动现金流入小计",
							67:  item["debtpaycash"],        //: "偿还债务支付的现金",
							68:  item["diviprofpaycash"],    //: "分配股利、利润或偿付利息所支付的现金",
							69:  item["subspaydivid"],       //: "其中：子公司支付给少数股东的股利，利润",
							70:  item["finrelacash"],        //: "支付其他与筹资活动有关的现金",
							71:  item["finoutfitse"],        //: "特殊处理本身不平筹资流出",
							72:  item["finoutfform"],        //: "特殊处理格式不同筹资流出",
							73:  item["fincashoutf"],        //: "筹资活动现金流出小计",
							74:  item["finnetitse"],         //: "特殊处理本身不平筹资净额",
							75:  item["finnetform"],         //: "特殊处理格式不同筹资净额",
							76:  item["finnetcflow"],        //: "筹资活动产生的现金流量净",
							77:  item["chgexchgchgs"],       //: "四、汇率变动对现金及现金等价物的影响",
							78:  item["netcashitse"],        //: "特殊处理本身不平现金净额",
							79:  item["netcashform"],        //: "特殊处理格式不同现金净额",
							80:  item["cashnetr"],           //: "五、现金及现金等价物净增加额",
							81:  item["inicashbala"],        //: "期初现金及现金等价物余额",
							82:  item["cashfinalitse"],      //: "特殊处理本身不平现金期末",
							83:  item["cashfinalform"],      //: "特殊处理格式不同现金期末",
							84:  item["finalcashbala"],      //: "六、期末现金及现金等价物余额",
							85:  item["netprofit"],          //: "净利润",
							86:  item["minysharrigh"],       //: "少数股东权益",
							87:  item["unreinveloss"],       //: "未确认的投资损失",
							88:  item["asseimpa"],           //: "资产减值准备",
							89:  item["assedepr"],           //: "固定资产折旧、油气资产折耗、生产性物资折旧",
							90:  item["realestadep"],        //: "投资性房地产折旧、摊销",
							91:  item["intaasseamor"],       //: "无形资产摊销",
							92:  item["longdefeexpenamor"],  //: "长期待摊费用摊销",
							93:  item["prepexpedecr"],       //: "待摊费用的减少",
							94:  item["accrexpeincr"],       //: "预提费用的增加",
							95:  item["dispfixedassetloss"], //: "处置固定资产、无形资产和其他长期资产的损失",
							96:  item["fixedassescraloss"],  //: "固定资产报废损失",
							97:  item["valuechgloss"],       //: "公允价值变动损失",
							98:  item["defeincoincr"],       //: "递延收益增加（减：减少）",
							99:  item["estidebts"],          //: "预计负债",
							100: item["finexpe"],            //: "财务费用",
							101: item["inveloss"],           //: "投资损失",
							102: item["defetaxassetdecr"],   //: "递延所得税资产减少",
							103: item["defetaxliabincr"],    //: "递延所得税负债增加",
							104: item["inveredu"],           //: "存货的减少",
							105: item["receredu"],           //: "经营性应收项目的减少",
							106: item["payaincr"],           //: "经营性应付项目的增加",
							107: item["unseparachg"],        //: "已完工尚未结算款的减少(减:增加)",
							108: item["unfiparachg"],        //: "已结算尚未完工款的增加(减:减少)",
							109: item["other"],              //: "其他",
							110: item["biznetscheitse"],     //: "特殊处理本身不平经营净额附表",
							111: item["biznetscheform"],     //: "特殊处理格式不同经营净额附表",
							112: item["biznetcflow"],        //: "经营活动产生现金流量净额",
							113: item["debtintocapi"],       //: "债务转为资本",
							114: item["expiconvbd"],         //: "一年内到期的可转换公司债券",
							115: item["finfixedasset"],      //: "融资租入固定资产",
							116: item["cashfinalbala"],      //: "现金的期末余额",
							117: item["cashopenbala"],       //: "现金的期初余额",
							118: item["equfinalbala"],       //: "现金等价物的期末余额",
							119: item["equopenbala"],        //: "现金等价物的期初余额",
							120: item["netcashscheitse"],    //: "特殊处理本身不平现金净额附表",
							121: item["netcashscheform"],    //: "特殊处理格式不同现金净额附表",
							122: item["cashneti"],           //: "现金及现金等价物的净增加额 "
						})
					}
				},
			},
		},
	},
}
