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

---

## 🏗️ Structs in Go

Structs = **blueprints** for grouping related data.
Think: JSON object but with type safety.

---

### Defining a Struct

```go
type order struct {
    id          string
    amount      float32
    status      string
    created_at  time.Time
    customer_id string
}
```

---

### Adding Methods to Structs

```go
func (o *order) changeStatus(status string) {
    o.status = status
}
```

* `func (o *order)` → receiver function (like methods in OOP).
* Pointer `*order` ensures the original struct gets updated.

---

### Constructor-like Function (Go has no real constructors)

```go
func newOrder(o order) *order {
    myorder := order{
        id:         o.id,
        status:     o.status,
        created_at: o.created_at,
        amount:     o.amount + (o.amount * 0.15), // add 15% tax
        customer_id:"user_bc328fca-4ec2-48d7-b313-206dbe7b9d99", // from session/uuid
    }
    return &myorder
}
```

* Go doesn’t have constructors → we **duct-tape it** with helper functions (prefix `new` for convention).

---

### Nested Structs

```go
type customer struct {
    name           string
    age            int
    is_plus_member bool
    orders         []order
    order_count    int
}
```

---

### Initializer for Nested Struct

```go
func newCustomer(c customer) *customer {
    mycustomer := c
    mycustomer.order_count = len(c.orders)
    return &mycustomer
}
```

---

### Inline Structs (for one-time use)

```go
inlineStruct := struct {
    name string
    age  int
}{"musheer", 17}
fmt.Println(inlineStruct)
```

* Useful when you need a quick struct without defining a type.

---

### Example in Action

```go
initializedOrder := *newOrder(order{
    id:        "23DQ",
    amount:    23.4,
    status:    "otw",
    created_at: time.Now(),
})

initializedOrder.changeStatus("out for delivery")

mycustomer := *newCustomer(customer{
    name:           "musheer",
    is_plus_member: true,
    orders:         []order{initializedOrder},
})

fmt.Println(mycustomer)
```

---

### 💡 Key Takeaways

* Structs = Go’s way of modeling objects.
* Methods use **receiver functions** (pointer receiver if you wanna mutate data).
* No built-in constructors → use custom `newXxx` functions.
* Nested structs let you embed relationships.
* Inline structs are great for quick, temporary shapes.
