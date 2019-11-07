package main

import (
	"fmt"
	"time"

	"github.com/m_quadra/kameria"
)

func main() {
	fmt.Println("kameria の test")

	ipv4Str := "192.168.0.1"
	ipv4Num, err := kameria.IPv4StringToInt(ipv4Str)
	fmt.Println(ipv4Str, "->", ipv4Num, err)
	ipv4Str, err = kameria.IPv4IntToString(ipv4Num)
	fmt.Println(ipv4Num, "->", ipv4Str, err)

	structTagMapTest()
	errorPrinterTest()
}

func structTagMapTest() {
	fmt.Println("\nStructTagMap(tag string, model interface{}) (map[string]interface{}, error)")

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
	if kameria.HasError(err) {
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

func errorPrinterTest() {
	fmt.Println("\nTest:", "HasError(err error) bool")

	timeStr := "1970-01-01 08:00:00"
	_, err := time.Parse("2006-01-02 15:04:051", timeStr)
	kameria.HasError(err)
}
