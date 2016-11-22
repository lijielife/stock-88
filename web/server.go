package web

import (
	"io"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"

	"net/http"

	"path/filepath"

	"github.com/flosch/pongo2"
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

//Template 模板
type Template struct {
}

//Render 渲染html
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	filename, err := filepath.Abs(name)
	if err != nil {
		return err
	}
	tmpl, err := pongo2.FromFile(filename)
	if err != nil {
		return err
	}
	var ctx pongo2.Context
	if data != nil {
		ctx = data.(map[string]interface{})
	}
	return tmpl.ExecuteWriter(ctx, w)
}

func main() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:123456@/finance?charset=utf8")
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Static("static", "static")
	e.GET("/", index)
	e.Start(":8080")
}

func index(c echo.Context) error {

	return c.HTML(http.StatusOK, "webapp/index.html")
}
