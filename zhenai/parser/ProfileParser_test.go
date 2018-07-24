package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestProfileParser(t *testing.T) {
	context, err := ioutil.ReadFile("profileParser_test_context.txt")
	if err != nil {
		panic(err)
	}

	result := ProfileParser(context, "小顺儿")
	fmt.Println(result.Items[0])
}
