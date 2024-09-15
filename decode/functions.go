package decode

import "fmt"

type Part string

const (
	StartOfBurst Part = "StartOfBurst"
	EndOfBurst   Part = "EndOfBurst"
	Mark         Part = "Mark"
	Zero         Part = "Zero"
	One          Part = "One"
)

func DecodeTiming(timing []int) (ret []Part) {
	for i, timingElement := range timing {
		if timingElement > 15000 {
			ret = append(ret, EndOfBurst)
			continue
		}
		if timingElement > 4000 {
			ret = append(ret, StartOfBurst)
			continue
		}

		if (i+1)%2 == 1 {
			ret = append(ret, Mark)

		} else {
			fmt.Println(timingElement)
			if timingElement >= 1000 {

				ret = append(ret, One)
			} else {
				ret = append(ret, Zero)
			}
		}
	}
	return
}

func ExtractSingleCommand(timings []int) []Part {
	d := DecodeTimingWithoutPreambleAndSkipEnd(timings)
	return d[0:16]
}

func DecodeTimingWithoutPreambleAndSkipEnd(timings []int) []Part {
	d := DecodeTimingWithoutPreamble(timings)
	return d[:len(d)-1]
}

func DecodeTimingWithoutPreamble(timings []int) []Part {
	return DecodeTiming(timings[272:])
}
