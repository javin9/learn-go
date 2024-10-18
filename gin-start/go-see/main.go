package main

import (
	"github/javin9/go-see/model"
	"github/javin9/go-see/pkg/setting"
	"github/javin9/go-see/routers"
	validator_translation "github/javin9/go-see/validator"
	"log"
)

func initSetup() {
	setting.Setup()
	model.Setup()
	validator_translation.Setup("zh")
}

func main() {
	initSetup()
	r := routers.InitRouters()
	err := r.Run(":3333")
	if err != nil {
		log.Fatalln("start service  error")
	}
}
