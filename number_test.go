package kameria

import "testing"

func TestArabic2NumStr(t *testing.T) {
	tsStr := "۰۱۲۳۴۵۶۷۸۹"
	if Arabic2NumStr(tsStr) != "0123456789" {
		t.Fail()
	}
}

func TestTibetan2NumStr(t *testing.T) {
	tsStr := "༠༡༢༣༤༥༦༧༨༩"
	if Tibetan2NumStr(tsStr) != "0123456789" {
		t.Fail()
	}
}
