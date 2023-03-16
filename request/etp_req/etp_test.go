package etpreq

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	baseUrl = "https://etp-ets.ru/44/catalog/procedure"
)

func TestUrlUpdate(t *testing.T) {
	url := "https://etp-ets.ru/44/catalog/procedure"
	idx := 0
	urlPage1 := urlUpdate(url, idx)
	idx++
	urlPage2 := urlUpdate(url, idx)
	idx++
	urlPage3 := urlUpdate(url, idx)
	assert.Equal(t, "https://etp-ets.ru/44/catalog/procedure", urlPage1)
	assert.Equal(t, "https://etp-ets.ru/44/catalog/procedure?page=2", urlPage2)
	assert.Equal(t, "https://etp-ets.ru/44/catalog/procedure?page=3", urlPage3)

}

func TestCancel(t *testing.T) {
	t.Run("close outging", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		req := New(baseUrl)
		ch := req.SendRequest(ctx)

		cancel()

		time.Sleep(100 * time.Microsecond)
		select {
		case _, ok := <-ch:
			if ok {
				t.Fatal("chan is not closed")
			}
		}
	})

	t.Run("close inside", func(t *testing.T) {

		req := New(baseUrl)
		ch := req.SendRequest(context.Background())

		req.Cancel()

		time.Sleep(100 * time.Microsecond)
		select {
		case _, ok := <-ch:
			if ok {
				t.Fatal("chan is not closed")
			}
		}
	})

}

func TestProccessing(t *testing.T) {
	t.Run("proccessig after run", func(t *testing.T) {
		req := New(baseUrl)
		req.SendRequest(context.Background())

		time.Sleep(100 * time.Microsecond)

		if !req.Proccessing() {
			t.Fatal("not proccessing")
		}

	})

	t.Run("proccessing before run", func(t *testing.T) {
		req := New(baseUrl)
		req.SendRequest(context.Background())

		time.Sleep(100 * time.Microsecond)

		req.Cancel()

		req.SendRequest(context.Background())

		if !req.Proccessing() {
			t.Fatal("not proccessing")
		}
	})
}
