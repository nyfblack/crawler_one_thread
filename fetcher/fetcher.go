package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

//获取网页内容，并转换为utf8格式输出
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyReader := bufio.NewReader(resp.Body)
	//根据页面内容判断网页的编码格式
	e := determineEncoding(bodyReader)
	//转换为utf8的编码
	utf8Read := transform.NewReader(bodyReader, e.NewDecoder())

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is error : %d", resp.StatusCode)
	}

	return ioutil.ReadAll(utf8Read)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("determine encoding is error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
