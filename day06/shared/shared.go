package shared

import (
	"math"
)

// Magic does some magic algorithm which is explained in maths.md
func Magic(t int, s int) int {
	ft := float64(t)
	fs := float64(s)
	trecord1 := (-ft + math.Sqrt(ft*ft-4*fs)) / (-2)
	trecord2 := (-ft - math.Sqrt(ft*ft-4*fs)) / (-2)
	n := math.Ceil(trecord2-1) - math.Floor(trecord1+1) + 1
	return int(n)
}
