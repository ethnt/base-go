/*
	base-go
		file: app.go
*/

/*
	= Appendix
		I. Imports
		II. Helpers and Main
		III. Controllers
*/

/*
	I. Imports
*/

package main

import (
	"github.com/hoisie/web"
	"github.com/hoisie/mustache"
)


/*
	II. Helpers and Main
*/

func render(file string, params map[string]string) string {
	return mustache.RenderFileInLayout(file, "views/layouts/application.mustache", params)
}

func main() {
	web.Get("/(.*)", index)
	web.Run("0.0.0.0:3000")
}


/*
	III. Controllers
*/

func index(val string) string {
	return render("views/index.mustache", map[string]string{"c":"world"})
}


