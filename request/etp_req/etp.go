package etpreq

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	req "github.com/qwertyqq2/test_task/request"
	"github.com/qwertyqq2/test_task/request/data"
)

type impl struct {
	baseUrl string

	process *process

	mups sync.RWMutex
}

func New(baseUrl string) req.Req {
	return &impl{
		baseUrl: baseUrl,
	}
}

func (etp *impl) SendRequest(ctx context.Context) <-chan *data.Data {
	ch := make(chan *data.Data)

	if etp.Proccessing() {
		close(ch)
		return ch
	}

	etp.mups.Lock()
	proc := newProc(ctx, etp.baseUrl, ch)
	etp.process = proc
	etp.mups.Unlock()

	proc.run()

	return ch
}

func (etp *impl) Proccessing() bool {
	etp.mups.RLock()
	defer etp.mups.RUnlock()

	if etp.process == nil {
		return false
	}

	return etp.process.proccessing
}

func (etp *impl) Cancel() {
	etp.mups.Lock()
	defer etp.mups.Unlock()

	etp.process.close()
}

type process struct {
	ctx    context.Context
	cancel func()

	url    string
	number uint
	ch     chan *data.Data
	idx    int

	proccessing bool

	muidx sync.Mutex
}

func newProc(parent context.Context, url string, ch chan *data.Data) *process {
	ctx, cancel := context.WithCancel(parent)
	return &process{
		ctx:         ctx,
		cancel:      cancel,
		url:         url,
		ch:          ch,
		idx:         0,
		proccessing: true,
	}
}

func (p *process) close() {
	p.cancel()
}

func (p *process) run() {
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

				p.findSelectors(resp)

				p.muidx.Lock()
				p.idx++
				p.muidx.Unlock()

			}
		}
	}()
}

func (p *process) findSelectors(resp *http.Response) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("err new doc")
		return
	}
	divs := doc.Find("tr")
	divs.Each(func(i int, s *goquery.Selection) {
		tsome := s.Find("td.row-procedure_name ")
		some := tsome.Text()
		tprice := s.Find("td.row-contract_start_price")
		unprice := tprice.Text()
		if strings.Contains(unprice, "RUB") {
			unprice = strings.ReplaceAll(unprice, "RUB", "")
		}
		unprice = strings.Replace(unprice, " ", "", -1)

		d := data.NewData(unprice, some)
		if !d.IsNil() {
			select {
			case <-p.ctx.Done():
				return
			case p.ch <- d:

			}
		}
	})
}

func urlUpdate(url string, idx int) string {
	if idx == 0 {
		return url
	}
	return url + "?page=" + fmt.Sprint(idx+1)
}
