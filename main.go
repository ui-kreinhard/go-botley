package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ui-kreinhard/go-botley/command"
)

func main() {
	program := command.Program{
		Name: "botleyTest",
		Bursts: []command.Burst{
			command.Backward,
			command.Rotate90,
			command.Rotate90,
		},
	}
	outputFlipperIrFile("/tmp/botley_synth.ir", program)
}

func outputFlipperIrFile(path string, program command.Program) error {

	raw := program.Render()

	rawString := ""
	for _, rawPart := range raw {
		rawString += strconv.Itoa(rawPart) + " "
	}

	dataParts := []string{
		"Filetype: IR signals file",
		"Version: 1",
		"#",
		"name: " + program.Name,
		"type: raw",
		"frequency: 38000",
		"duty_cycle: 0.330000",
		"data: " + rawString,
	}
	fmt.Println(strings.Join(dataParts, "\n"))

	os.Truncate(path, 0)
	return os.WriteFile(path, []byte(strings.Join(dataParts, "\n")), 0644)
}
