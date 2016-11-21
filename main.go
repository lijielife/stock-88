package main

import (
	"finance/model"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

var engine *xorm.Engine

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/finance?charset=utf8")
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		var assets []model.Asset
		err := engine.Where("code=?", "000333").Find(&assets)
		if err != nil {
			panic(err)
		}
		return c.JSON(http.StatusOK, assets)
	})
	e.Static("static", "static")
	e.Logger.Fatal(e.Start(":8080"))
}
