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
	zcfzb.Register()
}

var zcfzb = &Spider{
	Name:        "公司资产负债表",
	Description: "获取雪球公司资产负债表数据",
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
					Url:  fmt.Sprintf(zcfzburl, symbol),
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
					"reportdate",       //: "报表日期",
					"curfds",           //: "货币资金",
					"tradfinasset",     //: "交易性金融资产",
					"notesrece",        //: "应收票据",
					"accorece",         //: "应收账款",
					"prep",             //: "预付款项",
					"premrece",         //: "应收保费",
					"interece",         //: "应收利息",
					"dividrece",        //: "应收股利",
					"otherrece",        //: "其他应收款",
					"expotaxrebarece",  //: "应收出口退税",
					"subsrece",         //: "应收补贴款",
					"margrece",         //: "应收保证金",
					"intelrece",        //: "内部应收款",
					"inve",             //: "存货",
					"prepexpe",         //: "待摊费用",
					"unseg",            //: "待处理流动资产损益",
					"expinoncurrasset", //: "一年内到期的非流动资产",
					"othercurrasse",    //: "其他流动资产",
					"currassetitse",    //: "特殊处理本身不平流动资产",
					"currasseform",     //: "特殊处理格式不同流动资产",
					"totcurrasset",     //: "流动资产合计",
					"lendandloan",      //: "发放贷款及垫款",
					"avaisellasse",     //: "可供出售金融资产",
					"holdinvedue",      //: "持有至到期投资",
					"longrece",         //: "长期应收款",
					"equiinve",         //: "长期股权投资",
					"otherlonginve",    //: "其他长期投资",
					"inveprop",         //: "投资性房地产",
					"fixedasseimmo",    //: "固定资产原值",
					"accudepr",         //: "累计折旧",
					"fixedassenetw",    //: "固定资产净值",
					"fixedasseimpa",    //: "固定资产减值准备",
					"fixedassenet",     //: "固定资产净额",
					"consprog",         //: "在建工程",
					"engimate",         //: "工程物资",
					"fixedasseclea",    //: "固定资产清理",
					"prodasse",         //: "生产性生物资产",
					"comasse",          //: "公益性生物资产",
					"hydrasset",        //: "油气资产",
					"intaasset",        //: "无形资产",
					"deveexpe",         //: "开发支出",
					"goodwill",         //: "商誉",
					"logprepexpe",      //: "长期待摊费用",
					"tradshartrad",     //: "股权分置流通权",
					"defetaxasset",     //: "递延所得税资产",
					"othernoncasse",    //: "其他非流动资产",
					"noncasseitse",     //: "特殊处理本身不平非流动资产",
					"noncasseform",     //: "特殊处理格式不同非流动资产",
					"totalnoncassets",  //: "非流动资产合计",
					"totassetitse",     //: "特殊处理本身不平总资产",
					"totalasseform",    //: "特殊处理格式不同总资产",
					"totasset",         //: "资产总计",
					"shorttermborr",    //: "短期借款",
					"tradfinliab",      //: "交易性金融负债",
					"notespaya",        //: "应付票据",
					"accopaya",         //: "应付账款",
					"advapaym",         //: "预收款项",
					"copeworkersal",    //: "应付职工薪酬",
					"taxespaya",        //: "应交税费",
					"intepaya",         //: "应付利息",
					"divipaya",         //: "应付股利",
					"otherfeepaya",     //: "其他应交款",
					"margrequ",         //: "应付保证金",
					"intelpay",         //: "内部应付款",
					"otherpay",         //: "其他应付款",
					"accrexpe",         //: "预提费用",
					"expecurrliab",     //: "预计流动负债",
					"copewithreinrece", //: "应付分保账款",
					"inteticksett",     //: "国际票证结算",
					"dometicksett",     //: "国内票证结算",
					"defereve",         //: "一年内的递延收益",
					"shorttermbdspaya", //: "应付短期债券",
					"duenoncliab",      //: "一年内到期的非流动负债",
					"othercurreliabi",  //: "其他流动负债",
					"currliabitse",     //: "特殊处理本身不平流动负债",
					"currliabform",     //: "特殊处理格式不同流动负债",
					"totalcurrliab",    //: "流动负债合计",
					"longborr",         //: "长期借款",
					"bdspaya",          //: "应付债券",
					"longpaya",         //: "长期应付款",
					"specpaya",         //: "专项应付款",
					"expenoncliab",     //: "预计非流动负债",
					"longdefeinco",     //: "长期递延收益",
					"defeincotaxliab",  //: "递延所得税负债",
					"othernoncliabi",   //: "其他非流动负债",
					"longliabitse",     //: "特殊处理本身不平长期负债",
					"longliabform",     //: "特殊处理格式不同长期负债",
					"totalnoncliab",    //: "非流动负债合计",
					"totliabitse",      //: "特殊处理本身不平负债合计",
					"totliabform",      //: "特殊处理格式不同负债合计",
					"totliab",          //: "负债合计",
					"paidincapi",       //: "实收资本(或股本)",
					"capisurp",         //: "资本公积",
					"treastk",          //: "减：库存股",
					"specrese",         //: "专项储备",
					"rese",             //: "盈余公积",
					"generiskrese",     //: "一般风险准备",
					"unreinveloss",     //: "未确定的投资损失",
					"undiprof",         //: "未分配利润",
					"topaycashdivi",    //: "拟分配现金股利",
					"curtrandiff",      //: "外币报表折算差额",
					"sharrighitse",     //: "特殊处理本身不平股东权益",
					"sharrightform",    //: "特殊处理格式不同股东权益",
					"paresharrigh",     //: "归属于母公司股东权益合计",
					"minysharrigh",     //: "少数股东权益",
					"righaggritse",     //: "特殊处理本身不平股东权益",
					"rightaggrform",    //: "特殊处理格式不平股东权益",
					"righaggr",         //: "所有者权益(或股东权益)合计",
					"rightotitse",      //: "特殊处理本身不平负债及权益",
					"rightotform",      //: "特殊处理格式不同负债及权益",
					"totliabsharequi",  //: "负债和所有者权益",
					"settresedepo",     //: "结算备付金",
					"plac",             //: "拆出资金",
					"derifinaasset",    //: "衍生金融资产",
					"reinrece",         //: "应收分保账款",
					"reincontrese",     //: "应收分保合同准备金",
					"purcresaasset",    //: "买入返售金融资产",
					"cenbankborr",      //: "向中央银行借款",
					"deposit",          //: "吸收存款及同业存放",
					"fdsborr",          //: "拆入资金",
					"deriliab",         //: "衍生金融负债",
					"sellrepasse",      //: "卖出回购金融资产款",
					"copepoun",         //: "应付手续费及佣金",
					"insucontrese",     //: "保险合同准备金",
					"actitradsecu",     //: "代理买卖证券款",
					"actiundesecu",     //: "代理承销证券款"
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
							3:   item["reportdate"],       //: "报表日期",
							4:   item["curfds"],           //: "货币资金",
							5:   item["tradfinasset"],     //: "交易性金融资产",
							6:   item["notesrece"],        //: "应收票据",
							7:   item["accorece"],         //: "应收账款",
							8:   item["prep"],             //: "预付款项",
							9:   item["premrece"],         //: "应收保费",
							10:  item["interece"],         //: "应收利息",
							11:  item["dividrece"],        //: "应收股利",
							12:  item["otherrece"],        //: "其他应收款",
							13:  item["expotaxrebarece"],  //: "应收出口退税",
							14:  item["subsrece"],         //: "应收补贴款",
							15:  item["margrece"],         //: "应收保证金",
							16:  item["intelrece"],        //: "内部应收款",
							17:  item["inve"],             //: "存货",
							18:  item["prepexpe"],         //: "待摊费用",
							19:  item["unseg"],            //: "待处理流动资产损益",
							20:  item["expinoncurrasset"], //: "一年内到期的非流动资产",
							21:  item["othercurrasse"],    //: "其他流动资产",
							22:  item["currassetitse"],    //: "特殊处理本身不平流动资产",
							23:  item["currasseform"],     //: "特殊处理格式不同流动资产",
							24:  item["totcurrasset"],     //: "流动资产合计",
							25:  item["lendandloan"],      //: "发放贷款及垫款",
							26:  item["avaisellasse"],     //: "可供出售金融资产",
							27:  item["holdinvedue"],      //: "持有至到期投资",
							28:  item["longrece"],         //: "长期应收款",
							29:  item["equiinve"],         //: "长期股权投资",
							30:  item["otherlonginve"],    //: "其他长期投资",
							31:  item["inveprop"],         //: "投资性房地产",
							32:  item["fixedasseimmo"],    //: "固定资产原值",
							33:  item["accudepr"],         //: "累计折旧",
							34:  item["fixedassenetw"],    //: "固定资产净值",
							35:  item["fixedasseimpa"],    //: "固定资产减值准备",
							36:  item["fixedassenet"],     //: "固定资产净额",
							37:  item["consprog"],         //: "在建工程",
							38:  item["engimate"],         //: "工程物资",
							39:  item["fixedasseclea"],    //: "固定资产清理",
							40:  item["prodasse"],         //: "生产性生物资产",
							41:  item["comasse"],          //: "公益性生物资产",
							42:  item["hydrasset"],        //: "油气资产",
							43:  item["intaasset"],        //: "无形资产",
							44:  item["deveexpe"],         //: "开发支出",
							45:  item["goodwill"],         //: "商誉",
							46:  item["logprepexpe"],      //: "长期待摊费用",
							47:  item["tradshartrad"],     //: "股权分置流通权",
							48:  item["defetaxasset"],     //: "递延所得税资产",
							49:  item["othernoncasse"],    //: "其他非流动资产",
							50:  item["noncasseitse"],     //: "特殊处理本身不平非流动资产",
							51:  item["noncasseform"],     //: "特殊处理格式不同非流动资产",
							52:  item["totalnoncassets"],  //: "非流动资产合计",
							53:  item["totassetitse"],     //: "特殊处理本身不平总资产",
							54:  item["totalasseform"],    //: "特殊处理格式不同总资产",
							55:  item["totasset"],         //: "资产总计",
							56:  item["shorttermborr"],    //: "短期借款",
							57:  item["tradfinliab"],      //: "交易性金融负债",
							58:  item["notespaya"],        //: "应付票据",
							59:  item["accopaya"],         //: "应付账款",
							60:  item["advapaym"],         //: "预收款项",
							61:  item["copeworkersal"],    //: "应付职工薪酬",
							62:  item["taxespaya"],        //: "应交税费",
							63:  item["intepaya"],         //: "应付利息",
							64:  item["divipaya"],         //: "应付股利",
							65:  item["otherfeepaya"],     //: "其他应交款",
							66:  item["margrequ"],         //: "应付保证金",
							67:  item["intelpay"],         //: "内部应付款",
							68:  item["otherpay"],         //: "其他应付款",
							69:  item["accrexpe"],         //: "预提费用",
							70:  item["expecurrliab"],     //: "预计流动负债",
							71:  item["copewithreinrece"], //: "应付分保账款",
							72:  item["inteticksett"],     //: "国际票证结算",
							73:  item["dometicksett"],     //: "国内票证结算",
							74:  item["defereve"],         //: "一年内的递延收益",
							75:  item["shorttermbdspaya"], //: "应付短期债券",
							76:  item["duenoncliab"],      //: "一年内到期的非流动负债",
							77:  item["othercurreliabi"],  //: "其他流动负债",
							78:  item["currliabitse"],     //: "特殊处理本身不平流动负债",
							79:  item["currliabform"],     //: "特殊处理格式不同流动负债",
							80:  item["totalcurrliab"],    //: "流动负债合计",
							81:  item["longborr"],         //: "长期借款",
							82:  item["bdspaya"],          //: "应付债券",
							83:  item["longpaya"],         //: "长期应付款",
							84:  item["specpaya"],         //: "专项应付款",
							85:  item["expenoncliab"],     //: "预计非流动负债",
							86:  item["longdefeinco"],     //: "长期递延收益",
							87:  item["defeincotaxliab"],  //: "递延所得税负债",
							88:  item["othernoncliabi"],   //: "其他非流动负债",
							89:  item["longliabitse"],     //: "特殊处理本身不平长期负债",
							90:  item["longliabform"],     //: "特殊处理格式不同长期负债",
							91:  item["totalnoncliab"],    //: "非流动负债合计",
							92:  item["totliabitse"],      //: "特殊处理本身不平负债合计",
							93:  item["totliabform"],      //: "特殊处理格式不同负债合计",
							94:  item["totliab"],          //: "负债合计",
							95:  item["paidincapi"],       //: "实收资本(或股本)",
							96:  item["capisurp"],         //: "资本公积",
							97:  item["treastk"],          //: "减：库存股",
							98:  item["specrese"],         //: "专项储备",
							99:  item["rese"],             //: "盈余公积",
							100: item["generiskrese"],     //: "一般风险准备",
							101: item["unreinveloss"],     //: "未确定的投资损失",
							102: item["undiprof"],         //: "未分配利润",
							103: item["topaycashdivi"],    //: "拟分配现金股利",
							104: item["curtrandiff"],      //: "外币报表折算差额",
							105: item["sharrighitse"],     //: "特殊处理本身不平股东权益",
							106: item["sharrightform"],    //: "特殊处理格式不同股东权益",
							107: item["paresharrigh"],     //: "归属于母公司股东权益合计",
							108: item["minysharrigh"],     //: "少数股东权益",
							109: item["righaggritse"],     //: "特殊处理本身不平股东权益",
							110: item["rightaggrform"],    //: "特殊处理格式不平股东权益",
							111: item["righaggr"],         //: "所有者权益(或股东权益)合计",
							112: item["rightotitse"],      //: "特殊处理本身不平负债及权益",
							113: item["rightotform"],      //: "特殊处理格式不同负债及权益",
							114: item["totliabsharequi"],  //: "负债和所有者权益",
							115: item["settresedepo"],     //: "结算备付金",
							116: item["plac"],             //: "拆出资金",
							117: item["derifinaasset"],    //: "衍生金融资产",
							118: item["reinrece"],         //: "应收分保账款",
							119: item["reincontrese"],     //: "应收分保合同准备金",
							120: item["purcresaasset"],    //: "买入返售金融资产",
							121: item["cenbankborr"],      //: "向中央银行借款",
							122: item["deposit"],          //: "吸收存款及同业存放",
							123: item["fdsborr"],          //: "拆入资金",
							124: item["deriliab"],         //: "衍生金融负债",
							125: item["sellrepasse"],      //: "卖出回购金融资产款",
							126: item["copepoun"],         //: "应付手续费及佣金",
							127: item["insucontrese"],     //: "保险合同准备金",
							128: item["actitradsecu"],     //: "代理买卖证券款",
							129: item["actiundesecu"],     //: "代理承销证券款"
						})
					}
				},
			},
		},
	},
}
