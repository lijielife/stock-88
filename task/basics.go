package task

import (
	"finance/config"
	"finance/stock"
	"fmt"
	"strings"

	"time"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//NewBasicsTask 创建基本信息抓取任务
func NewBasicsTask() *BasicsTask {
	engine, err := xorm.NewEngine("mysql", config.MysqlConnectionString)
	if err != nil {
		panic(err)
	}
	return &BasicsTask{
		interval: 24 * time.Hour,
		engine:   engine,
	}
}

//BasicsTask 基本信息抓取任务
type BasicsTask struct {
	interval time.Duration
	engine   *xorm.Engine
}

//Run 运行任务
func (t *BasicsTask) Run() {
	var err error

	//初始化结构
	err = t.CreateTables()
	if err != nil {
		log.Panic(err)
	}
	var first = 2005

	for first <= time.Now().Year() {
		for i := 1; i <= 4; i++ {
			t.Report(first, i)
		}
		first++
	}
	fmt.Println("抓取完毕")
}

//CreateTables 创建表结构
func (t *BasicsTask) CreateTables() error {
	return t.engine.Sync(new(stock.Basics), new(stock.Report))
}

//Interval 任务执行间隔时间
func (t *BasicsTask) Interval() time.Duration {
	return t.interval
}

//List 同步股票列表基本信息
func (t *BasicsTask) List() {
	list, err := stock.Basicser.List()
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range list {
		if strings.ToLower(item.Code) == "code" {
			continue
		}
		err := t.InsertOrUpdate(new(stock.Basics), item, "code=? and name=?", item.Code, item.Name)
		if err != nil {
			log.Println(err, item)
		}
	}
}

//Report 同步股票报表数据到数据库
func (t *BasicsTask) Report(year, quarter int) {
	// var last stock.Report
	// // ok, err := t.engine.Cols("year", "quarter").Desc("year", "quarter").Get(&last)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	list, err := stock.Basicser.Report(year, quarter)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("本次数据量：", len(list))
	batch := map[string]stock.Report{}
	for i, item := range list {
		exists := t.exists(new(stock.Report), "code=? and name=? and year=? and quarter=?", item.Code, item.Name, year, quarter)
		if !exists {
			batch[fmt.Sprintf("%s%s%d%d", item.Code, item.Name, item.Year, item.Quarter)] = item
		}
		if len(batch) >= 100 || i == len(list)-1 {
			var tmp = []stock.Report{}
			for _, i := range batch {
				tmp = append(tmp, i)
			}
			affected, err := t.engine.Insert(&tmp)
			if err != nil {
				log.Println(err, tmp)
			}
			batch = map[string]stock.Report{}
			fmt.Println("report:", affected)
		}
	}
}

//Profit 同步股票利润数据到数据库
func (t *BasicsTask) Profit() {

}

//Operation 同步股票运营数据到数据库
func (t *BasicsTask) Operation() {

}

//Grow 同步股票成长数据到数据库
func (t *BasicsTask) Grow() {

}

//Debtpay 同步股票偿债数据到数据库
func (t *BasicsTask) Debtpay() {

}

//Cashflow 同步股票现金流数据到数据库
func (t *BasicsTask) Cashflow() {

}

//InsertOrUpdate 插入或更新
func (t *BasicsTask) InsertOrUpdate(st interface{}, data interface{}, where string, pk ...interface{}) error {
	total, err := t.engine.Where(where, pk...).Count(st)
	if err != nil {
		return err
	}

	if total == 0 {
		_, err = t.engine.Insert(data)
	} else {
		_, err = t.engine.Id(pk).Update(data)
	}
	return err
}

func (t *BasicsTask) exists(st interface{}, where string, pk ...interface{}) bool {
	count, err := t.engine.Where(where, pk...).Count(st)
	if err != nil {
		log.Println(err)
	}
	return count > 0
}
