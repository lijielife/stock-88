package web

import (
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func main() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:123456@/finance?charset=utf8")
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Static("static", "static")
	e.Static("/", "webapp")

}
