package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/ui-kreinhard/go-botley/command"
	"github.com/ui-kreinhard/go-botley/decode"
)

type Command struct {
	content []int
}

func (c *Command) render(commandCounter byte) []int {
	endOfBurst := 18048
	mark := 642
	return append(append(c.content, encodeByte(commandCounter)...), mark, endOfBurst)
}

type EndOfCommand struct {
	content []int
}

func (e *EndOfCommand) render(totalCommandCounter byte) []int {
	return append(append(e.content, encodeByte(totalCommandCounter)...), 642)
}

func encodeByte(value byte) []int {
	ret := []int{}
	for i := 0; i < 8; i++ {

		if value&1 == 1 {
			ret = append(ret, 1284)
		} else {
			ret = append(ret, 642)
		}
		ret = append(ret, 642)
		value = value >> 1
	}
	slices.Reverse(ret)
	return ret
}

func main() {
	// fmt.Println("encopde byte",encodeByte(2))
	preamble := []int{4255, 692, 629, 695, 630, 1330, 629, 1333, 679, 1282, 651, 1311, 648, 1329, 630, 1328, 631, 694, 631, 676, 649, 677, 648, 674, 651, 671, 654, 670, 653, 668, 631, 695, 630, 24783, 4254, 693, 630, 696, 630, 1329, 629, 1330, 630, 1334, 652, 1310, 649, 1312, 647, 1328, 631, 679, 646, 677, 648, 675, 650, 670, 654, 670, 653, 670, 629, 695, 630, 695, 631, 25139, 4260, 692, 628, 696, 629, 1333, 651, 1311, 678, 1306, 653, 1303, 656, 1301, 658, 1302, 629, 694, 658, 667, 658, 672, 653, 649, 676, 649, 676, 646, 678, 645, 652, 669, 631, 25196, 4256, 695, 626, 697, 655, 670, 655, 669, 656, 652, 673, 648, 677, 1287, 672, 649, 676, 1306, 653, 653, 672, 1304, 655, 1305, 654, 669, 656, 653, 672, 651, 674, 1303, 656, 25750, 4258, 692, 629, 695, 630, 695, 629, 695, 657, 669, 656, 667, 658, 1302, 657, 669, 656, 1301, 658, 667, 658, 1304, 655, 1302, 630, 696, 629, 696, 656, 1301, 631, 695, 657, 25749, 4258, 692, 629, 695, 630, 696, 657, 667, 658, 668, 657, 653, 672, 1303, 656, 668, 657, 1301, 658, 668, 657, 1303, 656, 1302, 630, 696, 656, 668, 657, 1301, 631, 1329, 630, 25142, 4258, 692, 629, 695, 630, 696, 656, 666, 659, 668, 657, 672, 653, 1303, 656, 668, 657, 1303, 656, 669, 656, 1302, 657, 1303, 629, 695, 657, 1303, 629, 695, 630, 696, 657, 25788, 4248, 702, 672, 651, 674, 651, 674, 647, 678, 644, 679, 643, 630, 1334, 679, 643, 629, 1335, 677, 645, 650, 1312, 677, 1283, 676, 651, 674, 1287, 672, 647, 678, 1286, 673, 25085}

	endOfCommand := EndOfCommand{
		content: []int{4283, 670, 626, 698, 627, 695, 629, 697, 628, 1331, 628, 1333, 651, 1311, 676, 1308},
	}

	backward := Command{
		content: []int{4256, 690, 631, 696, 629, 696, 655, 669, 655, 671, 654, 651, 674, 1305, 654, 671},
	}
	r90 := Command{
		content: []int{4259, 691, 631, 695, 658, 667, 658, 653, 672, 649, 676, 1306, 653, 650, 675, 649},
	}
	fmt.Println(decode.DecodeTiming(preamble))
	fmt.Println(decode.DecodeTiming(endOfCommand.content))
	fmt.Println(decode.DecodeTiming(backward.content))
	fmt.Println(decode.DecodeTiming(r90.content))

	// output(preamble, backward.render(1), r90.render(2), endOfCommand.render(3))

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
