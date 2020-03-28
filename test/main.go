package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/m_quadra/kameria"
)

func main() {
	fmt.Println("kameria の test")

	dayTimerTest()
	for { //for Test
	}
}

func dayTimerTest() {
	fmt.Println("\ndayTimer Test......")

	nowTime := time.Now()
	nowTime = nowTime.Add(time.Second * 1)

	dayTimer := kameria.DayTimer{
		Hour:   nowTime.Hour(),
		Minute: nowTime.Minute(),
		Second: nowTime.Second(),
		Handler: func() {
			fmt.Println("DayTimer running ~_~")
			rand.Seed(time.Now().UnixNano())
			fmt.Println(rand.Intn(10))
		},
		IsRepeat: true,
	}
	dayTimer.Start()

	// ctm := time.NewTimer(time.Second * 10)
	// go func() {
	// 	<-ctm.C
	// 	dayTimer.Cancel()
	// }()
}
