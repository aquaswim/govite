package govite

import (
	"fmt"
	"testing"
)

func TestUtils_appendIfNotExists(t *testing.T) {
	tc := []struct {
		target   []string
		src      string
		expected []string
	}{
		{
			target:   []string{},
			src:      "a",
			expected: []string{"a"},
		},
		{
			target:   []string{"a"},
			src:      "a",
			expected: []string{"a"},
		},
		{
			target:   []string{"a"},
			src:      "b",
			expected: []string{"a", "b"},
		},
	}

	testCompareArray := func(a []string, b []string) error {
		if len(a) != len(b) {
			return fmt.Errorf("array length not match")
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return fmt.Errorf("array expected to be equal, got unequal: %+v != %+v", a, b)
			}
		}
		return nil
	}

	for i, s := range tc {
		if err := testCompareArray(appendIfNotExists(s.target, s.src), s.expected); err != nil {
			t.Fatalf("test case #%d return err: %s", i, err)
		}
	}
}
