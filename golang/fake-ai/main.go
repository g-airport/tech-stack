package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		c, err := in.ReadString('\n')
		if err == nil {
			c = strings.Replace(c,"?","!",-1)
			c = strings.Replace(c,"？","！",-1)
			fmt.Println(c)
		}
	}
}
