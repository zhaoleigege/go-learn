package tdd

import "testing"

func TestAdd(t *testing.T) {
	assertFunc := func(t *testing.T, got, want int) {
		t.Helper()

		if got != want {
			t.Errorf("%d + %d want: %d, get: %d", 3, 4, want, got)
		}
	}

	t.Run("3 + 4的例子", func(t *testing.T) {
		sum := Add(3, 4)
		assertFunc(t, sum, 7)
	})

	t.Run("4 + 4的例子", func(t *testing.T) {
		sum := Add(4, 4)
		assertFunc(t, sum, 8)
	})
}

func BenchmarkAdd(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		sum = Add(i, sum)
	}
}
