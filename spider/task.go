package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
)

const (
	cookies  = "s=6v11gmzerx; xq_a_token=525dbd797ebd974a897c029c5738d43a562afd8a; xqat=525dbd797ebd974a897c029c5738d43a562afd8a; xq_r_token=c672fd7cb7f6e83a4267a18d8847c104bd37158d; xq_is_login=1; u=6234241459; xq_token_expire=Fri%20Dec%2009%202016%2021%3A03%3A15%20GMT%2B0800%20(CST); bid=cbbb8dbf096e52abc62fd8918a59e7d9_ivi2z4cx; snbim_minify=true; webp=1; __utmt=1; Hm_lvt_1db88642e346389874251b5a1eded6e3=1479128534; Hm_lpvt_1db88642e346389874251b5a1eded6e3=1479518640; __utma=1.33235178.1479129012.1479478640.1479518445.6; __utmb=1.9.9.1479518456882; __utmc=1; __utmz=1.1479129012.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)"
	listurl  = "https://xueqiu.com/stock/cata/stocklist.json?page=%d&size=90&order=desc&orderby=percent&type=11&_=1479520263018"
	gslrburl = "https://xueqiu.com/stock/f10/incstatement.json?symbol=%s&page=1&size=100&_=1479518640609"
	zcfzburl = "https://xueqiu.com/stock/f10/balsheet.json?symbol=%s&page=1&size=100&_=1479524795302"
	xjllburl = "https://xueqiu.com/stock/f10/cfstatement.json?symbol=%s&page=1&size=100&_=1479525698882"
)

var (
	header = http.Header{}
	stocks = []string{}
	engine *xorm.Engine
)

//Start 启动爬虫任务
func Start() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/finance?charset=utf8")
	if err != nil {
		log.Panic(err)
	}
	header.Add("Cookie", cookies)
	err = fillStocks(1)
	if err != nil {
		log.Panic(err)
	}
	for _, symbol := range stocks {
		zcfzb(symbol)
	}
}
func fillStocks(page int) (err error) {
	u := fmt.Sprintf(listurl, page)
	bts, err := httpget(u, header)
	if err != nil {
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(bts, &result)
	if err != nil {
		return
	}
	slist := result["stocks"].([]interface{})
	if len(slist) > 0 {
		for i := 0; i < len(slist); i++ {
			t := slist[i].(map[string]interface{})
			stocks = append(stocks, t["symbol"].(string))
		}
		page++
		return fillStocks(page)
	}
	return
}

func zcfzb(symbol string) {

}

func output() {

}

//Httpget 发起HTTP请求
func httpget(url string, header http.Header) (bts []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header = header
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	return ioutil.ReadAll(rsp.Body)
}
