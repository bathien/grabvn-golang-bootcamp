package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/avast/retry-go"
	"github.com/myteksi/hystrix-go/hystrix"
)

func main() {
	url := "http://localhost:8080/"
	//var body []byte
	for i := 0; i <100; i ++ {
		j := i
		retry.Do(
			func() error {
				fmt.Println("Request number:", j)
				resp, _ := http.Get(url)
				body, errBody := ioutil.ReadAll(resp.Body)
				if resp.StatusCode != 200 {
					fmt.Println(string(body))
					return errors.New("server error")
				}
				defer resp.Body.Close()
				if errBody != nil {
					return errBody
				}
				fmt.Println(string(body))
				return nil
			},
			retry.OnRetry(func(n uint, err error) {
				log.Printf("#%d: %s\n", n, err)
			}),
			retry.DelayType(retry.FixedDelay),
			retry.Attempts(2),
		)
	}


	hystrix.ConfigureCommand("my_circuit_breaker", hystrix.CommandConfig{
		Timeout:                     1000,
		MaxConcurrentRequests:       100,
		ErrorPercentThreshold:       25,
		QueueSizeRejectionThreshold: 100,
	})
	//var body []byte
	for i := 0; i <100; i ++ {
		j := i
		hystrix.Do("my_circuit_breaker", func() error {
			fmt.Println("Request number:", j)
			resp, _ := http.Get(url)
			body, errBody := ioutil.ReadAll(resp.Body)
			if resp.StatusCode != 200 {
				fmt.Println(string(body))
				return errors.New("server error")
			}
			defer resp.Body.Close()
			if errBody != nil {
				return errBody
			}
			fmt.Println(string(body))
			return nil
		}, func(e error) error {
			log.Printf("%s\n", e)
			return e
		})
	}
}
