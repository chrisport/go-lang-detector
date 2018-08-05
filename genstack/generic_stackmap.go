package playground // Bitflow

import (
	"reflect"
	"fmt"
)

type StackMap map[string][]interface{}

func (s StackMap) Push(key string, st interface{}) {
	_, ok := s[key]
	if !ok {
		s[key] = make([]interface{}, 0)
	}
	s[key] = append(s[key], st)
}

func (s StackMap) PopOrDefault(key string, defaultValue interface{}) interface{} {
	a := s[key]
	if len(a) == 0 {
		return defaultValue
	}

	p := a[len(a)-1]
	s[key] = a[:len(a)-1]
	return p
}

func (s StackMap) Pop(key string) interface{} {
	return s.PopOrDefault(key, nil)
}

func (s StackMap) Peek(key string) interface{} {
	a := s[key]
	if len(a) == 0 {
		return nil
	}
	return a[len(s[key])-1]
}

func (s StackMap) PopAll(key string, t interface{}) {
	valT := reflect.ValueOf(t)
	if (valT.Kind() != reflect.Ptr) || (reflect.Indirect(valT).Kind() != reflect.Slice|reflect.Array) {
		panic("PopAll requires input type to be pointer to slice: *[]<type>")
	}
	elemT := valT.Type().Elem().Elem()

	slice := reflect.MakeSlice(reflect.SliceOf(elemT), len(s[key]), len(s[key]))
	for a := 0; a < len(s[key]); a++ {
		curr := slice.Index(a)
		valA := reflect.ValueOf(s[key][a])
		if valA.Kind() != curr.Kind() {
			panic(fmt.Sprintf("Invalid type passed to PopAll, expected %v but have %v", valA.Type(), curr.Type()))
		}

		curr.Set(valA)
	}

	reflect.Indirect(valT).Set(slice)
}
