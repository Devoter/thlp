# thlp

thlp is a lightweight test helper. It has no dependencies and easy interface.

The package provides the following functions: `Equal`, `DeepEqual`, `Bytes`, `Ok`, `Cmp` and `Err`.

## Usage

### Check coverage

To check coverage you can use built-in Go instruments:

```sh
go test ./... -coverprofile=coverage.out 
go tool cover -html=coverage.out
```

### Equal

`Equal` calls `t.Fatalf` if expected and got arguments are not equal.

Example:

```go
func Something(arg int) int

func TestSomething(t *testing.T) {
    expected := 42
    result := Something(24)
    thlp.Equal(t, expected, result, "Expected: [%d], but got [%d]")
}
```

### DeelEqual

`DeepEqual` calls `t.Fatalf` if expected and got arguments are not equal deeply.

Example:

```go

func Something(arg int) []int

func TestSomething(t *testing.T) {
    expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
    result := Something(24)
    thlp.DeepEqual(t, expected, result, "Expected: [%v], but got [%v]")
}
```

### Bytes

Sometimes `reflect.DeepEqual` which is used under the hood of `DeepEqual` function, does not compares bytes slices correctly. Use `Bytes` to compare it.
`Bytes` calls `t.Fatalf` if expected and got bytes slices are not equal.

Example:

```go
func Something(arg int) []byte

func TestSomething(t *testing.T) {
    expected := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c}
    result := Something(24)
    thlp.Bytes(t, expected, result, "Expected: [%v], but got [%v]")
}
```

## Ok

`Ok` calls `t.Fatalf` if argument is not `true`.

Example:

```go
func Something(arg int) bool

func TestSomething(t *testing.T) {
    result := Something(24)
    thlp.Ok(t, result, "Result is not Ok")
}
```

### CMP

`Cmp` calls `t.Fatalf` if compare function returns `false`.

Example:

```go
func Something(arg int) interface{}
func Cmp(a, b interface{}) bool

func TestSomething(t *testing.T) {
    expected := struct{
        a int
        b float32
    }{
        a: 15,
        b: 22.2,
    }

    result := Something(24)
    thlp.Cmp(t, Cmp, expected, result, "Expected:\n%+v\ngot:\n%+v")
}
```

### Err

`Err` calls `t.Fatalf` if the pattern (regexp) is matched into error. If pattern is `""` and error is `nil` the assertion is `true`.

Example:

```go
func Something(arg int) error

func TestSomething(t *testing.T) {
    expected := "Error int [0-9]+ sector"

    result := Something(24)
    thlp.Err(t, expected, result, "Expected: [%v], but got [%v]")
}
```

## License

[MIT](LICENSE)
