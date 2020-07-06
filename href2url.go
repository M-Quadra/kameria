package kameria

import "strings"

// Href2Url href to url
func Href2Url(href, baseURL string) string {
	if strings.Contains(href, "://") {
		return href
	}

	if strings.HasPrefix(href, "//") {
		ary := strings.Split(baseURL, ":")
		if len(ary) <= 0 {
			return ""
		}

		return ary[0] + ":" + href
	}

	if strings.HasPrefix(href, "/") {
		ary := strings.Split(baseURL, "/")
		if len(ary) < 3 {
			return ""
		}

		ary = ary[0:3]
		return strings.Join(ary, "/") + href
	}

	if strings.Contains(href, "?") {
		ary := strings.Split(baseURL, "?")
		if len(ary) <= 0 {
			return ""
		}

		return ary[0] + href
	}

	ary := strings.Split(baseURL, "/")
	if len(ary) < 2 {
		return ""
	}

	ary[len(ary)-1] = href
	return strings.Join(ary, "/")
}
