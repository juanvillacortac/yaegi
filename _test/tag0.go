// The following comment line has the same effect as 'go run -tags=dummy'
// yaegi:tags dummy

package main

import _ "github.com/juanvillacortac/yaegi/_test/ct"

func main() {
	println("bye")
}

// Output:
// hello from ct1
// hello from ct3
// bye
