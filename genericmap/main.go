package main

import (
	server0 "github.com/ag9920/kstudy/genericmap/kitex_gen/kitex/test/server/exampleservice"
	"log"
)

func main() {
	svr := server0.NewServer(new(ExampleServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
