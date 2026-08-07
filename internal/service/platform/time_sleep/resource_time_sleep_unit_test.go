package time_sleep

import (
	"testing"
)

func TestValidDurationRegex(t *testing.T) {
	valid := []string{
		"1s", "30s", "5m", "2h", "500ms",
		"1.5s", "0.5m", "1.5h",
		"100ms", "999ms",
	}
	for _, v := range valid {
		if !validDuration.MatchString(v) {
			t.Errorf("expected %q to be valid but regex rejected it", v)
		}
	}
}

func TestInvalidDurationRegex(t *testing.T) {
	invalid := []string{
		"",
		"5",
		"5x",
		"5minutes",
		"5 s",
		"s5",
		"-5s",
		"1us",
		"1ns",
		"abc",
		"1.2.3s",
	}
	for _, v := range invalid {
		if validDuration.MatchString(v) {
			t.Errorf("expected %q to be invalid but regex accepted it", v)
		}
	}
}
