package main

import (
	"context"
	"fmt"
	"time"

	etpreq "github.com/qwertyqq2/test_task/request/etp_req"
)

const (
	baseUrl = "https://etp-ets.ru/44/catalog/procedure"
)

func main() {

	req := etpreq.New(baseUrl)

	ch := req.SendRequest(context.Background())

	for {
		select {
		case d := <-ch:
			fmt.Println(d.String())
			fmt.Println("\n")
			time.Sleep(1 * time.Second)
		}
	}
}
