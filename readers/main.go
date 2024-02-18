package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Nyanpasa")

	arr := make([]byte, 4)

	for {
		n, err := r.Read(arr)
		fmt.Printf("n = %d, err = %v, arr= %v\n", n, err, arr)
		fmt.Printf("arr n bytes: %q", arr[:n])
		if err == io.EOF {
			break
		}
	}
}
