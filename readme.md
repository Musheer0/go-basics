# 📘 GoLang Super Notes

This is my personal **GoLang cheat sheet** — a mix of syntax, examples, and deep explanations so I never have to Google the same thing twice.

---

## 🔁 Loops

Go has **only one loop**: the `for` loop. But it has multiple flavors.

### 1. While-style `for` loop

```go
var i int = 0
for i <= 3 {
    i++
}
```

* Works like a **while loop** in other languages.
* Condition checked at every iteration.

---

### 2. Range loop (like `for..in` in JS/Python)

```go
for i := range 69 {
    fmt.Printf("%d\n", i)
}
```

* `range N` gives numbers `0` → `N-1`.
* Super useful for iteration without declaring `i` manually.

---

### 3. Classic C-style `for` loop

```go
for i := 0; i <= 3; i++ {
    fmt.Printf("%d\n", i)
}
```

* Has **init**, **condition**, **increment** all in one line.

---

## ⚖️ If / Else

```go
for i := range 69 {
    if i == 17 {
        // do something
    } else if i > 64 {
        // do something
    } else {
        // do something
    }
}
```

* Pretty standard.
* No parentheses around conditions (Go hates them).

---

## 🔀 Switch

### Type switch (used with `interface{}`)

```go
func typeCheck(i interface{}) {
    switch i.(type) {
    case string:
        // do something
    case int:
        // do something
    default:
        // fallback
    }
}
```

* `i.(type)` checks the **actual type** of an interface value.
* Super handy for polymorphism-like situations.

---

## 📦 Arrays & Slices

### Arrays

```go
names := [2]string{"w", "s"} // fixed-size array
len(names)                   // → 2
```

### Slices (dynamic arrays)

```go
nums := make([]int, 2) // [0, 0], length=2, capacity=2
append(nums, 1)        // like JS `push`
cap(nums)              // gives capacity
```

* `var nums []int` → declares empty **nil slice**.

---

### Slice with custom capacity

```go
nums := make([]int, 2, 6) // len=2, cap=6 → [0,0]
```

### Copy slices

```go
var source = make([]int, 0, 5)
source = append(source, 1)

var target = make([]int, len(source))
copy(target, source) // copies values
```

### Slice operator

```go
nums := []int{1, 2, 4}

fmt.Print(nums[0:2]) // → [1,2]
fmt.Print(nums[:2])  // → [1,2]
fmt.Print(nums[2:])  // → [4]
```

* Works exactly like Python slicing.

---

### Compare slices

```go
import "slices"

slices.Equal(num1, num2) // returns true/false
```

---

## 🗺️ Maps

```go
m := make(map[string]string)

// Add
m["golang"] = "code"

// Get
val := m["golang"] // if not found → zero value ("", 0, false...)

// Delete
delete(m, "golang")

// Reset
clear(m)
```

### Check existence

```go
k, ok := m["value"]
// ok = true if key exists
```

### Compare maps

```go
import "maps"

m1 := map[string]int{"price": 565, "stock": 2}
m2 := map[string]int{"price": 545, "stock": 44}
fmt.Print(maps.Equal(m1, m2))
```

### Iterate over map / slice / string

```go
for i, num := range nums {
    fmt.Print(num)
}

for i, c := range "golang" {
    fmt.Println(i, c)
}
```

---

## 🛠️ Functions

### Basic function with multiple returns

```go
func add(a, b int, name string) (string, int) {
    fmt.Println(a + b)
    return name, a + b
}

name, sum := add(1, 2, "mushher")
fmt.Println(name, sum)
```

---

### Function as parameter

```go
func process(fn func(a int) int) int {
    return fn(1)
}
```

---

### Variadic function (like JS `...args`)

```go
func sum(nums ...int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}

sum(1, 2, 3, 4, 5)
arr := []int{1, 2, 3}
sum(arr...) // spread operator
```

---
## 🎯 Pointers in Go

Pointers = variables that **store memory addresses** instead of values.
Think: instead of holding the pizza, they hold the **house address where pizza lives**.

---

### Key Symbols

* `&` → **Address-of operator** → gets the memory address.
* `*` → **Dereference operator** → accesses the value at that address.

---

### Example

```go
package main

import "fmt"

// Function takes a pointer to an int
func printNum(num *int) {
    *num = 5 // change the value at that memory address
    fmt.Println("change", *num)
}

func main() {
    num := 1
    fmt.Println("before", num)

    // &num passes the memory address of num
    printNum(&num)

    fmt.Println("after", num)
}
```

---

### Output

```
before 1
change 5
after 5
```

---

### 💡 Explanation

1. `num := 1` → normal variable (holds value `1`).
2. `&num` → gives the memory address of `num`.
3. `printNum(&num)` → passes that address into the function.
4. Inside `printNum`, `*num = 5` changes the **actual value in memory**, so original `num` also updates.

---

### 🔑 Takeaways

* Use pointers when you want to **modify the original variable** inside functions.
* Without pointers, Go passes values by **copy** (like JS primitives).
* Think of `&` as “give me the address” and `*` as “open the box at this address.”



# ⚡ TL;DR Key Takeaways

* Go only has `for` loop but in 3 flavors.
* Arrays = fixed size. Slices = dynamic arrays (use these 99% of the time).
* Maps work like JS objects, but safer.
* Functions can return multiple values + accept variadic args.
* Type switches = Go’s way of handling polymorphism.
