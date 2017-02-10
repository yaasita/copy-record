package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var in_count, out_count int
	in_count = 0
	out_count = 0
	// in fp
	fp_in, err := os.Open("in.txt")
	check(err, "ファイルが見つかりません")
	defer fp_in.Close()
	// out fp
	fp_out, err := os.Create("out.txt")
	check(err, "書き込みに失敗しました")
	defer fp_out.Close()
	// read & write
	scanner := bufio.NewScanner(fp_in)
	for scanner.Scan() {
		str := scanner.Text()
		in_count = in_count + 1
		str_slice := strings.Split(str, "\t")
		loop, err := strconv.Atoi(str_slice[4])
		if err != nil {
			loop = 1
		}
		for i := 1; i <= loop; i++ {
			fp_out.Write([]byte(str + "\r\n"))
			out_count = out_count + 1
		}
	}
	var res string
	fmt.Println("処理終了")
	fmt.Printf("入力件数: %d\n", in_count)
	fmt.Printf("出力件数: %d\n", out_count)
	fmt.Println("何かキーを押してください")
	_, err = fmt.Scanln(&res)
}

func check(e error, s string) {
	if e != nil {
		fmt.Fprintln(os.Stderr, s)
		var res string
		fmt.Println("何かキーを押してください")
		_, _ = fmt.Scanln(&res)
		panic(e)
	}
}
