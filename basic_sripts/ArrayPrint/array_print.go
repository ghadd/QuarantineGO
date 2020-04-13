package arrayprint

import (
	"fmt"
)

func main() {
	arr := make([]int, 5)

	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
