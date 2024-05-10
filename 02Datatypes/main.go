package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
**  Almost all the values are types in Golnag similar to Java.
**
	1. String
	2. Boolean
	3. int
	4. Float
	5. complext number

	Advanced
	1. Array
	2. Slices
	3. Maps
	4. Structs
	5. Pointers

	Additional
	1. Functions
	2. Channels

*/

func main() {
	var username string = "name"

	fmt.Println(username)
	fmt.Printf(" variable is of type: %T \n", username)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input1, _ := strconv.ParseFloat(strings.TrimSpace(input), 32)
	fmt.Println(input1)
}

func ArrayFunc() {
	var arr [5]string
	var arr2 = [5]string{""}
	arr[0] =  "1"
	_ = len(arr)
	_ = len(arr2)
 }

 func SlicesFunc() {
	var slice = []int{}
	slice = append(slice, 1,2)
	slice = slice[1:]

 }
