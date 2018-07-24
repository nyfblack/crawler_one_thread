package parser

import (
	"io/ioutil"
	"testing"
)

//测试CityListParser
func TestCityListParser(t *testing.T) {
	context, err := ioutil.ReadFile("cityListParser_test_context.txt")
	if err != nil {
		t.Error("get context err :", err)
		return
	}

	result := CityListParser(context)

	//测试Parser的结果
	const num = 470
	expectedCitys := []string{
		"City: 阿坝", "City: 阿克苏", "City: 阿拉善盟",
	}
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(result.Items) != num {
		t.Errorf("City numbers is error ; expected is %d, but parser is %d",
			num, len(result.Items))
	}
	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("City is error : expected #%d is %s, but was %s",
				i, city, result.Items[i].(string))
		}
	}

	if len(result.Requests) != num {
		t.Errorf("Request numbers is error ; expected is %d, but parser is %d",
			num, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if url != result.Requests[i].Url {
			t.Errorf("Url is error : expected #%d is %s, but was %s",
				i, url, result.Requests[i].Url)
		}
	}
}
