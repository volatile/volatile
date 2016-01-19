package main

import (
	"github.com/volatile/compress"
	"github.com/volatile/core"
	"github.com/volatile/i18n"
	"github.com/volatile/log"
	"github.com/volatile/response"
	"github.com/volatile/route"
	"github.com/volatile/secure"
	"github.com/volatile/static"
	"golang.org/x/text/language"
)

func main() {
	i18n.Use(locales, language.English, true)

	if !core.Production {
		log.Use()
	}
	secure.Use(nil)
	compress.Use()
	static.Use(static.DefaultMaxAge)

	route.Get("^/$", func(c *core.Context) {
		response.Template(c, "hello", nil)
	})

	core.Run()
}
