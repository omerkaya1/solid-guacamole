package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/net/html"
)

func crawl(n *html.Node, result chan<- struct{}) {
	for _, a := range n.Attr {
		if n.Type != html.ElementNode {
			return
		}
		if a.Val == "aero_bcal_day_number" {
			result <- struct{}{}
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		crawl(c, result)
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	var (
		ticker  = time.NewTicker(time.Second)
		c       = http.Client{Timeout: time.Second * 5}
		result  = make(chan struct{}, 10)
		ackChan = make(chan os.Signal, 1)
	)
	defer ticker.Stop()

	signal.Notify(ackChan, syscall.SIGUSR1)

	go func() {
		for range result {
		NOTIFY:
			for {
				select {
				case <-ctx.Done():
					return
				case <-ackChan:
					break NOTIFY
				default:
					fmt.Print("\a")
					time.Sleep(time.Second)
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ticker.Reset(time.Minute * 3)
			body := strings.NewReader(fmt.Sprintf("bid=56&rnd=%d", rand.Int()%100+1))
			req, err := http.NewRequest(http.MethodPost,
				"https://appointment.mfa.gr/inner.php/reservations/aero/calendar", body)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
			resp, err := c.Do(req)
			if err != nil {
				fmt.Println(err)
				continue
			}
			switch resp.StatusCode {
			case http.StatusOK:
				data, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
					continue
				}
				_ = resp.Body.Close()

				doc, err := html.Parse(bytes.NewReader(data))
				if err != nil {
					log.Println(err)
					return
				}
				go crawl(doc, result)
			default:
				fmt.Println("status code: ", resp.StatusCode)
			}
		}
	}
}
