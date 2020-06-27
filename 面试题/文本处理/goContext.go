package main

import (
	"fmt"
	"regexp"
	"strings"
)

func squareSum(n int, arr [] int) int {
	var num int = int(arr[n-1]);
	
	if(n == 1){
		return int(arr[0]) * int(arr[0]);
	}else {
		if(num >= 0){
			return num * num * squareSum(n-1, arr);
		}else{
			return squareSum(n-1, arr);
		}
	}
}

func list(m) {
	if(m+1 > testfile0.length) {
		if(count == total){
			return 'successful!';
		}else{
			return 'error!';
		}
	}
	
	if( (m+1) % 2 == 0){
		re3, _ := regexp.Compile('\s');
		
		text := re3.ReplaceAllString(testfile0[m+1], ',');
		
		arr := strings.Split(re3, ',');
		
		var sum int = squareSum(int(testfile0[m]), arr);
		
		count ++;
		fmt.print(sum);
	
	}
	
	m++;
	return list(m);

}


func main() {
	
	file := os.Open("./abc.txt");
	defer file.Close();
	
	testfile0 := strings.Split(file, '\r\n');
	
	total := testfile0[0];
	count := 0;
	m := 0;
	list(m);
}