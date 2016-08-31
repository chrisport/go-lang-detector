package langdet_test

import (
	"strings"
	"testing"

	"github.com/chrisport/go-lang-detector/langdet"
	. "github.com/smartystreets/goconvey/convey"
)

func createMapRanking(tokensInRank ...string) map[string]int {
	rankMap := make(map[string]int)
	for i, token := range tokensInRank {
		rankMap[token] = i
	}
	return rankMap
}

func TestNew(t *testing.T) {
	Convey("Subject: New detector", t, func() {
		dd := langdet.NewDefaultLanguages()
		d := langdet.NewDetector()
		Convey("Detector should be initialized", func() {
			So(d.Languages, ShouldNotBeNil)
			So(dd.Languages, ShouldNotBeNil)
		})
	})
	Convey("Subject: New detector with default languages", t, func() {
		_ = langdet.NewDefaultLanguages()
		d := langdet.NewDetector()
		Convey("Detector should be initialized", func() {
			So(d.Languages, ShouldNotBeNil)
		})
	})
	Convey("Subject: New detector with languages from reader", t, func() {
		languageMapAsJson := "[{\"Profile\":{\"____t\":1,\"___t\":3,\"___t_\":5,\"__t\":7,\"__t_\":6,\"__t__\":9,\"_t\":15,\"_t_\":12,\"_t__\":2,\"_t___\":11,\"t\":4,\"t_\":8,\"t__\":14,\"t___\":13,\"t____\":10},\"Name\":\"english\"}]"
		reader := strings.NewReader(languageMapAsJson)
		d := langdet.NewWithLanguagesFromReader(reader)
		Convey("Detector should be initialized", func() {
			So(d.Languages, ShouldNotBeNil)
		})
	})
	Convey("Subject: Initialize DefaultLanguage with languages from reader", t, func() {
		languageMapAsJson := "[{\"Profile\":{\"____t\":1,\"___t\":3,\"___t_\":5,\"__t\":7,\"__t_\":6,\"__t__\":9,\"_t\":15,\"_t_\":12,\"_t__\":2,\"_t___\":11,\"t\":4,\"t_\":8,\"t__\":14,\"t___\":13,\"t____\":10},\"Name\":\"english\"}]"
		reader := strings.NewReader(languageMapAsJson)
		langdet.InitWithDefaultFromReader(reader)
		Convey("Detector should be initialized", func() {
			So(langdet.DefaultDetector.Languages, ShouldNotBeNil)
		})
	})
}

func TestAddLanguage(t *testing.T) {
	Convey("Subject: Add Language by text to new Detector", t, func() {
		d := langdet.Detector{}
		So(d.Languages, ShouldBeNil)

		en := "This is an english sentence"
		d.AddLanguageFromText(en, "en")

		Convey("Detector should get initialized and the language should be added", func() {
			So(d.Languages, ShouldNotBeNil)
			So(len(*d.Languages), ShouldEqual, 1)
			So((*d.Languages)[0].Name, ShouldEqual, "en")
		})
	})
	Convey("Subject: Add Language directly to new Detector", t, func() {
		d := langdet.Detector{}
		So(d.Languages, ShouldBeNil)

		d.AddLanguage(langdet.Language{Name: "en"})

		Convey("Detector should get initialized and the language should be added", func() {
			So(d.Languages, ShouldNotBeNil)
			So(len(*d.Languages), ShouldEqual, 1)
			So((*d.Languages)[0].Name, ShouldEqual, "en")
		})
	})
}

func TestClosest(t *testing.T) {
	Convey("Subject: Test GetClosestLanguage", t, func() {
		Convey("When finding a closest language", func() {
			s := "Hello I am english text, what is your language? I really dont know you say?"
			d := langdet.NewDetector()
			d.AddLanguageFromText(s, "english")
			d.AddLanguageFromText("Je parles français et toi?", "french")
			Convey("Should return string with the language name", func() {
				res := d.GetClosestLanguage(s)
				So(res, ShouldEqual, "english")
			})
		})
		Convey("When not finding a closest language", func() {
			s := "Hello I am english text, what is your language? I really dont know you say?"
			d := langdet.NewDetector()
			d.AddLanguageFromText("Je parles français et toi?", "french")
			Convey("Should return string \"undefined\"", func() {
				res := d.GetClosestLanguage(s)
				So(res, ShouldEqual, "undefined")
			})
		})
		Convey("When invalid minimum confidence", func() {
			d := langdet.NewDetector()
			d.MinimumConfidence = -19
			Convey("Should set confidence level to default", func() {
				_ = d.GetClosestLanguage("asd")
				So(d.MinimumConfidence, ShouldEqual, langdet.DefaultMinimumConfidence)
			})
		})
	})
	Convey("Subject: Test GetLanguages", t, func() {
		s := "Hello I am english text"
		d := langdet.NewDetector()
		d.AddLanguageFromText(s, "english")
		d.AddLanguageFromText("Je parles français et toi?", "french")
		Convey("Should return array with DetectionResults containing all languages", func() {
			res := d.GetLanguages(s)
			So(len(res), ShouldEqual, 2)
			So(res[0].Name, ShouldEqual, "english")
			So(res[1].Name, ShouldEqual, "french")
		})
	})

}
func TestGetDistance(t *testing.T) {
	Convey("Subject: Test getDistance", t, func() {
		Convey("same profiles should return distance 0", func() {
			rankMapA := createMapRanking("a", "b", "c")
			rankMapB := createMapRanking("a", "b", "c")
			dist := langdet.GetDistance(rankMapA, rankMapB, 10)
			So(dist, ShouldBeZeroValue)
		})

		Convey("same profiles with 1 rank swapped should return distance 2", func() {
			rankMapA := createMapRanking("a", "b", "c")
			rankMapB := createMapRanking("a", "c", "b")
			dist := langdet.GetDistance(rankMapA, rankMapB, 10)
			So(dist, ShouldEqual, 2)
		})

		Convey("same profiles except 1 token different should return distance 10 when maxDifference is 10", func() {
			rankMapA := createMapRanking("a", "b", "c")
			rankMapB := createMapRanking("a", "b", "d")
			dist := langdet.GetDistance(rankMapA, rankMapB, 10)
			So(dist, ShouldEqual, 10)
		})

		Convey("entirely different profiles with 3 tokens should return distance 30 if maxDistance is set to 10", func() {
			rankMapA := createMapRanking("a", "b", "c")
			rankMapB := createMapRanking("e", "f", "g")
			dist := langdet.GetDistance(rankMapA, rankMapB, 10)
			So(dist, ShouldEqual, 30)

		})
	})
}
