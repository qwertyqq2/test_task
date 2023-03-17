package req

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/qwertyqq2/test_task/request/data"
)

type impl struct {
	baseUrl string

	findFn FindFunc

	process *proccess

	mups sync.RWMutex
}

func New(baseUrl string, find FindFunc) Req {
	return &impl{
		baseUrl: baseUrl,
		findFn:  find,
	}
}

func (req *impl) SendRequest(ctx context.Context) <-chan *data.Data {
	ch := make(chan *data.Data)

	if req.Proccessing() {
		close(ch)
		return ch
	}

	req.mups.Lock()
	proc := newProc(ctx, req.baseUrl, ch, req.findFn)
	req.process = proc
	req.mups.Unlock()

	proc.run()

	return ch
}

func (req *impl) Proccessing() bool {
	req.mups.RLock()
	defer req.mups.RUnlock()

	if req.process == nil {
		return false
	}

	return req.process.Proccessing()
}

func (req *impl) Cancel() {
	req.mups.Lock()
	defer req.mups.Unlock()

	req.process.close()
}

type proccess struct {
	ctx    context.Context
	cancel func()

	findFn FindFunc

	url    string
	number uint
	ch     chan *data.Data
	idx    int

	proccessing bool

	muidx sync.RWMutex
}

func newProc(parent context.Context, url string, ch chan *data.Data, findFn FindFunc) *proccess {
	ctx, cancel := context.WithCancel(parent)
	return &proccess{
		ctx:         ctx,
		cancel:      cancel,
		url:         url,
		ch:          ch,
		idx:         0,
		proccessing: true,
		findFn:      findFn,
	}
}

func (p *proccess) close() {
	p.cancel()
}

func (p *proccess) run() {
	select {
	case <-p.ctx.Done():
		close(p.ch)
		return
	default:
	}
	go func() {
		defer func() {
			close(p.ch)
			p.proccessing = false
		}()
		for {
			select {
			case <-p.ctx.Done():
				return

			default:
				url := urlUpdate(p.url, p.idx)
				resp, err := http.Get(url)
				if err != nil {
					log.Println("err Get")
					return
				}

				p.findFn(p.ctx, resp, p.ch)

				p.muidx.Lock()
				p.idx++
				p.muidx.Unlock()

			}
		}
	}()
}

func (p *proccess) Proccessing() bool {
	return p.proccessing
}

func urlUpdate(url string, idx int) string {
	if idx == 0 {
		return url
	}
	return url + "?page=" + fmt.Sprint(idx+1)
}
