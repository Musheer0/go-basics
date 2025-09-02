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

---

## 🔌 Interfaces in Go

Interfaces in Go = **contracts**.
They define **what methods must exist**, not how they’re implemented.
Think: “If you sign this contract, you MUST have these functions.”

---

### Defining an Interface

```go
type ipayment interface {
    pay(amt float32)
    refunc(amt float32, accnt string)
}
```

* Any type that implements **all methods** in `ipayment` → is considered an `ipayment`.
* No need to explicitly say `implements` (like Java/C#). Go does it **implicitly**.

---

### Using Interfaces in Structs

```go
type payment struct {
    gateway ipayment
}

func (p payment) paymoney(amt float32) {
    p.gateway.pay(amt)
}
```

* `payment` struct has a `gateway` field of type `ipayment`.
* This means it can accept **any type** that satisfies the interface (PayPal, Razorpay, etc.).

---

### Implementation Example 1: PayPal

```go
type paypal struct{}

func (p paypal) pay(amt float32) {
    fmt.Println("method: paypal", amt)
}

func (p paypal) refunc(amt float32, acct string) {
    fmt.Println("refund\nmethod paypal amount:", amt, "account id", acct)
}
```

* Implements both `pay` and `refunc` → ✅ satisfies `ipayment`.

---

### Implementation Example 2: Razorpay

```go
type razorpay struct{}

func (p razorpay) pay(amt float32) {
    fmt.Println("method: razorpay", amt)
}
```

⚠️ `razorpay` **does not implement** `refunc`.
→ So it **cannot be used as `ipayment`** (compiler error if you try).

---

### Using It All Together

```go
func main() {
    newpayment := payment{
        gateway: paypal{}, // inject PayPal
    }
    newpayment.paymoney(45.66)

    // Can also directly call methods from gateway
    newpayment.gateway.refunc(33.22, "user_e3jd3s2")
}
```

#### Output:

```
method: paypal 45.66
refund
method paypal amount: 33.22 account id  user_e3jd3s2
```

---

### 💡 Key Takeaways

* Interfaces = contracts (define *what* methods must exist).
* Structs don’t declare they “implement” → compiler checks if methods match.
* This enables **polymorphism**: same function (`paymoney`) can work with different payment gateways.
* If a struct is missing even one method → ❌ it doesn’t satisfy the interface.
* Interfaces + struct embedding = Go’s version of flexible OOP.


---
## 🚦 Enums with `const` and `iota` in Go

Go doesn’t have native **enum** types like other languages, but we can simulate them using `const` + custom types.

---

### Defining a Custom Type + Constants with `iota`

```go
type status int

const (
    online status = iota
    offline
    busy
    disabled
)
```

* `type status int` → creates a **new type** (based on `int`).
* `iota` → auto-incrementing counter, resets to `0` for each `const` block.
* So here:

  * `online = 0`
  * `offline = 1`
  * `busy = 2`
  * `disabled = 3`

---

### Using It

```go
func printStatus(s status) {
    fmt.Println("user is", s)
}

func main() {
    printStatus(online)
}
```

#### Output:

```
user is 0
```

⚠️ By default, printing gives the **int value** (`0`, `1`, …).

---

### Alternative: String Constants

```go
type status string

const (
    online   status = "online"
    offline  status = "offline"
    busy     status = "busy"
    disabled status = "disabled"
)
```

* Here `status` is a string type.
* Printing gives **human-readable values**.

---

### 💡 Key Takeaways

* `iota` is great for **numeric enums** (fast, efficient).
* String constants are better when you need **readable logs or debugging**.
* Always wrap constants in a custom type (`status`) → improves type safety.

## 🌀 Generics in Go

Generics let you write **functions/structs** that can work with **any type** without duplicating code.

---

### Generic Functions

```go
func genericType[T comparable](items []T) {
    for i, t := range items {
        fmt.Println(i, t)
    }
}
```

* `[T comparable]` → type parameter.

  * `T` is a placeholder for any type.
  * `comparable` constraint → only types that support `==` and `!=`.
  * Common constraints:

    * `any` → literally any type.
    * `comparable` → ints, strings, bools, etc.
    * Union constraints → e.g., `[T int | float64]`.

Usage:

```go
genericType([]int{3, 2, 1})
genericType([]string{"m", "u"})
```

---

### Generic Structs

```go
type genericStruct[T any] struct {
    items []T
}
```

* Defines a struct where `items` can hold **any type**.
* When creating it, you decide the type:

```go
mg := genericStruct[string]{
    items: []string{"m", "u"},
}
```

---

### Alternative Syntax with Union Types

```go
// Only allow bool or string
func genericType[T bool | string](items []T) {
    for i, t := range items {
        fmt.Println(i, t)
    }
}
```

---

### 💡 Key Takeaways

* `[T any]` → universal placeholder for any type.
* Add constraints to restrict what types can be passed.
* Generics reduce **code duplication** (no need to write separate versions for `int`, `string`, etc.).
* Great for reusable **data structures** (stacks, queues, trees) and algorithms.


## ⚡ Goroutines + WaitGroups

Go handles concurrency using **goroutines** (lightweight threads) and sync tools like **WaitGroup** to make sure things finish before the program exits.

---

### Goroutines

```go
go task(i, &wg)
```

* `go` → runs `task` in a **separate goroutine**.
* Super lightweight (not like OS threads). You can spawn thousands without choking.
* But: once `main()` ends → all goroutines die instantly. That’s why we need a **WaitGroup**.

---

### WaitGroup

```go
var wg sync.WaitGroup
```

* `wg.Add(1)` → tells the WaitGroup *“one more goroutine to wait for.”*
* `wg.Done()` → signals *“this goroutine finished.”*
* `wg.Wait()` → blocks the main function until all added goroutines are done.

---

### Full Flow

```go
func task(i int, wg *sync.WaitGroup) {
    defer wg.Done() // always mark completion
    fmt.Println(i)
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i <= 10; i++ {
        wg.Add(1)
        go task(i, &wg)
    }
    wg.Wait() // wait until all goroutines finish
}
```

---

### Output

Order is **not guaranteed** because goroutines run concurrently:

```
3
1
0
5
2
...
```

---

### 💡 Key Takeaways

* `go funcName()` → starts a goroutine.
* Use `sync.WaitGroup` to stop `main` from exiting early.
* Always pair `wg.Add(1)` with a `wg.Done()` inside the goroutine.
* Execution order is random → don’t rely on it unless you use sync primitives (channels, mutex, etc.).


## 📡 Channels in Go

Channels are like **pipes** that goroutines can use to send data back and forth. They’re Go’s built-in concurrency communication tool.

---

### 🔑 Creating a Channel

```go
mychan := make(chan int) // unbuffered channel
```

* `chan int` → channel that only carries `int`s.
* By default, channels are **unbuffered** (send blocks until receive is ready).

---

### 📨 Sending and Receiving

```go
mychan <- 5      // send value into channel
val := <-mychan  // receive value from channel
```

* If no one’s on the other end, it blocks (waits).

---

### Example: Using Channels Instead of `WaitGroup`

```go
func sendMail(mailchan chan bool) {
    defer func() { mailchan <- true }()
    fmt.Println("email sent")
}

func main() {
    emailChan := make(chan bool)
    go sendMail(emailChan)
    <-emailChan // wait for signal
}
```

* Perfect for **waiting on one goroutine**.

---

### Example: Passing Data with Channels

```go
func add(numChan chan int, n1 int, n2 int) {
    numChan <- n1 + n2
}

func main() {
    mychan := make(chan int)
    go add(mychan, 4, 5)
    res := <-mychan
    fmt.Println(res) // 9
}
```

---

### ⚙️ Closing Channels

```go
close(emailChan)
```

* Closes a channel → signals that **no more values** will be sent.
* Needed for `range` over a channel, otherwise it blocks forever.

---

### Buffered Channels

```go
emailChan := make(chan string, 10) // holds up to 10 messages
```

* Doesn’t block until buffer is full.
* Useful when you want async queuing.

---

### ⏳ `select` for Multiple Channels

```go
chan1 := make(chan int)
chan2 := make(chan string)

go func() { chan1 <- 10 }()
go func() { chan2 <- "thee" }()

for i := 0; i < 2; i++ {
    select {
    case val1 := <-chan1:
        fmt.Println(val1)
    case val2 := <-chan2:
        fmt.Println(val2)
    }
}
```

* `select` waits for whichever channel is ready first.
* Think of it like a **switch-case for channels**.

---

### 🔒 Directional Channels

```go
func receiveOnly(em <-chan string) {
    msg := <-em // ✅ only receive allowed
}

func sendOnly(em chan<- string) {
    em <- "hi"  // ✅ only send allowed
}
```

* `<-chan T` → **receive-only**.
* `chan<- T` → **send-only**.
* Adds **type safety** → prevents accidental misuse.

---

### 💡 Key Takeaways

* Channels let goroutines communicate safely.
* Unbuffered = sync handoff, buffered = async queue.
* Always `close()` channels when you’re done sending.
* `select` is your friend for multiplexing multiple channels.
* Directional channels (`<-chan`, `chan<-`) are great for clean APIs.

## 🔒 Mutex (Mutual Exclusion) in Go

When multiple goroutines try to update the same variable at the same time → **race condition** happens.
A `sync.Mutex` (lock) prevents that by allowing **only one goroutine** to access the critical section at a time.

---

### Example: Safe Counter with Mutex

```go
package main

import (
	"fmt"
	"sync"
)

type test struct {
	sum int
	mu  sync.Mutex
}

func (p *test) add(wb *sync.WaitGroup) {
	defer func() {
		wb.Done()
		p.mu.Unlock() // unlock after done
	}()

	p.mu.Lock()     // lock before modifying shared data
	p.sum += 1
}

func main() {
	var wb sync.WaitGroup
	p := test{sum: 0}

	for i := 0; i <= 100; i++ {
		wb.Add(1)
		go p.add(&wb)
	}

	wb.Wait()
	fmt.Println(p.sum)
}
```

---

### 🛠️ Breakdown

* `mu sync.Mutex` → the lock itself.
* `p.mu.Lock()` → blocks other goroutines until we’re done.
* `p.mu.Unlock()` → releases lock so next goroutine can enter.
* Always `defer Unlock()` so you don’t forget to release.

---

### ⚠️ Without Mutex

```go
p.sum += 1
```

If many goroutines run this at once → some increments get lost, final `sum` < expected.

---

### ✅ With Mutex

* Only one goroutine can increment at a time.
* Ensures `sum` ends up exactly `101`.

---

### 🔑 Key Takeaways

* Use `sync.Mutex` for safe access to **shared state**.
* Always lock → update → unlock.
* Combine with `sync.WaitGroup` to wait for goroutines to finish.
* Forgetting `Unlock` = **deadlock** (everything freezes).

---

Want me to also add the **`sync.RWMutex`** version (for read-heavy code where multiple reads can happen in parallel but writes still lock)?

