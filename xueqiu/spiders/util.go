package spiders

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var stockList []string

func stocks() []string {
	if stockList != nil {
		return stockList
	}
	bts, err := ioutil.ReadFile("stocks.txt")
	if err != nil {
		panic(err)
	}
	list := strings.Split(string(bts), "\n")
	fmt.Println(list)
	stockList = list
	return list
}
