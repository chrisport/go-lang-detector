package playground // Bitflow
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var sampleProp0 = SampleType{id: "0"}
var sampleProp1 = SampleType{id: "1"}

type SampleType struct {
	id string
}

func TestGivenStack_whenPeek_thenReturnLastInserted(t *testing.T) {
	s := make(StackMap, 0)
	s.Push("prop", sampleProp0)
	s.Push("prop", sampleProp1)

	prop := s.Peek("prop")
	prop = s.Peek("prop")
	prop = s.Peek("prop")

	assert.Equal(t, prop, sampleProp1)
}

func TestGivenInexistingStack_whenPeekOrDefault_thenReturnDefault(t *testing.T) {
	s := make(StackMap, 0)
	s.Push("prop", sampleProp0)

	res := s.PeekOrDefault("non-existing_stack", sampleProp1)

	assert.Equal(t, res, sampleProp1)
}

func TestGivenSampleStack_whenPeekOrDefault_thenReturnDefault(t *testing.T) {
	s := make(StackMap, 0)
	s.Push("prop", sampleProp0)

	res := s.PeekOrDefault("prop", sampleProp1)


	assert.Equal(t, res , sampleProp0)
}

func TestGivenStackOncePopped_whenPop_thenReturnLastInserted(t *testing.T) {
	s := make(StackMap, 0)
	s.Push("prop", sampleProp0)
	s.Push("prop", sampleProp1)

	prop := s.Pop("prop")

	assert.Equal(t, prop, sampleProp1)
}

func TestGivenStackPopped_whenPopThenCurrent_thenReturnSecondLastInserted(t *testing.T) {
	s := make(StackMap, 0)
	s.Push("prop", sampleProp0)
	s.Push("prop", sampleProp1)
	prop := s.Pop("prop")

	prop = s.Peek("prop")

	assert.Equal(t, prop, sampleProp0)
}

func TestGivenEmptyStack_whenCurrentOrPop_thenReturnNil(t *testing.T) {
	s := make(StackMap, 0)

	popped := s.Pop("prop")
	current := s.Peek("prop")

	assert.Nil(t, popped)
	assert.Nil(t, current)
}

func TestSampleTypeStack_whenPopAll_thenFillArrayWithAllElements(t *testing.T) {
	s := make(StackMap, 0)

	s.Push("prop", sampleProp0)
	s.Push("prop", sampleProp1)

	var allProps []SampleType
	s.PopAll("prop", &allProps)

	assert.Equal(t, allProps[0], sampleProp0)
	assert.Equal(t, allProps[1], sampleProp1)
}

func TestSampleTypeStackWith5Elements_whenLen_thenReturn5(t *testing.T) {
	s := make(StackMap, 0)
	s.Push("sampleStack", sampleProp0)
	s.Push("sampleStack", sampleProp0)
	s.Push("sampleStack", sampleProp0)
	s.Push("sampleStack", sampleProp0)
	s.Push("sampleStack", sampleProp0)

	res := s.Len("sampleStack")

	assert.Equal(t, res, 5)
	assert.Equal(t, res, len(s["sampleStack"]))
}

func TestSampleTypeStack_whenPopAllWithWrongArrayType_thenPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	s := make(StackMap, 0)
	s.Push("prop", SampleType{id: "0"})
	s.Push("prop", SampleType{id: "1"})

	var allProps []int
	s.PopAll("prop", &allProps)

	// must panic
}

func TestSampleTypeStack_whenPopAllWithWrongType_thenPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	s := make(StackMap, 0)
	s.Push("prop", SampleType{id: "0"})
	s.Push("prop", SampleType{id: "1"})

	var allProps []SampleType
	s.PopAll("prop", allProps)

	// must panic
}
