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
					"headerTitle1",     //: "流动资产",
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
					"headerTitle2",     //: "非流动资产",
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
					"headerTitle3",     //: "流动负债",
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
					"headerTitle4",     //: "非流动负债",
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
					"headerTitle5",     //: "所有者权益",
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
							4:   item["headerTitle1"],     //: "流动资产",
							5:   item["curfds"],           //: "货币资金",
							6:   item["tradfinasset"],     //: "交易性金融资产",
							7:   item["notesrece"],        //: "应收票据",
							8:   item["accorece"],         //: "应收账款",
							9:   item["prep"],             //: "预付款项",
							10:  item["premrece"],         //: "应收保费",
							11:  item["interece"],         //: "应收利息",
							12:  item["dividrece"],        //: "应收股利",
							13:  item["otherrece"],        //: "其他应收款",
							14:  item["expotaxrebarece"],  //: "应收出口退税",
							15:  item["subsrece"],         //: "应收补贴款",
							16:  item["margrece"],         //: "应收保证金",
							17:  item["intelrece"],        //: "内部应收款",
							18:  item["inve"],             //: "存货",
							19:  item["prepexpe"],         //: "待摊费用",
							20:  item["unseg"],            //: "待处理流动资产损益",
							21:  item["expinoncurrasset"], //: "一年内到期的非流动资产",
							22:  item["othercurrasse"],    //: "其他流动资产",
							23:  item["currassetitse"],    //: "特殊处理本身不平流动资产",
							24:  item["currasseform"],     //: "特殊处理格式不同流动资产",
							25:  item["totcurrasset"],     //: "流动资产合计",
							26:  item["headerTitle2"],     //: "非流动资产",
							27:  item["lendandloan"],      //: "发放贷款及垫款",
							28:  item["avaisellasse"],     //: "可供出售金融资产",
							29:  item["holdinvedue"],      //: "持有至到期投资",
							30:  item["longrece"],         //: "长期应收款",
							31:  item["equiinve"],         //: "长期股权投资",
							32:  item["otherlonginve"],    //: "其他长期投资",
							33:  item["inveprop"],         //: "投资性房地产",
							34:  item["fixedasseimmo"],    //: "固定资产原值",
							35:  item["accudepr"],         //: "累计折旧",
							36:  item["fixedassenetw"],    //: "固定资产净值",
							37:  item["fixedasseimpa"],    //: "固定资产减值准备",
							38:  item["fixedassenet"],     //: "固定资产净额",
							39:  item["consprog"],         //: "在建工程",
							40:  item["engimate"],         //: "工程物资",
							41:  item["fixedasseclea"],    //: "固定资产清理",
							42:  item["prodasse"],         //: "生产性生物资产",
							43:  item["comasse"],          //: "公益性生物资产",
							44:  item["hydrasset"],        //: "油气资产",
							45:  item["intaasset"],        //: "无形资产",
							46:  item["deveexpe"],         //: "开发支出",
							47:  item["goodwill"],         //: "商誉",
							48:  item["logprepexpe"],      //: "长期待摊费用",
							49:  item["tradshartrad"],     //: "股权分置流通权",
							50:  item["defetaxasset"],     //: "递延所得税资产",
							51:  item["othernoncasse"],    //: "其他非流动资产",
							52:  item["noncasseitse"],     //: "特殊处理本身不平非流动资产",
							53:  item["noncasseform"],     //: "特殊处理格式不同非流动资产",
							54:  item["totalnoncassets"],  //: "非流动资产合计",
							55:  item["totassetitse"],     //: "特殊处理本身不平总资产",
							56:  item["totalasseform"],    //: "特殊处理格式不同总资产",
							57:  item["totasset"],         //: "资产总计",
							58:  item["headerTitle3"],     //: "流动负债",
							59:  item["shorttermborr"],    //: "短期借款",
							60:  item["tradfinliab"],      //: "交易性金融负债",
							61:  item["notespaya"],        //: "应付票据",
							62:  item["accopaya"],         //: "应付账款",
							63:  item["advapaym"],         //: "预收款项",
							64:  item["copeworkersal"],    //: "应付职工薪酬",
							65:  item["taxespaya"],        //: "应交税费",
							66:  item["intepaya"],         //: "应付利息",
							67:  item["divipaya"],         //: "应付股利",
							68:  item["otherfeepaya"],     //: "其他应交款",
							69:  item["margrequ"],         //: "应付保证金",
							70:  item["intelpay"],         //: "内部应付款",
							71:  item["otherpay"],         //: "其他应付款",
							72:  item["accrexpe"],         //: "预提费用",
							73:  item["expecurrliab"],     //: "预计流动负债",
							74:  item["copewithreinrece"], //: "应付分保账款",
							75:  item["inteticksett"],     //: "国际票证结算",
							76:  item["dometicksett"],     //: "国内票证结算",
							77:  item["defereve"],         //: "一年内的递延收益",
							78:  item["shorttermbdspaya"], //: "应付短期债券",
							79:  item["duenoncliab"],      //: "一年内到期的非流动负债",
							80:  item["othercurreliabi"],  //: "其他流动负债",
							81:  item["currliabitse"],     //: "特殊处理本身不平流动负债",
							82:  item["currliabform"],     //: "特殊处理格式不同流动负债",
							83:  item["totalcurrliab"],    //: "流动负债合计",
							84:  item["headerTitle4"],     //: "非流动负债",
							85:  item["longborr"],         //: "长期借款",
							86:  item["bdspaya"],          //: "应付债券",
							87:  item["longpaya"],         //: "长期应付款",
							88:  item["specpaya"],         //: "专项应付款",
							89:  item["expenoncliab"],     //: "预计非流动负债",
							90:  item["longdefeinco"],     //: "长期递延收益",
							91:  item["defeincotaxliab"],  //: "递延所得税负债",
							92:  item["othernoncliabi"],   //: "其他非流动负债",
							93:  item["longliabitse"],     //: "特殊处理本身不平长期负债",
							94:  item["longliabform"],     //: "特殊处理格式不同长期负债",
							95:  item["totalnoncliab"],    //: "非流动负债合计",
							96:  item["totliabitse"],      //: "特殊处理本身不平负债合计",
							97:  item["totliabform"],      //: "特殊处理格式不同负债合计",
							98:  item["totliab"],          //: "负债合计",
							99:  item["headerTitle5"],     //: "所有者权益",
							100: item["paidincapi"],       //: "实收资本(或股本)",
							101: item["capisurp"],         //: "资本公积",
							102: item["treastk"],          //: "减：库存股",
							103: item["specrese"],         //: "专项储备",
							104: item["rese"],             //: "盈余公积",
							105: item["generiskrese"],     //: "一般风险准备",
							106: item["unreinveloss"],     //: "未确定的投资损失",
							107: item["undiprof"],         //: "未分配利润",
							108: item["topaycashdivi"],    //: "拟分配现金股利",
							109: item["curtrandiff"],      //: "外币报表折算差额",
							110: item["sharrighitse"],     //: "特殊处理本身不平股东权益",
							111: item["sharrightform"],    //: "特殊处理格式不同股东权益",
							112: item["paresharrigh"],     //: "归属于母公司股东权益合计",
							113: item["minysharrigh"],     //: "少数股东权益",
							114: item["righaggritse"],     //: "特殊处理本身不平股东权益",
							115: item["rightaggrform"],    //: "特殊处理格式不平股东权益",
							116: item["righaggr"],         //: "所有者权益(或股东权益)合计",
							117: item["rightotitse"],      //: "特殊处理本身不平负债及权益",
							118: item["rightotform"],      //: "特殊处理格式不同负债及权益",
							119: item["totliabsharequi"],  //: "负债和所有者权益",
							120: item["settresedepo"],     //: "结算备付金",
							121: item["plac"],             //: "拆出资金",
							122: item["derifinaasset"],    //: "衍生金融资产",
							123: item["reinrece"],         //: "应收分保账款",
							124: item["reincontrese"],     //: "应收分保合同准备金",
							125: item["purcresaasset"],    //: "买入返售金融资产",
							126: item["cenbankborr"],      //: "向中央银行借款",
							127: item["deposit"],          //: "吸收存款及同业存放",
							128: item["fdsborr"],          //: "拆入资金",
							129: item["deriliab"],         //: "衍生金融负债",
							130: item["sellrepasse"],      //: "卖出回购金融资产款",
							131: item["copepoun"],         //: "应付手续费及佣金",
							132: item["insucontrese"],     //: "保险合同准备金",
							133: item["actitradsecu"],     //: "代理买卖证券款",
							134: item["actiundesecu"],     //: "代理承销证券款"
						})
					}
				},
			},
		},
	},
}
