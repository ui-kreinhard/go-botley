package command

import "slices"

type Part string

const (
	StartOfBurst Part = "StartOfBurst"
	EndOfBurst   Part = "EndOfBurst"
	Mark         Part = "Mark"
	Zero         Part = "Zero"
	One          Part = "One"
)

func (p Part) render() int {
	switch p {
	case StartOfBurst:
		return 4150
	case EndOfBurst:
		return 18042
	case Mark:
		return 642
	case One:
		return 1284
	case Zero:
		return 642
	}
	return 0
}

type Program struct {
	Bursts []Burst
	Name   string
}

func renderPreabmel() (ret []int) {
	parts := []Part{StartOfBurst, Zero, Mark, Zero, Mark, One, Mark, One, Mark, One, Mark, One, Mark, One, Mark, One, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, EndOfBurst, StartOfBurst, Zero, Mark, Zero, Mark, One, Mark, One, Mark, One, Mark, One, Mark, One, Mark, One, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, EndOfBurst, StartOfBurst, Zero, Mark, Zero, Mark, One, Mark, One, Mark, One, Mark, One, Mark, One, Mark, One, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, EndOfBurst, StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, One, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, EndOfBurst, StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, One, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, EndOfBurst, StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, One, Mark, Zero, Mark, Zero, Mark, One, Mark, One, Mark, EndOfBurst, StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, One, Mark, Zero, Mark, One, Mark, Zero, Mark, Zero, Mark, EndOfBurst, StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, One, Mark, Zero, Mark, One, Mark, Zero, Mark, One, Mark, EndOfBurst}
	for _, part := range parts {
		ret = append(ret, part.render())
	}
	return
}

func renderEndOfCommand(commandCount byte) (ret []int) {
	parts := []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, One, Mark, One, Mark, One}

	for _, part := range parts {
		ret = append(ret, part.render())
	}
	ret = append(ret, encodeByte(commandCount)...)
	ret = append(ret, Mark.render())
	return
}

func (p *Program) Render() []int {
	ret := renderPreabmel()
	commandCount := 1
	for i, burst := range p.Bursts {
		renderedBurst := burst.render(byte(i + 1))
		ret = append(ret, renderedBurst...)
		commandCount++
	}
	ret = append(ret, renderEndOfCommand(byte(commandCount))...)
	return ret
}

type Burst struct {
	parts []Part
}

func (b *Burst) render(commandCounter byte) (ret []int) {
	for _, part := range b.parts {
		ret = append(ret, part.render())
	}
	ret = append(ret, encodeByte(commandCounter)...)
	ret = append(ret, Mark.render())
	ret = append(ret, EndOfBurst.render())
	return
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
