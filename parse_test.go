package humantime

import (
	"testing"
	"time"
)

var format = "2006-01-02 15:04:05"

func MakeTime() time.Time {	
	return time.Date(2014, 8, 11, 6, 14, 23, 123456789, time.UTC)
}

func TestHumanTime(t *testing.T) {
	if Parse("test").Format(format) != "2014" {
		t.Log(Parse("test").Format(format))
		t.Errorf("HumanTime")
	}
}
