package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	line, _, _ := reader.ReadLine()

	n, _ := strconv.Atoi(string(line))

	m := big.NewInt(1)

	tmp := big.NewInt(0)

	tmp10 := big.NewInt(int64(10))
	tmp100x := big.NewInt(int64(1000000000))

	for {

		m.Mul(m, big.NewInt(int64(n)))

		tmp = big.NewInt(m.Int64())

		if tmp.Mod(tmp, tmp10).Int64() == 0 {

			for {
				tmp = big.NewInt(m.Int64())
				if tmp.Mod(tmp, tmp10).Int64() == 0 {
					m.Div(m, tmp10)

				} else {
					break
				}
			}
		}

		m.Mod(m, tmp100x)

		n--
		if n == 1 {
			break
		}
	}

	fmt.Println(m.String())
}
