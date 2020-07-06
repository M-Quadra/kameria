package kameria

import "testing"

func TestHref2Url(t *testing.T) {
	href0 := "https://www.google.com/"
	if Href2Url(href0, "") != href0 {
		t.Fail()
	}

	href1 := "//www.baidu.com"
	if Href2Url(href1, href0) != "https://www.baidu.com" {
		t.Fail()
	}

	href2 := "/index.html"
	if Href2Url(href2, href0) != "https://www.google.com/index.html" {
		t.Fail()
	}

	href3 := "?wd=2"
	baseURL3 := "https://www.baidu.com/s?wd=1"
	if Href2Url(href3, baseURL3) != "https://www.baidu.com/s?wd=2" {
		t.Fail()
	}

	href4 := "/search?q=1"
	if Href2Url(href4, href0) != "https://www.google.com/search?q=1" {
		t.Fail()
	}
}
