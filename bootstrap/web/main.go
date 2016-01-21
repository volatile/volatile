package main

import (
	"flag"
	"math/rand"
	"time"

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
	// Initialization

	if core.Production {
		rand.Seed(time.Now().UTC().UnixNano())
	}
	flag.Parse()

	i18n.Use(locales, language.English, true)
	response.TemplatesFuncs(i18n.TemplatesFuncs)

	// Handlers

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
