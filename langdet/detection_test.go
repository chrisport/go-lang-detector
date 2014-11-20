package langdet

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func createMapRanking(tokensInRank ...string) map[string]int {
	rankMap := make(map[string]int)
	for i, token := range tokensInRank {
		rankMap[token] = i
	}
	return rankMap
}

func TestGetDistance(t *testing.T) {
	Convey("Subject: Test getDistance", t, func() {
		Convey("same profiles should return distance 0", func() {
			rankMapA := createMapRanking("a", "b", "c")
			rankMapB := createMapRanking("a", "b", "c")
			dist := getDistance(rankMapA, rankMapB, 10)
			So(dist, ShouldBeZeroValue)
		})

		Convey("same profiles with 1 rank swapped should return distance 2", func() {
			rankMapA := createMapRanking("a", "b", "c")
			rankMapB := createMapRanking("a", "c", "b")
			dist := getDistance(rankMapA, rankMapB, 10)
			So(dist, ShouldEqual, 2)
		})

		Convey("same profiles except 1 token different should return distance 10 when maxDifference is 10", func() {
			rankMapA := createMapRanking("a", "b", "c")
			rankMapB := createMapRanking("a", "b", "d")
			dist := getDistance(rankMapA, rankMapB, 10)
			So(dist, ShouldEqual, 10)
		})

		Convey("entirely different profiles with 3 tokens should return distance 30 if maxDistance is set to 10", func() {
			rankMapA := createMapRanking("a", "b", "c")
			rankMapB := createMapRanking("e", "f", "g")
			dist := getDistance(rankMapA, rankMapB, 10)
			So(dist, ShouldEqual, 30)

		})
	})
}
