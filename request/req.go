package req

import (
	"context"

	"github.com/qwertyqq2/test_task/request/data"
)

type Req interface {
	SendRequest(ctx context.Context) <-chan *data.Data

	Cancel()

	Proccessing() bool
}
