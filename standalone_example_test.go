package wordcase

import (
	"fmt"
)

func ExampleSnakeCase() {
	fmt.Println(SnakeCase("ONE nineFive"))
	// Output: one_nine_five
}
