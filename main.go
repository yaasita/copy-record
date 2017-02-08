package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// in fp
	fp_in, err := os.Open("in.txt")
	check(err)
	defer fp_in.Close()
	// out fp
	fp_out, err := os.Create("out.txt")
	check(err)
	defer fp_out.Close()
	// read & write
	scanner := bufio.NewScanner(fp_in)
	for scanner.Scan() {
		str := scanner.Text()
		str_slice := strings.Split(str, "\t")
		loop, err := strconv.Atoi(str_slice[2])
		check(err)
		for i := 1; i <= loop; i++ {
			fp_out.Write([]byte(str + "\r\n"))
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
