//This is a comment

package main

import (
	"fmt"
	"time"
	"math/rand"
)


func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Hello, this is the n-th, boring, hello world app...")
	fmt.Println("I tell you what time is it, because i'm kind: ", time.Now())
	fmt.Println("And a random number: ", rand.Intn(10))
}
