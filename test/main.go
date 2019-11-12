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
	mathCmpTest()
	unique4Test()
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
	fmt.Println("\nError Printer Test......")

	timeStr := "1970-01-01 08:00:00"
	_, err := time.Parse("2006-01-02 15:04:051", timeStr)
	kameria.HasError(err)
}

func mathCmpTest() {
	fmt.Println("\nMath Cmp Test......")

	fmt.Println("kameria.Max4Int(1, 3)\t\t", kameria.Max4Int(1, 3))
	fmt.Println("kameria.Min4Int(1, 5)\t\t", kameria.Min4Int(1, 5))
	fmt.Println("kameria.Limit4Int(1, 12, 6)\t", kameria.Limit4Int(1, 12, 6))
	fmt.Println("kameria.Limit4Int(1, -12, 6)\t", kameria.Limit4Int(1, -12, 6))
}

func unique4Test() {
	fmt.Println("\nUnique 4 Test......")

	tsStrAry := []string{"app", "google", "m$", "m$", "m$", "m$", "noDieBook"}
	tsIntAry := []int{9700, 9800, 9900, 7920, 7920, 7920, 7920}
	tsInt64Ary := []int64{3700, 3800, 3900, 3950, 3950, 3950, 3950}

	fmt.Println("kameria.Unique4String\t", kameria.Unique4String(tsStrAry))
	fmt.Println("kameria.Unique4Int\t", kameria.Unique4Int(tsIntAry))
	fmt.Println("kameria.Unique4Int64\t", kameria.Unique4int64(tsInt64Ary))
}
