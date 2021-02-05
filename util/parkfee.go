package util

import (
	"ProManageSystem/model/parkmodel"
	"time"
)

func Calcutefee(t1, t2 time.Time) int {
	timediff := t2.Sub(t1).Minutes()
	if timediff < 15 {
		return 0
	}
	h := (int(timediff) / 60) + 1
	return h * parkmodel.HourFee
}
