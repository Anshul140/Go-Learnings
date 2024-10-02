package iteration

import "testing"

func TestRepeater(t *testing.T) {
	got := charRepeater("a")
	expected := "aaaaa"

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func charRepeater(charater string) string {
	var repeated string

	for i := 0; i < 5; i++ {
		repeated = repeated + charater
	}

	return repeated
}

func BenchmarkCharRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		charRepeater("a")
	}
}
