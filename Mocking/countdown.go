package Mocking

import (
	"fmt"
	"io"
	"time"
)

func Countdown( writer io.Writer)  {

	fmt.Fprintf(writer,"3\n")
	time.Sleep(1 * time.Second)
	fmt.Fprintf(writer,"2\n")
	time.Sleep(1 * time.Second)
	fmt.Fprintf(writer,"1\n")
	time.Sleep(1 * time.Second)
	fmt.Fprintf(writer,"Go!")
}
