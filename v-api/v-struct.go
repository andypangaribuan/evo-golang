package v_api

import "github.com/labstack/echo/v4"


/* ============================================
	Created by andy pangaribuan on 2020/04/01
	Copyright BoltIdea. All rights reserved.
   ============================================ */
type Context struct {
	echo *echoContext
	Log smLog
}


type SMEchoApi struct {
	port int
	e    *echo.Echo
}


type echoContext struct {
	c echo.Context
}


type smLog struct {
	context *Context
}
