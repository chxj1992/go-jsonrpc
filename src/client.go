package src

import (
	"log"
	"github.com/gorilla/rpc/v2/json2"
	"net/http"
	"bytes"
)

type Client struct {
	Url string
}

func (p Client) Call(method string, args interface{}, result interface{}) {

	message, err := json2.EncodeClientRequest(method, args)
	if err != nil {
		log.Fatalf("%s", err)
	}
	req, err := http.NewRequest("POST", p.Url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", p.Url, err)
	}
	defer resp.Body.Close()

	err = json2.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	log.Printf("%s", result)
}

type HelloArgs struct {
	Who string
}
type HelloReply struct {
	Message string
}