package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Benchmark_ReflectiveUnmarshall(b *testing.B) {
	processor := TestProcessor{}
	msg := CustomMsg{
		Msg: "TEST",
		Id:  1,
	}
	var targetType reflect.Type = processor.getType()
	data, err := json.Marshal(msg)
	panicOnErr(err)

	for i := 0; i < b.N; i++ {
		targetInstance := reflect.New(targetType).Interface()
		err = json.Unmarshal(data, targetInstance)
		panicOnErr(err)
	}
}

func Benchmark_NormalUnmarshall(b *testing.B) {
	msg := CustomMsg{
		Msg: "TEST",
		Id:  1,
	}
	data, err := json.Marshal(msg)
	panicOnErr(err)

	for i := 0; i < b.N; i++ {
		var targetInstance = &CustomMsg{}
		err = json.Unmarshal(data, targetInstance)
		panicOnErr(err)
	}

}
