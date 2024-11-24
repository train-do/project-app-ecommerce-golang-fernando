// package main

// import "fmt"

// func adder() func(int) int {
// 	sum := 0
// 	return func(x int) int {
// 		sum += x
// 		return sum
// 	}
// }

//	func main() {
//		pos, neg := adder(), adder()
//		for i := 0; i < 10; i++ {
//			fmt.Println(
//				pos(i),
//				neg(i),
//			)
//		}
//	}
package main

import (
	"log"
	"net/http"

	"github.com/train-do/project-app-ecommerce-golang-fernando/router"
)

// import "fmt"

// func outer() func() int {
// 	count := 0 // variabel yang ada di dalam scope outer
// 	fmt.Println(count, "----")
// 	return func() int {
// 		fmt.Println(count, "****")
// 		count++ // inner function dapat mengakses count
// 		fmt.Println(count, "++++")
// 		return count
// 	}
// }

//	func main() {
//		closureFunc := outer()     // Memanggil outer yang mengembalikan fungsi inner
//		fmt.Println(closureFunc()) // Output: 1
//		fmt.Println(closureFunc()) // Output: 2
//		fmt.Println(closureFunc()) // Output: 3
//	}
func main() {
	router, logger, err := router.RouterAPI()
	if err != nil {
		logger.Panic(err.Error())
	}
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
