package select_

import (
	"errors"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(urls ...string) (string, error) {
	return ConfigurableRacer(tenSecondTimeout, urls...)
}

func ConfigurableRacer(timeout time.Duration, urls ...string) (string, error) {
	ch := make(chan string)

	for _, url := range urls {
		go func(u string) {
			http.Get(u)
			ch <- u
		}(url)
	}

	select {
	case result := <-ch:
		return result, nil
	case <-time.After(timeout):
		return "", errors.New("request timeout")
	}
}
