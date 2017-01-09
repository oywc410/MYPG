package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//範囲
var m, n int

func main() {

	//島
	arr := make(map[int](map[int]int))
	//位置確認判断用
	arrCheck := make(map[int](map[int]bool))

	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()
	constr := string(line)

	nm := strings.Split(constr, " ")

	n, _ = strconv.Atoi(nm[0])
	m, _ = strconv.Atoi(nm[1])

	j := 0
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		constr := string(line)

		if constr == "\\q" {
			break
		}

		bs := strings.Split(constr, " ")
		arr[j] = make(map[int]int)
		arrCheck[j] = make(map[int]bool)

		for i, t := range bs {
			arr[j][i], _ = strconv.Atoi(t)
			arrCheck[j][i] = false
		}

		j++

		if j == m {
			break
		}
	}

	//fmt.Println("-----------")

	//fmt.Println(n, m)
	//fmt.Println(arr)

	//fmt.Println("-----------")

	//島数
	con := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if arr[i][j] == 0 {
				arrCheck[i][j] = true
			}

			if arrCheck[i][j] == false {
				//島あり判断していない場合は島の位置全部「判断済」に設定する
				setSiMa(&arrCheck, arr, i, j)
				//fmt.Println(arrCheck, i, j)
				con++
			}
		}
	}

	fmt.Println(con)
}

//島判設定用
func setSiMa(arrCheckP *map[int](map[int]bool), arr map[int](map[int]int), i, j int) bool {

	if arr[i][j] == 0 {
		return false
	}

	arrCheck := *arrCheckP

	if arrCheck[i][j] == true {
		return true
	}

	arrCheck[i][j] = true

	//各方法関連している位置を「判断済」を設定する
	if i-1 >= 0 {
		setSiMa(arrCheckP, arr, i-1, j)
	}

	if j+1 < n {
		setSiMa(arrCheckP, arr, i, j+1)
	}

	if i+1 < m {
		setSiMa(arrCheckP, arr, i+1, j)
	}

	if j-1 >= 0 {
		setSiMa(arrCheckP, arr, i, j-1)
	}

	return true
}
