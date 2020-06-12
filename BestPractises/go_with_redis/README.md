## Go with Redis

### Install the Redis Driver

```shell    
go get github.com/go-redis/redis
```

file main.go

```shell
#run file
go run Best\ Practises/go_with_redis/main.go
```

```go
package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Run ...
func Run() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err1 := client.Ping().Result()
	fmt.Println(pong, err1)

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist

}

// main function
func main() {
	Run()
}

```