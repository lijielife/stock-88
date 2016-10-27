package stock

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list, err := Basicser.List()
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("获取沪深上市公司基本情况失败")
	}
	if list[0].Code == "" {
		t.Fatal("获取沪深上市公司基本情况数据失败")
	}
	fmt.Println(list[1])
}

func TestReport(t *testing.T) {
	list, err := Basicser.Report(2014, 1)
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("获取业绩报表失败")
	}
	if list[0].Code == "" {
		t.Fatal("获取业绩报表数据错误")
	}
	if len(list) <= 1000 {
		t.Fatal("业绩报表数据量不足，查询错误")
	}
	fmt.Println(list[0])
}

func TestProfit(t *testing.T) {
	list, err := Basicser.Profit(2014, 1)
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("获取盈利能力失败")
	}
	if list[0].Code == "" {
		t.Fatal("获取盈利能力数据错误")
	}
	if len(list) <= 1000 {
		t.Fatal("盈利能力数据量不足，查询错误")
	}
	fmt.Println(list[0])
}

func TestOperation(t *testing.T) {
	list, err := Basicser.Operation(2014, 1)
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("获取运营能力失败")
	}
	if list[0].Code == "" {
		t.Fatal("获取运营能力数据错误")
	}
	if len(list) <= 1000 {
		t.Fatal("运营能力数据量不足，查询错误")
	}
	fmt.Println(list[0])
}

func TestGrowth(t *testing.T) {
	list, err := Basicser.Growth(2014, 1)
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("获取成长能力失败")
	}
	if list[0].Code == "" {
		t.Fatal("获取成长能力数据错误")
	}
	if len(list) <= 1000 {
		t.Fatal("成长能力数据量不足，查询错误")
	}
	fmt.Println(list[0])
}

func TestDebtpaying(t *testing.T) {
	list, err := Basicser.Debtpaying(2014, 1)
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("获取偿债能力失败")
	}
	if list[0].Code == "" {
		t.Fatal("获取偿债能力数据错误")
	}
	if len(list) <= 1000 {
		t.Fatal("偿债能力数据量不足，查询错误")
	}
	fmt.Println(list[0])
}

func TestCashflow(t *testing.T) {
	list, err := Basicser.Debtpaying(2014, 1)
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("获取现金流量失败")
	}
	if list[0].Code == "" {
		t.Fatal("获取现金流量数据错误")
	}
	if len(list) <= 1000 {
		t.Fatal("现金流量数据量不足，查询错误")
	}
	fmt.Println(list[0])
}
