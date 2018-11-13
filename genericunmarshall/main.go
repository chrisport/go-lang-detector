package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type CustomMsg struct {
	Msg string
	Id  int
}

func main() {
	processor := TestProcessor{}
	msg := CustomMsg{
		Msg: "TEST",
		Id:  1,
	}
	var targetType reflect.Type = processor.getType()
	data, err := json.Marshal(msg)
	panicOnErr(err)

	fmt.Printf("\n%v\n%v\n", msg, string(data))
	targetInstance := reflect.New(targetType).Interface()
	err = json.Unmarshal(data, targetInstance)
	panicOnErr(err)
	processor.process(targetInstance)
}

func panicOnErr(err error) {
	time.Sleep(500 * time.Millisecond)

	if err != nil {
		panic(err)
	}
}

type ProcessorFunction interface {
	//process(*pubsub.Message) *model.Message
	process(interface{})
	getType() reflect.Type
}

type TestProcessor struct {
}

func (t *TestProcessor) process(input interface{}) {
	inputMessage := input.(*CustomMsg)

	if inputMessage.Msg != "TEST" || inputMessage.Id != 1 {
		panic("FML")
	}
}

func (t *TestProcessor) getType() reflect.Type {
	return reflect.TypeOf(CustomMsg{})
}
