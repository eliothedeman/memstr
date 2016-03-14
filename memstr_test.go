package memstr

import "testing"

var (
	memoryTests = map[string]int64{
		"1k": 1000,
		"2K": 2048,
	}
)

func TestParseMemory(t *testing.T) {
	for k, v := range memoryTests {
		n, err := Parse(k)
		if err != nil {
			t.Error(err)
		}
		if n != v {
			t.Fail()
		}
	}
}
