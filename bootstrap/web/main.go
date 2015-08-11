package main

import (
	"github.com/volatile/compress"
	"github.com/volatile/core"
	"github.com/volatile/i18n"
	"github.com/volatile/log"
	"github.com/volatile/response"
	"github.com/volatile/route"
	"github.com/volatile/static"
)

func main() {
	if !core.Production {
		log.Use()
	}
	compress.Use()
	i18n.Use(locales, "en", true)
	static.Use(static.DefaultMaxAge)

	route.Get("^/$", func(c *core.Context) {
		response.View(c, "hello", nil)
	})

	core.Run()
}
