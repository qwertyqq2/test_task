package req

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/qwertyqq2/test_task/request/data"
)

const (
	rtsUrl = "https://www.rts-tender.ru/poisk/poisk-commercial-tenderi" //base url for rts
	etpUrl = "https://etp-ets.ru/44/catalog/procedure"                  // base url for etp
)

// Here are implementations of the FindFunc type for the rustender site and others

var etpFindFunc = func(ctx context.Context, resp *http.Response, ch chan *data.Data) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("err new doc")
		return
	}
	divs := doc.Find("tr")
	divs.Each(func(i int, s *goquery.Selection) {
		tsome := s.Find("td.row-procedure_name ")
		some := tsome.Text()
		var (
			name, number string
		)
		if strings.Contains(some, "(") {
			split := strings.Split(some, "(")
			name = split[0]
			number = split[1]
		}

		if strings.Contains(number, ")") {
			number = strings.Replace(number, ")", "", -1)
		}

		if !checkNumber(number) {
			return
		}
		d := data.NewData(number, name)
		if !d.IsNil() {
			select {
			case <-ctx.Done():
				return
			case ch <- d:

			}
		}
	})
}

var rtsFindFunc = func(ctx context.Context, resp *http.Response, ch chan *data.Data) {
	selectNum := func(str string) string {
		if !strings.Contains(str, "№") {
			return ""
		}
		var (
			number string
		)
		{
			split := strings.Split(str, "№")
			some := split[1]
			var last string
			for _, s := range some {
				if (s < 'a' || s > 'я') && (s < 'А' || s > 'Я') {
					continue
				}
				last = string(s)
				break
			}
			split = strings.Split(some, last)
			number = split[0]
		}
		return number
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("err new doc")
		return
	}
	divs := doc.Find("div.card-item")
	divs.Each(func(i int, s *goquery.Selection) {
		var (
			number, name string
		)
		tnum := s.Find("a")
		numstr := tnum.Text()
		numstr = strings.ReplaceAll(numstr, "РТС-ТЕНДЕРПОДРОБНЕЕ", "")
		number = selectNum(numstr)
		number = cleanSpace(number)

		tsome := s.Find("div.card-item__title")
		some := tsome.Text()
		name = cleanSpace(some)

		d := data.NewData(number, name)
		if !d.IsNil() {
			select {
			case <-ctx.Done():
				return
			case ch <- d:

			}
		}

	})
}

func NewRts() Req {
	return New(rtsUrl, rtsFindFunc)
}

func NewEtp() Req {
	return New(etpUrl, etpFindFunc)
}

func cleanSpace(s string) string {
	s = strings.ReplaceAll(s, "  ", "")
	s = strings.ReplaceAll(s, "\n", "")
	return s
}

func checkNumber(str string) bool {
	for _, s := range str {
		if (s < 'a' || s > 'z') && (s < 'A' || s > 'Z') && (s < 'a' || s > 'я') && (s < 'А' || s > 'Я') {
			continue
		}
		return false
	}
	return true
}

func cleanFromEnd(str string) string {
	var (
		n   = len(str) - 1
		res = str
	)
	for i := n; i >= 0; i-- {
		if checkNumber(string(str[i])) {
			res = str[:i]
			continue
		}
		return res
	}

	return res
}
