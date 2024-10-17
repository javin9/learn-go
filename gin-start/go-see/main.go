package main

import (
	"github/javin9/go-see/pkg/setting"
	"github/javin9/go-see/routers"
	"log"
)

func initSetup() {
	setting.Setup()
}

func main() {
	initSetup()
	r := routers.InitRouters()
	err := r.Run(":3333")
	if err != nil {
		log.Fatalln("start service  error")
	}
}
