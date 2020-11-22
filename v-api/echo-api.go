package v_api

import (
	"github.com/andypangaribuan/evo-golang/clog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */

//noinspection GoUnusedExportedFunction
func BuildEcho(port int, logMiddleware clog.EchoMiddleware) *SMEchoApi {
	e := echo.New()

	if logMiddleware != nil {
		e.Use(logMiddleware.Logger)
	}

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "")
	})

	sm := SMEchoApi{port, e}
	return &sm
}


func (slf *SMEchoApi) Serve() {
	slf.e.Logger.Fatal(slf.e.Start(":" + strconv.Itoa(slf.port)))
}


func (slf *SMEchoApi) POST(path string, handler HandlerFunc) {
	slf.e.POST(path, func(c echo.Context) error {
		ctx := Context{
			echo: &echoContext{c: c},
		}
		ctx.Log = smLog{context: &ctx}

		return handler(ctx)
	})
}
