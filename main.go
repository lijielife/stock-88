package main

import (
	"finance/task"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	task.Start()
	time.Sleep(30 * 60 * time.Second)
	// data, err := stock.List()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, item := range data {
	// 	fmt.Println(item)
	// }
	// engine, err := xorm.NewEngine("mysql", "root:123456@/finance?charset=utf8")
	// if err != nil {
	// 	panic(err)
	// }

	// type (
	// 	Basics struct {
	// 		Code             string
	// 		Name             string
	// 		Industry         string
	// 		Area             string
	// 		Pe               string
	// 		Outstanding      string
	// 		Totals           string
	// 		TotalAssets      string
	// 		LiquidAssets     string
	// 		FixedAssets      string
	// 		Reserved         string
	// 		ReservedPerShare string
	// 		Eps              string
	// 		Bvps             string
	// 		Pb               string
	// 		TimeToMarket     string
	// 	}
	// )

	//000333,美的集团,家用电器,
	//广东,9.05,634169.19,642493.44,
	//16750275,11683425,2100000.5,
	//1283200.25,2,1.478,8.61,3.11,19931112
	// item := stock.Basics{
	// 	Code:             "000333",
	// 	Name:             "美的集团",
	// 	Industry:         "家用电器",
	// 	Area:             "广东",
	// 	Pe:               "9.05",
	// 	Outstanding:      "634169.19",
	// 	Totals:           "642493.44",
	// 	TotalAssets:      "16750275",
	// 	LiquidAssets:     "11683425",
	// 	FixedAssets:      "2100000.5",
	// 	Reserved:         "1283200.25",
	// 	ReservedPerShare: "2",
	// 	Eps:              "1.478",
	// 	Bvps:             "8.61",
	// 	Pb:               "3.11",
	// 	TimeToMarket:     "19931112",
	// }
	// id, err := engine.Insert(item)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(id)
}
