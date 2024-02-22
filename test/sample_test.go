package main

import "testing"

func TestMultiple(t *testing.T) {

	t.Run("Positive", func(t *testing.T) {
		t.Parallel()
		t.Log("Positive")

		var x, y, result = 2, 2, 4

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf(" %d != %d", realResult, result)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		t.Parallel()
		t.Log("Negative")
		var x, y, result = -2, 4, -8

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf(" %d != %d", realResult, result)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		t.Parallel()
		t.Log("Positive")

		var x, y, result = 2, 2, 4

		realResult := Add(x, y)

		if realResult != result {
			t.Errorf("%d != %d", realResult, result)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		t.Parallel()
		t.Log("Negative")

		var x, y, result = -2, -2, -4

		realResult := Add(x, y)

		if realResult != result {
			t.Errorf("%d != %d", realResult, result)
		}
	})
}
