package main

import (
	"log"

	"github.com/EriconYu/BaXingNengLiang/handlers"
	"github.com/devfeel/dotweb"
)

func main() {
	port := 8091
	log.SetPrefix("---- ")
	log.Println(" Start ... 八星能量 @", port, "... ")
	dotapp := dotweb.New()
	handlers.InitRoute(dotapp.HttpServer)
	if e := dotapp.StartServer(port); e != nil {
		log.Println("Start server error", e.Error())
	}
}
