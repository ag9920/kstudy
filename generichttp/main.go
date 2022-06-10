package main

import (
	"log"

	api "github.com/ag9920/kstudy/generichttp/kitex_gen/http/bizservice"
)

func main() {
	svr := api.NewServer(new(BizServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
