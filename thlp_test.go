package thlp_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/Devoter/thlp"
)

type assertMock struct {
	result string
}

func (am *assertMock) Fatalf(format string, arguments ...interface{}) {
	am.result = fmt.Sprintf(format, arguments...)
}

func (am *assertMock) Helper() {}

func TestEqual(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)

		Equal(am, 127, 127, "expected %d, got %d")
		if am.result != "" {
			t.Fatalf("Result was not expected, but got: \"%s\"\n", am.result)
		}
	})

	t.Run("NotEqual", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)
		expected := "expected text is \"expected text\", but got \"got text\"\n"

		Equal(am, "expected text", "got text", "expected text is \"%s\", but got \"%s\"")
		if expected != am.result {
			t.Fatalf("Expected result is \"%s\", but got \"%s\"\n", expected, am.result)
		}
	})
}

func TestDeepEqual(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)

		got := struct {
			count int
			data  []byte
		}{
			count: 127,
			data:  []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'},
		}

		expected := struct {
			count int
			data  []byte
		}{
			count: 127,
			data:  []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'},
		}

		DeepEqual(am, expected, got, "expected:\n%+v\ngot:\n%+v")
		if am.result != "" {
			t.Fatalf("Result was not expected, but got: \"%s\"\n", am.result)
		}
	})

	t.Run("NotEqual", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)

		got := struct {
			count int
			data  []int8
		}{
			count: 127,
			data:  []int8{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'},
		}

		expected := struct {
			count int64
			data  []int8
		}{
			count: 127,
			data:  []int8{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'},
		}

		res := fmt.Sprintf("expected:\n%+v\ngot:\n%+v\n", expected, got)

		DeepEqual(am, expected, got, "expected:\n%+v\ngot:\n%+v")
		if res != am.result {
			t.Fatalf("Expected result is \"%s\", but got \"%s\"\n", res, am.result)
		}
	})
}

func TestBytes(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)

		got := []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'}
		expected := []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'}

		Bytes(am, expected, got, "expected:\n%+v\ngot:\n%+v")
		if am.result != "" {
			t.Fatalf("Result was not expected, but got: \"%s\"\n", am.result)
		}
	})

	t.Run("NotEqual", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)

		got := []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd'}
		expected := []byte{'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd', '!'}

		res := fmt.Sprintf("expected:\n%+v\ngot:\n%+v\n", expected, got)

		Bytes(am, expected, got, "expected:\n%+v\ngot:\n%+v")
		if res != am.result {
			t.Fatalf("Expected result is \"%s\", but got \"%s\"\n", res, am.result)
		}
	})
}

func TestOk(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)

		Ok(am, true, "assertion failed")
		if am.result != "" {
			t.Fatalf("Result was not expected, but got: \"%s\"\n", am.result)
		}
	})

	t.Run("NotOk", func(t *testing.T) {
		t.Parallel()

		am := new(assertMock)
		res := "assertion failed\n"

		Ok(am, false, "assertion failed")
		if res != am.result {
			t.Fatalf("Expected result is \"%s\", but got \"%s\"\n", res, am.result)
		}
	})
}

func TestCmp(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)
		cmp := func(a, b interface{}) bool { return true }

		Cmp(am, cmp, 100, 100, "expected: %d, but got: %d")
		if am.result != "" {
			t.Fatalf("Result was not expected, but got: \"%s\"\n", am.result)
		}
	})

	t.Run("NotEqual", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)
		cmp := func(a, b interface{}) bool { return false }

		res := "expected: 100, but got: 100\n"

		Cmp(am, cmp, 100, 100, "expected: %d, but got: %d")
		if res != am.result {
			t.Fatalf("Expected result is \"%s\", but got \"%s\"\n", res, am.result)
		}
	})
}

func TestErr(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)

		Err(am, "test error", errors.New("Extended: test error"), "expected: \"%v\", but got: \"%v\"")
		if am.result != "" {
			t.Fatalf("Result was not expected, but got: \"%s\"\n", am.result)
		}
	})

	t.Run("NotEqual", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)
		res := "expected: \"test error\", but got: \"Extended: test err\"\n"

		Err(am, "test error", errors.New("Extended: test err"), "expected: \"%v\", but got: \"%v\"")
		if res != am.result {
			t.Fatalf("Expected result is \"%s\", but got \"%s\"\n", res, am.result)
		}
	})

	t.Run("NilError", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)

		Err(am, "", nil, "expected: \"%v\", but got: \"%v\"")
		if am.result != "" {
			t.Fatalf("Result was not expected, but got: \"%s\"\n", am.result)
		}
	})

	t.Run("NilErrorWithPattern", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)
		res := "expected: \"test\", but got: \"<nil>\"\n"

		Err(am, "test", nil, "expected: \"%v\", but got: \"%v\"")
		if res != am.result {
			t.Fatalf("Expected result is \"%s\", but got \"%s\"\n", res, am.result)
		}
	})

	t.Run("NotNilErrorWithPattern", func(t *testing.T) {
		t.Parallel()
		am := new(assertMock)
		res := "expected: \"\", but got: \"new error\"\n"

		Err(am, "", errors.New("new error"), "expected: \"%v\", but got: \"%v\"")
		if res != am.result {
			t.Fatalf("Expected result is \"%s\", but got \"%s\"\n", res, am.result)
		}
	})
}
