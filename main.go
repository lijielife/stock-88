package main

import (
	"finance/stock"
	"fmt"
)

func main() {
	data, err := stock.List()
	if err != nil {
		panic(err)
	}
	for _, item := range data {
		fmt.Println(item)
	}
}
