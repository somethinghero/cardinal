package random

import (
	"fmt"
)

func hitJudgeTest(k float64) int {
	hit := false
	times := 0
	for hit == false {
		hit = HitJudge(k)
		times++
	}

	return times
}

func hitJudgeTestOne(k float64, testCount int) {
	ex := 0
	times := 0
	for i := 0; i < testCount; i++ {
		times = hitJudgeTest(k)
		ex += times

		//fmt.Printf("hitJudgeTest for hit rate:%v, times:%v \n", k, times)
	}

	fmt.Printf("hitJudgeTest for hit rate:%v, ex:%v \n", k, float32(ex)/float32(testCount))
}

//HitJudgeTests test func
func HitJudgeTests() {

	testCount := 50
	k := 0.9

	hitJudgeTestOne(k, testCount)

	k = 0.92
	hitJudgeTestOne(k, testCount)

	k = 0.93
	hitJudgeTestOne(k, testCount)

	k = 0.99
	hitJudgeTestOne(k, testCount)

	k = 0.999
	hitJudgeTestOne(k, testCount)

}
