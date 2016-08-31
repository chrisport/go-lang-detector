package langdet_test

import (
	"testing"

	"github.com/chrisport/go-lang-detector/langdet"
	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkCalculateElapsedTimeInMillis(b *testing.B) {
	sampleText := "TEXT"

	for n := 0; n < b.N; n++ {
		_ = langdet.CreateOccurenceMap(sampleText, 5)
	}

}

func TestCreateProfile(t *testing.T) {
	sampleText := "TEXT"
	Convey("Subject: Test create profile\n", t, func() {
		Convey("result of 'TEXT' should contain when n=3:  T:2, E:1, X:1, T:1, ...", func() {
			result := langdet.CreateOccurenceMap(sampleText, 3)
			So(result["T"], ShouldEqual, 2)
			So(result["E"], ShouldEqual, 1)
			So(result["X"], ShouldEqual, 1)
			So(result["_T"], ShouldEqual, 1)
			So(result["TE"], ShouldEqual, 1)
			So(result["EX"], ShouldEqual, 1)
			So(result["XT"], ShouldEqual, 1)
			So(result["T_"], ShouldEqual, 1)
			So(result["_T"], ShouldEqual, 1)
			So(result["__T"], ShouldEqual, 1)
			So(result["_TE"], ShouldEqual, 1)
			So(result["TEX"], ShouldEqual, 1)
			So(result["EXT"], ShouldEqual, 1)
			So(result["XT_"], ShouldEqual, 1)
			So(result["T__"], ShouldEqual, 1)
		})
	})
}

func TestCreateProfileWithObscure(t *testing.T) {
	sampleText := "...TE X123123T"
	Convey("Subject: Test create profile\n", t, func() {
		Convey("result of 'TEXT' should contain when n=3:  T:2, E:1, X:1, T:1, ...", func() {
			result := langdet.CreateOccurenceMap(sampleText, 3)
			So(result["T"], ShouldEqual, 2)
			So(result["E"], ShouldEqual, 1)
			So(result["X"], ShouldEqual, 1)
			So(result["_T"], ShouldEqual, 1)
			So(result["TE"], ShouldEqual, 1)
			So(result["T_"], ShouldEqual, 1)
			So(result["_T"], ShouldEqual, 1)
			So(result["XT"], ShouldEqual, 1)
			So(result["__T"], ShouldEqual, 1)
			So(result["__X"], ShouldEqual, 1)
			So(result["_XT"], ShouldEqual, 1)
			So(result["XT_"], ShouldEqual, 1)
			So(result["_TE"], ShouldEqual, 1)
			So(result["TE_"], ShouldEqual, 1)
			So(result["XT_"], ShouldEqual, 1)
			So(result["T__"], ShouldEqual, 1)
		})
	})

}

func TestRanking(t *testing.T) {
	sampleText := "AABBCC"
	Convey("Subject: Test create Ranking Lookup Map\n", t, func() {
		Convey("AABBCC should result in A, B and C on rank 1-3", func() {
			result := langdet.CreateOccurenceMap(sampleText, 5)
			ranking := langdet.CreateRankLookupMap(result)
			So(ranking["A"], ShouldBeBetween, 0, 4)
			So(ranking["B"], ShouldBeBetween, 0, 4)
			So(ranking["C"], ShouldBeBetween, 0, 4)

		})
	})

}
