package helper

import (
	"math"
	"time"
)

func TimeCheck(reqTime int64) bool {
	now := time.Now().UnixMilli()
	return math.Abs(float64(reqTime-now)) > 60000
}
