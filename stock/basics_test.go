package stock

import "testing"

func TestList(t *testing.T) {
	list, err := Basicser.List()
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("列表为空")
	}
	if list[0].Code == "" {
		t.Fatal("获取的列表错误")
	}
	t.Log(list[0])
}

func TestProfit(t *testing.T) {
	list, err := Basicser.Profit(2014, 1)
	if err != nil {
		t.Error(err)
	}
	if len(list) <= 0 {
		t.Fatal("报表为空")
	}
	if list[0].Code == "" {
		t.Fatal("获取的报表数据错误")
	}
	if len(list) <= 1000 {
		t.Fatal("数据量不足，查询错误")
	}
}
