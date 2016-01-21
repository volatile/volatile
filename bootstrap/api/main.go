package main

import (
	"flag"
	"math/rand"
	"strconv"
	"time"

	"github.com/volatile/compress"
	"github.com/volatile/core"
	"github.com/volatile/cors"
	"github.com/volatile/log"
	"github.com/volatile/response"
	"github.com/volatile/route"
	"github.com/volatile/secure"
)

func main() {
	// Initialization

	if core.Production {
		rand.Seed(time.Now().UTC().UnixNano())
	}
	flag.Parse()

	// Handlers

	if !core.Production {
		log.Use()
	}
	secure.Use(nil)
	cors.Use(nil)
	compress.Use()

	route.Get("^/resources/(?P<id>[0-9]+)$", func(c *core.Context, params map[string]string) {
		id, _ := strconv.Atoi(params["id"])
		response.JSON(c, map[string]interface{}{
			"id":   id,
			"type": "Resource",
		})
	})

	core.Run()
}
