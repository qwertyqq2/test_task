package req

import (
	"context"
	"net/http"

	"github.com/qwertyqq2/test_task/request/data"
)

// FindFunc is a function that parses a web page after receiving
// a response from the site resp. After processing, it sends data to the ch channel
type FindFunc func(ctx context.Context, resp *http.Response, ch chan *data.Data)

// Req is an interface that starts and cancels the process of searching for data from a web page.
type Req interface {

	//SendRequest sends a request to a web page and returns
	// a channel through which we will receive data from it.
	SendRequest(ctx context.Context) <-chan *data.Data

	//Cancel completes the process of receiving data, while closing the channel
	Cancel()

	//are we in the process of getting the data?
	Proccessing() bool
}
