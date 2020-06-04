package kameria

import "strings"

var arabic2numMap = map[string]string{
	"۰": "0",
	"۱": "1",
	"۲": "2",
	"۳": "3",
	"۴": "4",
	"۵": "5",
	"۶": "6",
	"۷": "7",
	"۸": "8",
	"۹": "9",
}

var tibetan2numMap = map[string]string{
	"༠": "0",
	"༡": "1",
	"༢": "2",
	"༣": "3",
	"༤": "4",
	"༥": "5",
	"༦": "6",
	"༧": "7",
	"༨": "8",
	"༩": "9",
}

// Arabic2NumStr Arabic-Indic Digit to [0,9]string
func Arabic2NumStr(arabic string) string {
	for k, v := range arabic2numMap {
		arabic = strings.ReplaceAll(arabic, k, v)
	}
	return arabic
}

// Tibetan2NumStr Tibetan Digit to [0,9]string
func Tibetan2NumStr(arabic string) string {
	for k, v := range tibetan2numMap {
		arabic = strings.ReplaceAll(arabic, k, v)
	}
	return arabic
}
