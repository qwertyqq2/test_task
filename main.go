package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	req "github.com/qwertyqq2/test_task/request"
)

type conf struct {
	rtsParsing bool // parse rus-tender
	etpParsing bool // parse fabricant-tender
}

func ParseFlags() (conf, error) {
	rtsParsing := flag.Bool("rts", false, "parse rus-tender")
	etpParsing := flag.Bool("etp", false, "parse fabricant-tender")
	flag.Parse()
	return conf{
		rtsParsing: *rtsParsing,
		etpParsing: *etpParsing,
	}, nil
}

func main() {

	help := flag.Bool("h", false, "Display Help")
	config, err := ParseFlags()
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

	if *help {
		fmt.Println("This program demonstrates a simple parsing of tender sites")
		fmt.Println()
		flag.PrintDefaults()
		return
	}

	var (
		ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	)

	defer cancel()

	if config.etpParsing && config.rtsParsing {
		log.Fatal("you want two at once?")
	}

	if config.etpParsing {
		etp := req.NewEtp()
		ch := etp.SendRequest(ctx)
		for {
			select {
			case <-ctx.Done():
				return

			case d := <-ch:
				fmt.Println(d.String())
				time.Sleep(100 * time.Millisecond)
			}
		}
	}

	if config.rtsParsing {
		rts := req.NewRts()
		ch := rts.SendRequest(ctx)
		for {
			select {
			case <-ctx.Done():
				return

			case d := <-ch:
				fmt.Println(d.String())
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
