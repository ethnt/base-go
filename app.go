package main

import (
	"github.com/hoisie/web"
	"github.com/hoisie/mustache"
)

func index(val string) string {
	return mustache.RenderFileInLayout("views/index.mustache", "views/layouts/application.mustache", map[string]string{"c":"world"})
}

func main() {
	web.Get("/(.*)", index)
	web.Run("0.0.0.0:3000")
}
