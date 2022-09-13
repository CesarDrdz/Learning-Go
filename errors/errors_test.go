package timeparse

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"19:00:12", true},
		{"1:3:33", true},
		{"bad", false},
		{"1:3:33", false},
		{"0:59:59", true},
		{"", false},
		{"34:35", false},
		{"aa:bb:cc", false},
		{"4:36:", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
