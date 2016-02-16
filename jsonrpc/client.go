package jsonrpc

import (
	"log"
	"github.com/gorilla/rpc/v2/json2"
	"net/http"
	"bytes"
)

type Client struct {
	Url string
}

func (p Client) Call(method string, args interface{}, result interface{}) error {

	message, err := json2.EncodeClientRequest(method, args)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	req, err := http.NewRequest("POST", p.Url, bytes.NewBuffer(message))
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error in sending request to %s. %s", p.Url, err)
		return err
	}
	defer resp.Body.Close()

	err = json2.DecodeClientResponse(resp.Body, result)
	if err != nil {
		log.Printf("Couldn't decode response. %s", err)
		return err
	}

	return err
}
