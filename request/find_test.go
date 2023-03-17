package req

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/qwertyqq2/test_task/request/data"
	"github.com/stretchr/testify/assert"
)

var (
	ch chan *data.Data
)

func TestCheckNumber(t *testing.T) {
	str := "2342345A32132"
	if checkNumber(str) {
		t.Fatal("is not true")
	}

	str = "234234532132"
	assert.Equal(t, true, checkNumber(str))
}

func TestCleanFromEnd(t *testing.T) {
	str := "qeeqwe    , "
	res := cleanFromEnd(str)
	assert.Equal(t, res, "qeeqwe")
}

func TestEtpFind(t *testing.T) {
	ch = make(chan *data.Data)
	ctx := context.Background()

	t.Run("test first page", func(t *testing.T) {
		resp, err := http.Get(etpUrl)
		if err != nil {
			t.Fatal(err)
		}

		go etpFindFunc(ctx, resp, ch)
		select {
		case d := <-ch:
			if d.IsNil() {
				t.Fatal("nil pack")
			}
			fmt.Println(d.String())
		}
	})
}

func TestRts(t *testing.T) {

	ch = make(chan *data.Data)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	t.Run("test first page", func(t *testing.T) {
		resp, err := http.Get(rtsUrl)
		if err != nil {
			t.Fatal(err)
		}

		go rtsFindFunc(ctx, resp, ch)
		select {
		case d := <-ch:
			if d.IsNil() {
				t.Fatal("nil pack")
			}
			fmt.Println(d.String())
		}
	})
}
