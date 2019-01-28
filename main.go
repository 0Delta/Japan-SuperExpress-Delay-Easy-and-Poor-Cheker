/*
 * Japan SuperExpress Delay Easy (and Poor) Cheker
 * @author 0Delta
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"time"
)

func sub(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "ERR : Request"
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "ERR : Do"
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "ERR : Read"
	}
	r := regexp.MustCompile(`<meta property="og:description" content=.*/>`)
	ret := r.FindAllStringSubmatch(string(body), 1)
	return ret[0][0]
}

func main() {
	var wg sync.WaitGroup
	args := [][]string{
		{"北海道", "https://transit.yahoo.co.jp/traininfo/detail/637/0/"},
		{"　東北", "https://transit.yahoo.co.jp/traininfo/detail/1/0/"},
		{"　秋田", "https://transit.yahoo.co.jp/traininfo/detail/6/0/"},
		{"　山形", "https://transit.yahoo.co.jp/traininfo/detail/5/0/"},
		{"　上越", "https://transit.yahoo.co.jp/traininfo/detail/3/0/"},
		{"　北陸", "https://transit.yahoo.co.jp/traininfo/detail/624/0/"},
		{"東海道", "https://transit.yahoo.co.jp/traininfo/detail/7/0/"},
		{"　山陽", "https://transit.yahoo.co.jp/traininfo/detail/8/0/"},
		{"　九州", "https://transit.yahoo.co.jp/traininfo/detail/410/0/"}}

	for _, a := range args {
		wg.Add(1)
		go func(place string, url string) {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("ERR : Request")
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Println("ERR : Do")
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println("ERR : Read")
			}
			r := regexp.MustCompile(`<meta property="og:description" content=.*/>`)
			ret := r.FindAllStringSubmatch(string(body), 1)
			fmt.Println(place + ret[0][0])
			// return ret[0][0]
			wg.Done()
		}(a[0], a[1])
		time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}
