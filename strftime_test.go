package strftime

import (
	"testing"
	"time"
)

var (
	timeStr   = "Mon Oct 08 17:48:51 2018"
	layoutStr = "Mon Jan 02 15:04:05 2006"
	s         string
)

func TestLayout(t *testing.T) {
	layout := Layout("%c")
	if layout != layoutStr {
		t.Fatal("Expected equal layouts ", layout, " and ", layoutStr)
	}
}

func TestParse(t *testing.T) {
	gt, err := time.Parse(layoutStr, timeStr)
	if err != nil {
		t.Fatal("error parsing with time.Parse", err)
	}

	st, err := Parse("%c", timeStr)
	if err != nil {
		t.Fatal("error parsing with strftime", err)
	}

	if !st.Equal(gt) {
		t.Fatal("Expected equal times ", st, " and ", gt)
	}
}

func TestFormat(t *testing.T) {
	st, err := Parse("%c", timeStr)
	if err != nil {
		t.Fatal("error parsing with strftime", err)
	}

	s := Format("%c", st)

	if s != timeStr {
		t.Fatal("Expected equal formats ", s, " and ", timeStr)
	}
}

func BenchmarkFormat(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		s = Layout(`%c %x %X %Z time format`)
	}
}
