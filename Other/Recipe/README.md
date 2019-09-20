## Recipe of Go

### Anonymous Functions

- Go supports anonymous functions, which can form closures.
- Anonymous functions are useful when you want to define a function inline without having to name it.

```go
package main

import "fmt"

func Equals(a int) func(int) bool {
	return func(x int) bool {
		return a == x
	}
}

func main() {
	fmt.Printf("%t\n", Equals(1)(1))
}

```

Anonymous Functions

We can assign function return to a variable.

```go
package main

import "fmt"

func main() {
        dump := func(str string) string {
		return str
	}

	fmt.Println(dump("Hi anonymous!"))
}
```

### Higher-order Functions

A function:

- Can take another function as an argument.
- Define a function as a return type.

```go
package main

import (
"fmt"
"strings"
)

func hello(hi func(string) string) func(string) string {

    return func(name string) string {
    	return "Hello, " + hi(name)
    }

}

func main() {
    fmt.Println(hello(func(name string) string { return strings.ToUpper(name) })("John Doe"))
}
```

### Function Collections

We can use functions in all the same places that we can use regular data types

```go
package main

import "fmt"

type binaryFunctions func(int, int) int

func main() {

	fns := []binaryFunctions{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
		func(x, y int) int { return x % y },
	}

	x, y := 12, 5

	for i := 0; i < len(fns); i++ {
		fmt.Println(fns[i](x, y))
	}
}
```

### Functions in Struct

We can use type of struct to be container of functions.

```go
package main

import "fmt"

type operations struct {
    name string
    function func(int, int) int
}

func main() {
    ops := []operations{
        {"add", func(x, y int) int { return x + y }},
        {"sub", func(x, y int) int { return x - y }},
        {"mul", func(x, y int) int { return x * y }},
        {"div", func(x, y int) int { return x / y }},
        {"mod", func(x, y int) int { return x % y }},
    }

    x, y := 12, 5

    for i:=0; i < len(ops); i++ {
    	fmt.Printf("%s %d %d = %d\n", ops[i].name, x, y, ops[i].function(x, y))
    }
}
```

### Functions in Map

We can use map to be container of functions.

```go
package main

import "fmt"

var operations map[string]func(int, int) int

func init() {
	operations = make(map[string]func(int, int) int)
	operations["add"] = func(x, y int) int { return x + y }
	operations["sub"] = func(x, y int) int { return x - y }
	operations["mul"] = func(x, y int) int { return x * y }
	operations["div"] = func(x, y int) int { return x / y }
	operations["mod"] = func(x, y int) int { return x % y }
}

func main() {
	x, y := 12, 5

	fmt.Println(operations["add"](x, y))
	fmt.Println(operations["sub"](x, y))
	fmt.Println(operations["mul"](x, y))
	fmt.Println(operations["div"](x, y))
	fmt.Println(operations["mod"](x, y))
}
```

### Tail Recursion

- Tail recursion is recursion where the function calls itself at the end ("tail") of the function.
- Go currently does not optimize for recursion.

```go
package main

import "fmt"

func sum(n, accumulator int) int {
	if n == 1 {
		return accumulator
	} else {
		return sum(n-1, n+accumulator)
	}
}

func sigma(number int) int {
	return sum(number, 1)
}

func main() {
	fmt.Println(sigma(10000000))
}
```

Tail Recursion

We can avoid blowing out memory by using channel

```go
package main

import "fmt"

func sum(number int, accumulator int, result chan int) {
	if number == 1 {

		result <- accumulator
		return
	}

	go sum(number-1, accumulator+number, result)
}

func main() {

	result := make(chan int)

	sum(10000000, 0, result)

	fmt.Println(<-result)
}
```

### Function Types as Interface Values

Function types in Go can also have methods.

```go
package main

import (
        "fmt"
)

type binaryFuction func(int, int) int

func Add(x, y int) int {
        return x + y
}

func (f binaryFuction) Name() string {
        return "Add"
}

func main() {
        fmt.Println(binaryFuction(Add)(4, 5))
        fmt.Println(binaryFuction(Add).Name())
}
```

### Channels of Functions

We can even make channels of functions.

```go
package main

import (
	"fmt"
)

func Produce(channel chan func(), functions []func()) {
	for i := 0; i < len(functions); i++ {
		channel <- functions[i]
	}
	close(channel)
}

func main() {
	x := 4
	functions := []func(){
		func() { x += 1 },
		func() { x -= 1 },
		func() { x *= 2 },
		func() { x /= 2 },
		func() { x *= x },
	}

	channel := make(chan func())

	go Produce(channel, functions)

	for fn := range channel {
		fn()
		fmt.Println(x)
	}
}
```

### Multiple Return Values

A function can return multiple values as output.

```go
package main

import "fmt"

func SalaryLevel(score float64) (string, string) {
	var (
		skill string
		salary string
	)

	switch {
	case score >= 80:
		skill = "dewa"
		salary = "extra ordinary"

	default:
		skill = "so-so"
		salary = "average"
	}

	return skill, salary
}

func main() {
	result1, result2 := SalaryLevel(99)
	fmt.Println("Skill: ", result1, "\nSalary: ", result2)
}
```

### Named Return Parameters

The return of a function can be given names and used as regular variables.

```go
package main

import "fmt"

func MultipleValues(score float64) (levelInt int, levelString string) {
	switch {
	case score <= 33:
		levelInt = 1
		levelString = "1st"

	case score <= 66:
		levelInt = 2
		levelString = "2nd"

	default:
		levelInt = 3
		levelString = "3th"
	}

	return
}

func main() {
	resultInt, resultString := MultipleValues(66.5)
	fmt.Println(resultInt, resultString)
}
```

### Closures

A closure is a function value that references variables from outside its body.

```go
package main

import "fmt"

func makeEvenGenerator() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
```

### Monoids

- 'mappend', an associative binary operation that can reduced two elements to one.
- If we don’t have an element and we need one we can start with the 'identity' element.
- 'mappend' is associative so 'mappend(mappend(a, b), c) == mappend(a, mappend(b, c))' with any values in a, b, and c.
- 'mappend(identity, a) == a' with any value for a.

```go
package main

import (
	"fmt"
	"log"
)

// sum monoid
var identity int = 0
var mappend = func(a, b int) int {
	return a + b
}

func main() {

	// n could be replaced with any value of type int
	n := 42
	identityVerified := mappend(identity, n) == n && mappend(n, identity) == n
	if !identityVerified {
		log.Fatal("invalid identity")
	}

	// a, b, and c could be any values of our given type
	a := 1
	b := 9
	c := 40
	associativityVerified := mappend(mappend(a, b), c) == mappend(a, mappend(b, c))
	if !associativityVerified {
		log.Fatal("not associative")
	}

	fmt.Println("You've got a monoid")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := identity
	for _, n := range numbers {
		sum = mappend(sum, n)
	}

	fmt.Printf("The sum of %v is %d\n", numbers, sum)
}
```

### Compose

A variable or a function could be composed by some [composed] functions.

```go
package main

import "fmt"

/*F is type of function with an integer as return*/
type F func(i int) int

/*A compose function*/
func (f F) compose(inner F) F {
	return func(i int) int { return f(inner(i)) }
}

func main() {
	var f1 F = func(i int) int {
		return i * 2
	}

	var f2 F = func(i int) int {
		return i + 1
	}

	var f3 F = func(i int) int {
		return i + 10
	}

	f := f1.compose(f2).compose(f3)

	fmt.Println(f(4))
}
```

### Lazy Evaluation

Lazy evaluation delays the evaluation of the expression until it is needed

```go
package main

import "fmt"

func Interge() <-chan int {
    yield := make(chan int)
    count := 0
    go func() {
        for {
            yield <- count
            count++
        }
    }()
    return yield
}

var resume <-chan int

func getInetge() int {
    return <-resume
}

func main() {
    resume = Interge()
    fmt.Println(getInetge())
    fmt.Println(getInetge())
    fmt.Println(getInetge())
    fmt.Println(getInetge())
    fmt.Println(getInetge())
    fmt.Println(getInetge())
}
```

original source from https://slides.com/handarusakti/fun-go
