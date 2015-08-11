package main

import (
	"strconv"

	"github.com/volatile/compress"
	"github.com/volatile/core"
	"github.com/volatile/cors"
	"github.com/volatile/log"
	"github.com/volatile/response"
	"github.com/volatile/route"
)

func main() {
	if !core.Production {
		log.Use()
	}
	compress.Use()
	cors.Use(nil)

	route.Get("^/resources/(?P<id>[0-9]+)$", func(c *core.Context, params map[string]string) {
		id, _ := strconv.Atoi(params["id"])
		response.JSON(c, map[string]interface{}{
			"id":   id,
			"type": "Resource",
		})
	})

	core.Run()
}
