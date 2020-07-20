package main

import (
	"context"
	"fmt"
	proto "github.com/leehom/go-twirp"
	"log"
	"net/http"
)

func main() {
	client := proto.NewEchoProtobufClient("http://localhost:8080", &http.Client{})

	response, err := client.Say(context.Background(), &proto.Request{Text: "Hello World"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response:%s\n", response.GetText())
}