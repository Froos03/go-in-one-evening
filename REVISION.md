# Go in One Evening — 10-Minute Revision

A recap of every exercise solved so far: what it taught, the key code, and the gotchas worth remembering.

**Progress:** Modules 1–5 complete, plus `06-errors` 01–02 and the three mini-projects (HTTP server, `ls`, HTTP bad request). `06-errors/03-handling-errors` is in progress.

---

## 1. Start — the shape of a Go program

| Exercise | Lesson |
|---|---|
| `02-hello` | Every runnable program is `package main` + `func main()`. |
| `03-main` | `main()` is the entry point; it takes no args and returns nothing. |
| `04-imports` | Use standard-library packages with `import "fmt"`. |
| `05-comments` | `//` comments out a line, which is handy for disabling code. |

```go
package main

import "fmt"

func main() {
	fmt.Print("I ")       // no newline
	fmt.Print("<3 ")
	// fmt.Print("Hello ") // commented out → never runs
	fmt.Println("Go!")     // adds newline; Println separates args with spaces
}
```

**Remember:** `Print` adds no newline. `Println` adds spaces between args and a newline at the end.

---

## 2. Variables

| Exercise | Lesson |
|---|---|
| `01-variables` | Short declaration `x := value` infers the type. |
| `02-types` | `var x int` declares (zero value), `x = 31` assigns. Basic types: `int`, `float64`, `bool`, `string`. |
| `03-const` | `const hoursInDay = 24` can never be reassigned. |
| `04-operators` | `* + - /` follow normal precedence: `price * qty + shipping`. |
| `05-globals` | Package-level variables must use `var`/`const`, **not** `:=`. |

```go
var price = 40.0     // global: var keyword required
const taxRate = 0.1  // global constant

func main() {
	var days int     // declared → zero value 0
	days = 31        // assigned
	pi := 3.1415     // inferred float64
	tax := taxRate * price
}
```

**Gotchas**
- `:=` only works **inside functions**.
- An unused local variable is a **compile error** in Go.
- Zero values: `0`, `""`, `false`, `nil`.

---

## 3. Functions

| Exercise | Lesson |
|---|---|
| `01-call` | Define with `func Name() {}` and call as many times as you like. |
| `02-arguments` | Parameters have a type **after** the name: `func Greet(name string)`. |
| `03-return` | The return type goes after the parameters. Same-type params can share a type: `a, b, c int`. |

```go
func Greet(name string) {
	fmt.Println("Hello, " + name)  // + concatenates strings
}

func Sum(a, b, c, d, e int) int {
	return a + b + c + d + e
}
```

**Remember:** a **Capitalized** name (`Sum`) is *exported*, meaning visible outside the package. A lowercase name (`hello`) is private to the package.

### Project: HTTP server

```go
func main() {
	http.HandleFunc("/hello", hello)             // route → handler
	log.Fatal(http.ListenAndServe(":8080", nil)) // start server, exit on error
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", name)            // write response to w
}
```

---

## 4. Collections

| Exercise | Lesson |
|---|---|
| `01-arrays` | Arrays have a **fixed size that is part of the type**: `[4]string{...}`. |
| `02-slices` | Slices have no size in the brackets (`[]string{...}`) and can grow. |
| `03-len` | `len(x)` returns the number of elements. |

```go
var colors  = [5]string{"red", "green", "blue"} // array
var systems = []string{"linux", "macos", "windows"} // slice

len(colors)  // 5 ← not 3! the remaining 2 slots are "" (zero values)
len(systems) // 3
```

**Gotcha:** an array's `len` is its declared size, not how many values you wrote.

### Project: `ls`

```go
func listFiles(dirname string) []string {
	var dirs []string                   // nil slice, ready to append
	files, _ := os.ReadDir(dirname)     // _ discards the error
	for _, f := range files {           // range gives index, value
		dirs = append(dirs, f.Name())   // append returns the new slice!
	}
	return dirs
}
// strings.Join(names, " ") → "a b c"
```

**Remember:** you must write `x = append(x, v)`. A bare `append(x, v)` does nothing useful.

---

## 5. Conditionals

| Exercise | Lesson |
|---|---|
| `01-if` | No parentheses around conditions; braces are required. |
| `02-else` | `if / else if / else` chains. |
| `03-switch` | `switch` on a value; there is **no fallthrough** by default, so no `break` is needed. `default` catches everything else. |
| `04-booleans` | `&&` (and), `\|\|` (or), `!` (not). You can return a boolean expression directly. |
| `05-returning-early` | Guard clauses: handle the bad case first, then `return`. |

```go
if x == 0 {
	return "zero"
} else if x < 0 {
	return "negative"
} else {
	return "positive"
}

switch d {
case "N":
	return "north"
default:
	return "unknown"
}

func In20thCentury(year int) bool {
	return year > 1900 && year <= 2000   // no if needed
}

func ResetPassword(code int) {
	if code != 2022 {
		return            // early return keeps the happy path unindented
	}
	Password = "new-password"
}
```

### Project: HTTP bad request

```go
name := r.URL.Query().Get("name")   // read ?name=... query param
if name == "" {
	w.WriteHeader(http.StatusBadRequest) // 400
	return                                // early return!
}
fmt.Fprintf(w, "Hello, %s", name)
```

---

## 6. Errors *(in progress)*

| Exercise | Lesson |
|---|---|
| `01-multiple-returns` | Functions can return several values: `(string, string)`. |
| `02-returning-errors` | Idiomatic Go returns `(result, error)`, with `nil` meaning "no error". |
| `03-handling-errors` | 🚧 Next up: check `err != nil` from `os.Mkdir`. |

```go
func Swap(x, y string) (string, string) {
	return y, x               // (a temp variable works too)
}
x, y = Swap(x, y)

func Divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("division by zero") // zero value + error
	}
	return x / y, nil                             // result + no error
}

result, err := Divide(100, 50)
if err != nil {
	// handle it
}
```

**Remember:** an error is a value. Return it last, check `if err != nil` right away, and return a zero value for the other results when something fails.

---

## Cheat Sheet

| Concept | Syntax |
|---|---|
| Short declare (in functions) | `x := 5` |
| Declare with type | `var x int` |
| Constant | `const x = 5` |
| Function | `func Name(a, b int) int { ... }` |
| Multiple returns | `func f() (int, error)` |
| Array / slice | `[3]int{...}` / `[]int{...}` |
| Length | `len(s)` |
| Append | `s = append(s, v)` |
| Loop over collection | `for i, v := range s { }` |
| Ignore a value | `_` |
| Error creation | `errors.New("msg")` |
| Exported name | starts with a **C**apital letter |
| Print | `Print`, `Println`, `Fprintf(w, "%s", v)` |
