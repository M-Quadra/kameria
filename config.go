package kameria

import "time"

// TimeLocation default time.Local
var TimeLocation = time.Local

func init() {
	if TimeLocation != nil {
		return
	}

	TimeLocation = time.Local
}
