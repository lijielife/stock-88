package stock

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

var (
	//Basicser 基本信息获取器
	Basicser = NewBasicser()
)

//NewBasicser create a Basicser
func NewBasicser() BasicsMux {
	return BasicsMux{
		listurl:   "http://218.244.146.57/static/all.csv",
		reporturl: "http://vip.stock.finance.sina.com.cn/q/go.php/vFinanceAnalyze/kind/%s/index.phtml?s_i=&s_a=&s_c=&reportdate=%d&quarter=%d&p=%d&num=60",
	}
}

type (
	//BasicsMux 基本面信息获取器
	BasicsMux struct {
		reporturl string
		listurl   string
	}
	//DataProc 数据处理类型
	DataProc func(*goquery.Document)
	//Basics 基本信息
	/*
		code,代码
		name,名称
		industry,细分行业
		area,地区
		pe,市盈率
		outstanding,流通股本
		totals,总股本(万)
		totalAssets,总资产(万)
		liquidAssets,流动资产
		fixedAssets,固定资产
		reserved,公积金
		reservedPerShare,每股公积金
		eps,每股收益
		bvps,每股净资
		pb,市净率
		timeToMarket,上市日期
	*/
	Basics struct {
		Code             string `xorm:"varchar(16) notnull pk"`
		Name             string `xorm:"varchar(16) notnull pk"`
		Industry         string `xorm:"varchar(16) notnull"`
		Area             string `xorm:"varchar(16) notnull"`
		Pe               string `xorm:"float"`
		Outstanding      string `xorm:"double"`
		Totals           string `xorm:"double"`
		TotalAssets      string `xorm:"double"`
		LiquidAssets     string `xorm:"double"`
		FixedAssets      string `xorm:"double"`
		Reserved         string `xorm:"double"`
		ReservedPerShare string `xorm:"float"`
		Eps              string `xorm:"float"`
		Bvps             string `xorm:"float"`
		Pb               string `xorm:"float"`
		TimeToMarket     string `xorm:"date"`
	}
	//Report 业绩数据
	/*
	   code,代码
	   name,名称
	   eps,每股收益
	   eps_yoy,每股收益同比(%)
	   bvps,每股净资产
	   roe,净资产收益率(%)
	   epcf,每股现金流量(元)
	   net_profits,净利润(万元)
	   profits_yoy,净利润同比(%)
	   distrib,分配方案
	   report_date,发布日期
	*/
	Report struct {
		Code       string `xorm:"varchar(16) notnull pk"`
		Name       string `xorm:"varchar(16) notnull"`
		Eps        string `xorm:"float pk"`
		EpsYoy     string `xorm:"float"`
		Bvps       string `xorm:"float"`
		Roe        string `xorm:"float"`
		Epcf       string `xorm:"float"`
		NetProfits string `xorm:"float"`
		ProfitsYoy string `xorm:"double"`
		Distrib    string `xorm:"varchar(32)"`
		Year       int    `xorm:"pk"`
		Quarter    int    `xorm:"pk"`
	}
	//Profit 利润数据
	/*
		code,代码
		name,名称
		roe,净资产收益率(%)
		net_profit_ratio,净利率(%)
		gross_profit_rate,毛利率(%)
		net_profits,净利润(万元)
		eps,每股收益
		business_income,营业收入(百万元)
		bips,每股主营业务收入(元)
	*/
	Profit struct {
		Code            string
		Name            string
		Roe             string
		NetProfitRatio  string
		GrossProfitRate string
		NetProfits      string
		Eps             string
		BusinessIncome  string
		Bips            string
		Year            int
		Quarter         int
	}
	//Operation 运营能力数据
	/*
		code,代码
		name,名称
		arturnover,应收账款周转率(次)
		arturndays,应收账款周转天数(天)
		inventory_turnover,存货周转率(次)
		inventory_days,存货周转天数(天)
		currentasset_turnover,流动资产周转率(次)
		currentasset_days,流动资产周转天数(天)
	*/
	Operation struct {
		Code                 string
		Name                 string
		Arturnover           string
		Arturndays           string
		InventoryTurnover    string
		InventoryDays        string
		CurrentassetTurnover string
		CurrentassetDays     string
		Year                 int
		Quarter              int
	}
	//Grow 成长能力数据
	/*
		code,代码
		name,名称
		mbrg,主营业务收入增长率(%)
		nprg,净利润增长率(%)
		nav,净资产增长率
		targ,总资产增长率
		epsg,每股收益增长率
		seg,股东权益增长率
	*/
	Grow struct {
		Code    string
		Name    string
		Mbrg    string
		Nprg    string
		Nav     string
		Targ    string
		Epsg    string
		Seg     string
		Year    int
		Quarter int
	}
	//Debtpay  偿债能力数据
	/*
		code,代码
		name,名称
		currentratio,流动比率
		quickratio,速动比率
		cashratio,现金比率
		icratio,利息支付倍数
		sheqratio,股东权益比率
		adratio,股东权益增长率
	*/
	Debtpay struct {
		Code         string
		Name         string
		Currentratio string
		Quickratio   string
		Cashratio    string
		Icratio      string
		Sheqratio    string
		Adratio      string
		Year         int
		Quarter      int
	}

	//Cashflow 现金流量数据
	/*
		code,代码
		name,名称
		cf_sales,经营现金净流量对销售收入比率
		rateofreturn,资产的经营现金流量回报率
		cf_nm,经营现金净流量与净利润的比率
		cf_liabilities,经营现金净流量对负债比率
		cashflowratio,现金流量比率
	*/
	Cashflow struct {
		Code          string
		Name          string
		CfSales       string
		Rateofreturn  string
		CfNm          string
		CfLiabilities string
		Cashflowratio string
		Year          int
		Quarter       int
	}
)

//List 获取沪深上市公司基本情况
func (b *BasicsMux) List() (data []Basics, err error) {
	bts, err := httpget(b.listurl)
	if err != nil {
		return
	}
	enc := mahonia.NewDecoder("gbk")
	r := csv.NewReader(enc.NewReader(bytes.NewReader(bts)))
	lines, err := r.ReadAll()
	if err != nil {
		return
	}
	data = make([]Basics, len(lines))
	for i, line := range lines {

		data[i] = Basics{
			Code:             line[0],
			Name:             line[1],
			Industry:         line[2],
			Area:             line[3],
			Pe:               line[4],
			Outstanding:      line[5],
			Totals:           line[6],
			TotalAssets:      line[7],
			LiquidAssets:     line[8],
			FixedAssets:      line[9],
			Reserved:         line[10],
			ReservedPerShare: line[11],
			Eps:              line[12],
			Bvps:             line[13],
			Pb:               line[14],
			TimeToMarket:     b.parseDate(line[15]),
		}
	}
	return
}

//Report 获取业绩报表数据
func (b *BasicsMux) Report(year, quarter int) ([]Report, error) {
	data := []Report{}
	err := b.report(year, quarter, "report", func(doc *goquery.Document) {
		//解析HTML
		doc.Find("#dataTable tbody tr").Each(func(i int, q *goquery.Selection) {
			tds := q.Find("td")
			report := Report{
				Code:       tds.Eq(0).Find("a").Text(),
				Name:       tds.Eq(1).Find("a").Text(),
				Eps:        b.parseNumber(tds.Eq(2).Text()),
				EpsYoy:     b.parseNumber(tds.Eq(3).Text()),
				Bvps:       b.parseNumber(tds.Eq(4).Text()),
				Roe:        b.parseNumber(tds.Eq(5).Text()),
				Epcf:       b.parseNumber(tds.Eq(6).Text()),
				NetProfits: b.parseNumber(tds.Eq(7).Text()),
				ProfitsYoy: b.parseNumber(tds.Eq(8).Text()),
				Distrib:    b.parseNumber(tds.Eq(9).Text()),
				Year:       year,
				Quarter:    quarter,
			}
			data = append(data, report)
		})
	})
	return data, err
}

//Profit 获取盈利能力数据
func (b *BasicsMux) Profit(year, quarter int) ([]Profit, error) {
	data := []Profit{}
	err := b.report(year, quarter, "profit", func(doc *goquery.Document) {
		//解析HTML
		doc.Find("#dataTable tbody tr").Each(func(i int, q *goquery.Selection) {
			profit := Profit{
				Code:            q.Find("td").Eq(0).Find("a").Text(),
				Name:            q.Find("td").Eq(1).Find("a").Text(),
				Roe:             q.Find("td").Eq(2).Text(),
				NetProfitRatio:  q.Find("td").Eq(3).Text(),
				GrossProfitRate: q.Find("td").Eq(4).Text(),
				NetProfits:      q.Find("td").Eq(5).Text(),
				Eps:             q.Find("td").Eq(6).Text(),
				BusinessIncome:  q.Find("td").Eq(7).Text(),
				Bips:            q.Find("td").Eq(8).Text(),
				Year:            year,
				Quarter:         quarter,
			}
			data = append(data, profit)
		})
	})
	return data, err
}

//Operation 获取运营能力数据
func (b *BasicsMux) Operation(year, quarter int) ([]Operation, error) {
	data := []Operation{}
	err := b.report(year, quarter, "operation", func(doc *goquery.Document) {
		//解析HTML
		doc.Find("#dataTable tbody tr").Each(func(i int, q *goquery.Selection) {
			operation := Operation{
				Code:                 q.Find("td").Eq(0).Find("a").Text(),
				Name:                 q.Find("td").Eq(1).Find("a").Text(),
				Arturnover:           q.Find("td").Eq(2).Text(),
				Arturndays:           q.Find("td").Eq(3).Text(),
				InventoryTurnover:    q.Find("td").Eq(4).Text(),
				InventoryDays:        q.Find("td").Eq(5).Text(),
				CurrentassetTurnover: q.Find("td").Eq(6).Text(),
				CurrentassetDays:     q.Find("td").Eq(7).Text(),
				Year:                 year,
				Quarter:              quarter,
			}
			data = append(data, operation)
		})
	})
	return data, err
}

//Growth 获取成长能力数据
func (b *BasicsMux) Growth(year, quarter int) ([]Grow, error) {
	data := []Grow{}
	err := b.report(year, quarter, "grow", func(doc *goquery.Document) {
		//解析HTML
		doc.Find("#dataTable tbody tr").Each(func(i int, q *goquery.Selection) {
			grow := Grow{
				Code:    q.Find("td").Eq(0).Find("a").Text(),
				Name:    q.Find("td").Eq(1).Find("a").Text(),
				Mbrg:    q.Find("td").Eq(2).Text(),
				Nprg:    q.Find("td").Eq(3).Text(),
				Nav:     q.Find("td").Eq(4).Text(),
				Targ:    q.Find("td").Eq(5).Text(),
				Epsg:    q.Find("td").Eq(6).Text(),
				Seg:     q.Find("td").Eq(7).Text(),
				Year:    year,
				Quarter: quarter,
			}
			data = append(data, grow)
		})
	})
	return data, err
}

// Debtpaying 获取偿债能力数据
func (b *BasicsMux) Debtpaying(year, quarter int) ([]Debtpay, error) {
	data := []Debtpay{}
	err := b.report(year, quarter, "grow", func(doc *goquery.Document) {
		//解析HTML
		doc.Find("#dataTable tbody tr").Each(func(i int, q *goquery.Selection) {
			debtpay := Debtpay{
				Code:         q.Find("td").Eq(0).Find("a").Text(),
				Name:         q.Find("td").Eq(1).Find("a").Text(),
				Currentratio: q.Find("td").Eq(2).Text(),
				Quickratio:   q.Find("td").Eq(3).Text(),
				Cashratio:    q.Find("td").Eq(4).Text(),
				Icratio:      q.Find("td").Eq(5).Text(),
				Sheqratio:    q.Find("td").Eq(6).Text(),
				Adratio:      q.Find("td").Eq(7).Text(),
				Year:         year,
				Quarter:      quarter,
			}
			data = append(data, debtpay)
		})
	})
	return data, err
}

// Cashflow 获取现金流量数据
func (b *BasicsMux) Cashflow(year, quarter int) ([]Cashflow, error) {
	data := []Cashflow{}
	err := b.report(year, quarter, "grow", func(doc *goquery.Document) {
		//解析HTML
		doc.Find("#dataTable tbody tr").Each(func(i int, q *goquery.Selection) {
			cashflow := Cashflow{
				Code:          q.Find("td").Eq(0).Find("a").Text(),
				Name:          q.Find("td").Eq(1).Find("a").Text(),
				CfSales:       q.Find("td").Eq(2).Text(),
				Rateofreturn:  q.Find("td").Eq(3).Text(),
				CfNm:          q.Find("td").Eq(4).Text(),
				CfLiabilities: q.Find("td").Eq(5).Text(),
				Cashflowratio: q.Find("td").Eq(6).Text(),
				Year:          year,
				Quarter:       quarter,
			}
			data = append(data, cashflow)
		})
	})
	return data, err
}

//递归获取每页数据
func (b *BasicsMux) report(year, quarter int, page string, proc DataProc) error {
	var recall func(year, quarter, pageno int) error

	recall = func(year, quarter, pageno int) (err error) {
		url := fmt.Sprintf(b.reporturl, page, year, quarter, pageno)

		bts, err := httpget(url)
		if err != nil {
			return err
		}
		enc := mahonia.NewDecoder("gbk")
		doc, err := goquery.NewDocumentFromReader(enc.NewReader(bytes.NewReader(bts)))
		if err != nil {
			return err
		}
		//解析HTML
		proc(doc)
		if _, ok := doc.Find(".pages a").Last().Attr("onclick"); ok {
			pageno++
			// time.Sleep(500 * time.Millisecond)
			recall(year, quarter, pageno)
		}
		return
	}
	return recall(year, quarter, 1)
}

func (b *BasicsMux) parseNumber(val string) string {
	if val == "--" || val == "" {
		val = "0"
	}
	return val
}

func (b *BasicsMux) parseDate(val string) string {
	if val == "0" || val == "" {
		val = time.Now().Format("2006-01-02")
	}
	return val
}
