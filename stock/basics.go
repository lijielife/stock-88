package stock

import (
	"encoding/csv"
	"net/http"

	"github.com/axgle/mahonia"
)

type (
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
		code             string
		name             string
		industry         string
		area             string
		pe               string
		outstanding      string
		totals           string
		totalAssets      string
		liquidAssets     string
		fixedAssets      string
		reserved         string
		reservedPerShare string
		eps              string
		bvps             string
		pb               string
		timeToMarket     string
	}
)

//List 获取股票列表及基本信息
func List() (data []Basics, err error) {
	rsp, err := http.Get(ALL_STOCK_BASICS_FILE)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	enc := mahonia.NewDecoder("gbk")
	r := csv.NewReader(enc.NewReader(rsp.Body))
	lines, err := r.ReadAll()
	if err != nil {
		return
	}
	data = make([]Basics, len(lines))
	for i, line := range lines {
		data[i] = Basics{
			code:             line[0],
			name:             line[1],
			industry:         line[2],
			area:             line[3],
			pe:               line[4],
			outstanding:      line[5],
			totals:           line[6],
			totalAssets:      line[7],
			liquidAssets:     line[8],
			fixedAssets:      line[9],
			reserved:         line[10],
			reservedPerShare: line[11],
			eps:              line[12],
			bvps:             line[13],
			pb:               line[14],
			timeToMarket:     line[15],
		}
	}
	return
}
