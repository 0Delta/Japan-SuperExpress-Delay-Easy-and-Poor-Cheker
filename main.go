/*
 * Japan SuperExpress Delay Easy (and Poor) Cheker
 * @author 0Delta
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"
)

var REG = regexp.MustCompile(`<meta property="og:description" content=.*/>`)

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

			defer wg.Done()

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERR : Request")
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERR : Do")
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERR : Read")
			}
			ret := REG.FindAllStringSubmatch(string(body), 1)
			fmt.Println(place + ret[0][0])
			// return ret[0][0]
		}(a[0], a[1])
		time.Sleep(10 * time.Millisecond)
	}
	wg.Wait()
}
