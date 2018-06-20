package main

import (
	_ "testbeego/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/http"
)

func main() {

	var UrlManager = func(ctx *context.Context) {
		url := ctx.Input.URL()
		if url != "/" && url != "/login" {
			_, ok := ctx.Input.Session("login").(string)
			if !ok {
				ctx.Redirect(http.StatusFound, "/")
			}
		}
	}
	beego.InsertFilter("/*",beego.BeforeRouter, UrlManager)
	beego.Run()
}

