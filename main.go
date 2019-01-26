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
	fmt.Println("北海道" + sub("https://transit.yahoo.co.jp/traininfo/detail/637/0/"))
	fmt.Println("　東北" + sub("https://transit.yahoo.co.jp/traininfo/detail/1/0/"))
	fmt.Println("　秋田" + sub("https://transit.yahoo.co.jp/traininfo/detail/6/0/"))
	fmt.Println("　山形" + sub("https://transit.yahoo.co.jp/traininfo/detail/5/0/"))
	fmt.Println("　上越" + sub("https://transit.yahoo.co.jp/traininfo/detail/3/0/"))
	fmt.Println("　北陸" + sub("https://transit.yahoo.co.jp/traininfo/detail/624/0/"))
	fmt.Println("東海道" + sub("https://transit.yahoo.co.jp/traininfo/detail/7/0/"))
	fmt.Println("　山陽" + sub("https://transit.yahoo.co.jp/traininfo/detail/8/0/"))
	fmt.Println("　九州" + sub("https://transit.yahoo.co.jp/traininfo/detail/410/0/"))
}
