package main

import (
	"context"
	"log"

	"github.com/ag9920/kstudy/kitex_gen/api"
	"github.com/ag9920/kstudy/kitex_gen/api/kstudy"
	"github.com/cloudwego/kitex/client"
)

func main() {
	client, err := kstudy.NewClient("kstudy", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.Request{
		Message1: "message1",
		Message2: "message2",
	}
	resp, err := client.Concat(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
