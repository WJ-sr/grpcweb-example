package main

import (
	"honnef.co/go/js/dom"
	r "myitcv.io/react"

	"github.com/johanbrandhorst/grpcweb-example/client/container"
)

//go:generate gopherjs build app.go -m -o html/client.js
//go:generate bash -c "cd compiled/ && go run assets_generate.go"
//go:generate bash -c "rm html/*.js*"

func main() {
	domTarget := dom.GetWindow().Document().GetElementByID("app")

	r.Render(container.Container(), domTarget)
}
