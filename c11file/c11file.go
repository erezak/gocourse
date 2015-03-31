package main

import (
	"fmt"
	"github.com/erezak/gocourse/c11file/c11file"
	"strconv"
)

func iToFormattedA(number int64) string {
	formattedNumber := strconv.FormatInt(number, 10)
	for i := len(formattedNumber); i > 3; {
		i -= 3
		formattedNumber = formattedNumber[:i] + "," + formattedNumber[i:]
	}
	return formattedNumber
}

func main() {
	size, err := c11file.SumFilesSizes("C:\\")
	if err != nil {
		fmt.Println(err.Error)
	} else {
		fmt.Printf("Files total size in %s: %smb (%s bytes)\n", "C:\\",
			iToFormattedA(size/1024/1024), iToFormattedA(size))
	}
}
