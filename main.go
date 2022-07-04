package main

import (
	"github.com/flyotlin/go-todo/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run()
}
