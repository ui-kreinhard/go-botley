package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ui-kreinhard/go-botley/decode"
)

func convertToInt(input string) (ret []int) {
	parts := strings.Split(input, " ")
	for _, part := range parts {
		value, _ := strconv.Atoi(part)
		ret = append(ret, value)
	}
	return
}

func main() {
	data := convertToInt(os.Args[1])
	fmt.Println(decode.DecodeTimingWithoutPreambleAndSkipEnd(data))
	fmt.Println(convertAndJoin(decode.ExtractSingleCommand(data)))
}

func convertAndJoin(commands []decode.Part) string {
	ret := []string{}
	for _, command := range commands {
		ret = append(ret, string(command))
	}
	return strings.Join(ret, ", ")
}
