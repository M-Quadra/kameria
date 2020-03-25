package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/m_quadra/kameria"
)

func main() {
	fmt.Println("kameria の test")

	structTagMapTest()

	dayTimerTest()
	for { //for Test
	}
}

func structTagMapTest() {
	fmt.Println("\nStruct Tag Map Test......")

	type tsModel struct {
		KeyName string `dbkey:"key1"`
		KeyUser string `dbkey:"key2"`
	}

	var model tsModel
	needMap := map[string]interface{}{
		"key1": &model.KeyName,
		"key2": &model.KeyUser,
	}
	fmt.Println("needMap:", needMap)

	tagMap, err := kameria.StructTagMap("dbkey", &model)
	if err != nil {
		fmt.Println(".StructTagMap error")
	} else {
		fmt.Println(" tagMap:", tagMap)
	}
	for k, v := range tagMap {
		if needMap[k] != v {
			fmt.Println("???", k, "fail")
		}
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
