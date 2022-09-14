package iteration

import "fmt"

func Repeat(character string, repeat int) string {
	var repeated string

	for i := 0; i < repeat; i++ {
		repeated += character
	}
	return repeated
}

func main() {
	fmt.Printf(Repeat("a", 10))
}
