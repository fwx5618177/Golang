package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var output []int

func main() {
	var err1 error
	var testfile0 []string

	// That's the test file name
	// this demo get data from the test.txt
	file, err1 := ioutil.ReadFile("./test.txt")
	if err1 != nil {
		fmt.Println(err1)
	}

	testfile0 = strings.Split(string(file), "\r\n")

	total, _ := strconv.Atoi(testfile0[0])
	count := 0
	m := 0
	list(m, testfile0, count, total)
	// fmt.Println(output)
	outPrint(0, output)
}

func squareSum(n int, arr []string) int {
	// fmt.Println("传进来的参数，代表个数:", n)
	// fmt.Println("此时指定奇数行的值：", arr[n-1])
	num, _ := strconv.Atoi(arr[n-1])

	if n == 1 {
		num1, _ := strconv.Atoi(arr[0])
		// fmt.Println("最后一个的平方:", num1*num1)
		return num1 * num1
	}
	if num >= 0 {
		// fmt.Println("平方:", num*num)
		return num*num + squareSum(n-1, arr)
	}

	return squareSum(n-1, arr)

}

func outPrint(i int, output []int) {

	if i+1 > len(output) {
		return
	}
	// fmt.Println("i", i)
	// fmt.Println("output[i]", output[i])
	fmt.Println(output[i])
	i++
	outPrint(i, output)
	return
}

func list(m int, testfile0 []string, count int, total int) {
	if m+1 >= len(testfile0) {
		// fmt.Println(output)
		return
	}

	if (m+1)%2 == 0 {
		var sum int = 0
		// fmt.Println("m:", m)
		// fmt.Println("testfile0[m+1]:", testfile0[m+1])
		pretext := regexp.MustCompile(`\s`)

		text := pretext.ReplaceAllString(testfile0[m+1], `,`)
		// fmt.Println("得到的数组：", text)
		// fmt.Println("上一行,指定的数字个数:", testfile0[m])

		var arr []string = strings.Split(text, ",")
		num, _ := strconv.Atoi(testfile0[m])
		sum = squareSum(num, arr)

		count++
		// fmt.Println("total：", sum)
		output = append(output, sum)
		// fmt.Println(output)
		// outPrint(0, output)
		// return output
	}

	m++
	list(m, testfile0, count, total)
	// fmt.Println(output)
	return
}
