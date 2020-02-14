package duration

import (
	"github.com/goccha/log"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	duration := "P1Y2M3DT4H25M36S"
	now, err := time.Parse(time.RFC3339, "2020-02-13T10:50:00+09:00")
	if err != nil {
		t.Errorf("%v", err)
	}
	log.Debug("now=%v", now)
	expected, err := time.Parse(time.RFC3339, "2021-04-16T15:15:36+09:00")
	if err != nil {
		t.Errorf("%v", err)
	}
	tm := Add(now, duration)
	if expected.Equal(tm) {
		log.Debug("add=%v", tm)
	} else {
		t.Errorf("expected=%s, actual=%s", expected, tm)
	}
}

func TestParse(t *testing.T) {
	duration := "T4H"
	d := Parse(duration)
	if d == (4 * time.Hour) {
		log.Debug("OK")
	} else {
		t.Fail()
	}
}
