package stock

const (
	// 所有股票基本信息列表
	ALL_STOCK_BASICS_FILE = "http://218.244.146.57/static/all.csv"
	// 个股历史数据，日，周，月。参考 K_TYPE
	DAY_PRICE_URL = "http://api.finance.ifeng.com/%s/?code=%s&type=last"

	// 报表数据
	REPORT_URL = "http://vip.stock.finance.sina.com.cn/q/go.php/vFinanceAnalyze/kind/mainindex/index.phtml?s_i=&s_a=&s_c=&reportdate=%s&quarter=%s&p=%s&num=%s"
)

var (
	K_TYPE = map[string]string{"D": "akdaily", "W": "akweekly", "M": "akmonthly"}
)
