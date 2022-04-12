package main

import (
	api "github.com/ag9920/kstudy/kitex_gen/api/kstudy"
	"log"
)

func main() {
	svr := api.NewServer(new(KStudyImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
