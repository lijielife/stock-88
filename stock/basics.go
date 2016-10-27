package stock

import (
	"bytes"
	"encoding/csv"
	"fmt"

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
		Code             string
		Name             string
		Industry         string
		Area             string
		Pe               string
		Outstanding      string
		Totals           string
		TotalAssets      string
		LiquidAssets     string
		FixedAssets      string
		Reserved         string
		ReservedPerShare string
		Eps              string
		Bvps             string
		Pb               string
		TimeToMarket     string
	}

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
		Code        string
		Name        string
		Eps         string
		Eps_yoy     string
		Bvps        string
		Roe         string
		Epcf        string
		NetProfits  string
		ProfitsYoy  string
		Distrib     string
		Report_date string
	}
	//Profit 利润报表
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
	}
)

//List 获取列表及基本信息
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
			TimeToMarket:     line[15],
		}
	}
	return
}

//Report 获取报表基本信息
func (b *BasicsMux) Report(year, quarter int) (data []Profit, err error) {
	return nil, nil
}

//Profit 利润表
func (b *BasicsMux) Profit(year, quarter int) (data []Profit, err error) {
	pageno := 1
	data = []Profit{}
	var recall func(year, quarter, pageno int, data *[]Profit) error

	recall = func(year, quarter, pageno int, data *[]Profit) (err error) {
		url := fmt.Sprintf(b.reporturl, "profit", year, quarter, pageno)

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
		fmt.Println(doc.Find("#dataTable tbody tr").Length())
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
			}
			*data = append(*data, profit)
		})
		if pageno == 3 {
			return
		}
		if _, ok := doc.Find(".pages a").Last().Attr("onclick"); ok {
			pageno++
			recall(year, quarter, pageno, data)
		}
		return
	}

	recall(year, quarter, pageno, &data)
	return
}
