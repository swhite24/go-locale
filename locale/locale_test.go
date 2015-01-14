package locale

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	ls := Read("en-US,en;q=0.8")
	fmt.Println(ls)

	l := ls.Best()
	fmt.Println(l)
}
