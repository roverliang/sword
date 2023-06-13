package util

import (
	"testing"
	"time"
)

func TestGetCNMothDay(t *testing.T) {
	ti := time.Now()
	t.Log(GetCNMothDay(ti))
}

func TestCheckContinuity(t *testing.T) {
	l := make(JsonTimeSlice, 0)
	l = append(l, JsonTime{
		Time: time.Now(),
	})
	l = append(l, JsonTime{
		Time: time.Now().Add(-48 * time.Hour),
	})
	l = append(l, JsonTime{
		Time: time.Now().Add(24 * time.Hour),
	})
	b := CheckContinuity(l)
	t.Log(b)
}

func TestGetBuTimeStr(t *testing.T) {
	now := time.Now()
	ts := []time.Time{
		now.Add(-20 * time.Second),
		now.Add(-59 * time.Second),
		now.Add(-60 * time.Second),
		now.Add(-61 * time.Second),
		now.Add(-(3600 - 1) * time.Second),
		now.Add(-3600 * time.Second),
		now.Add(-(3600 + 1) * time.Second),
		now.Add(-(3600*24 - 1) * time.Second),
		now.Add(-3600 * 24 * time.Second),
		now.Add(-(3600*24 + 1) * time.Second),
		now.Add(-(3600*24*7 - 1) * time.Second),
		now.Add(-3600 * 24 * 7 * time.Second),
		now.Add(-(3600*24*7 + 1) * time.Second),
		now.Add(-(3600*24*30 - 1) * time.Second),
		now.Add(-3600 * 24 * 30 * time.Second),
		now.Add(-(3600*24*30 + 1) * time.Second),
		now.Add(-(3600*24*360 - 1) * time.Second),
		now.Add(-3600 * 24 * 360 * time.Second),
		now.Add(-(3600*24*360 + 1) * time.Second),
	}
	for _, v := range ts {
		t.Log(GetBuTimeStr(v, now))
	}
}
