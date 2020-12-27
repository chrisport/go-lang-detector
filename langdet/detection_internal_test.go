package langdet

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"unicode"
)

type mockComparator struct {
}

func (f *mockComparator) CompareTo(lazyLookupMap func() map[string]int, originalText string, maxRank int) DetectionResult {
	return DetectionResult{"fake", 99}
}

func (f *mockComparator) GetName() string { return "fake" }

func TestLanguageComparator(t *testing.T) {
	Convey("given no language needs lookup map ", t, func() {
		d := NewDetector()
		mocklLM := func(text string, nDepth int) func() map[string]int {
			return func() map[string]int { panic("shouldn't be needed") }
		}
		lazyLookupMap, mocklLM = mocklLM, lazyLookupMap
		d.AddLanguageComparators(&mockComparator{})
		res := d.GetClosestLanguage("some dummy text")
		Convey("Never call lazyLookupMap and return result", func() {
			So(d.Languages, ShouldNotBeNil)
			So(res, ShouldEqual, "fake")
		})
		lazyLookupMap, mocklLM = mocklLM, lazyLookupMap
	})
}

func TestChinese(t *testing.T) {
	Convey("given chinese is added", t, func() {
		d := NewDetector()
		clc := UnicodeRangeLanguageComparator{"chinese", unicode.Han}
		d.AddLanguageComparators(&clc)
		res := d.GetLanguages("not chinese")
		Convey("then do not detect non-chinese as Chinese", func() {
			So(len(res), ShouldEqual, 1)
			So(res[0].Name, ShouldEqual, "chinese")
			So(res[0].Confidence, ShouldEqual, 0)
		})

		res = d.GetLanguages("漢字")
		Convey("then do detect Chinese language", func() {
			So(len(res), ShouldEqual, 1)
			So(res[0].Name, ShouldEqual, "chinese")
			So(res[0].Confidence, ShouldEqual, 100)
		})
	})
}
