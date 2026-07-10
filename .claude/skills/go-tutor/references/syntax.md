# Go Syntax Module

A contrast-driven syntax reference for a learner who already knows another
language. Teach these **on demand** as builds require them — don't dump it all
at once. For each item, prefer the flawed-code-first pattern: show a broken
snippet, ask for the fix, then explain.

## 1. Declarations & zero values
```go
var x int            // 0        — zero value, never nil/undefined
var s string         // ""       — not nil
name := "London"     // short decl, only inside functions; type inferred
const Max = 100      // constants, no := , can be untyped
```
- No `let`/`const`-hoisting quirks. `:=` declares + assigns; `=` reassigns.
- **Unused local variables and imports are compile ERRORS**, not warnings.

## 2. Functions & multiple returns
```go
func divide(a, b int) (int, error) {   // multiple return values
    if b == 0 {
        return 0, errors.New("divide by zero")
    }
    return a / b, nil
}
q, err := divide(10, 2)
```
- The `(value, error)` pair is THE Go idiom. Always handle `err` before using
  the value. No exceptions-as-control-flow.
- Named returns `(q int, err error)` exist but use sparingly.

## 3. Error handling (no try/catch)
```go
resp, err := http.Get(url)
if err != nil {
    return fmt.Errorf("fetching %s: %w", url, err)  // %w wraps
}
defer resp.Body.Close()   // runs when function returns
```
- `panic`/`recover` exist but are for truly exceptional cases, not normal flow.
- `defer` = scheduled cleanup, runs LIFO at function exit.

## 4. Structs & tags (no classes)
```go
type Weather struct {
    City string  `json:"name"`      // struct tag maps JSON key -> field
    Temp float64 `json:"main.temp"` // (nested needs care — teach later)
}
w := Weather{City: "London", Temp: 12.5}
w2 := &Weather{}     // pointer to struct
```
- Capitalized field/func names are **exported** (public); lowercase = package-private.
- No inheritance — use composition (embedding).

## 5. Methods: value vs pointer receivers
```go
func (w Weather) Summary() string { ... }   // value receiver: copy, read-only
func (w *Weather) SetTemp(t float64) { w.Temp = t }  // pointer: can mutate
```
- Rule of thumb: pointer receiver if it mutates or the struct is large; be
  consistent within a type.

## 6. Slices, arrays, maps
```go
arr := [3]int{1, 2, 3}          // fixed-size array (rare)
nums := []int{1, 2, 3}          // slice (the workhorse)
nums = append(nums, 4)          // append returns a new slice header
m := map[string]int{"a": 1}
v, ok := m["a"]                 // comma-ok idiom; ok=false if missing
```
- Slices are references to a backing array; arrays are value types (copied).
- Zero value of a slice/map is `nil`; you can range a nil slice, but must
  `make` a map before writing to it.

## 7. Control flow
```go
if v, err := f(); err != nil { ... }   // statement + condition
for i := 0; i < n; i++ { }             // classic
for i, x := range nums { }             // range
for cond { }                           // "while"
for { }                                // infinite
switch x { case 1: ...; default: ... } // no fallthrough by default
```

## 8. Interfaces (implicit)
```go
type Fetcher interface {
    Fetch(city string) (Weather, error)
}
// Any type with a matching Fetch method satisfies Fetcher — no "implements".
```
- "Accept interfaces, return structs." Keep interfaces small (1–3 methods).
- This is the key to testability: swap a real HTTP fetcher for a fake in tests.

## 9. Packages & tooling
```
go mod init example/weather   // create module
go run .                      // compile + run
go build                      // produce binary
go test ./...                 // run tests
gofmt -w .   /  go vet ./...  // format / lint
```
- One `package main` with `func main()` = an executable.
- Formatting is not a preference: `gofmt` is canonical.

## Common gotchas to quiz on
- Unused import/variable → won't compile.
- `:=` vs `=` (redeclaration in same scope).
- Loop variable capture in goroutines/closures (fixed in Go 1.22+, still worth
  understanding).
- Forgetting to check `err`; forgetting `defer Close()`.
- Comparing a value receiver method mutating a copy (change is lost).
