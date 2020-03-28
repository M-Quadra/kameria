package kameria

import (
	"fmt"
	"testing"
)

func TestStructTagMap(t *testing.T) {
	fmt.Println("[kameria] StructTagMap Test...")

	type tsModel struct {
		KeyName string `tag:"key1"`
		KeyUser string `tag:"key2"`
	}

	model := tsModel{}
	needMap := map[string]interface{}{
		"key1": &model.KeyName,
		"key2": &model.KeyUser,
	}
	fmt.Println("needMap:", needMap)

	tagMap, err := StructTagMap("tag", &model)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println("tagMap:", tagMap)

	for k, v := range tagMap {
		if needMap[k] != v {
			t.Fail()
		}
	}
}
