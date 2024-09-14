package decode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	content := []int{4256, 690, 631, 696, 629, 696, 655, 669, 655, 671, 654, 651, 674, 1305, 654, 671, 654, 651, 674, 648, 677, 648, 651, 672, 627, 696, 629, 696, 628, 696, 655, 1304, 629, 18059}
	decoded := DecodeTiming(content)
	expected := []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, EndOfBurst}
	assert.Equal(t, expected, decoded)
}
