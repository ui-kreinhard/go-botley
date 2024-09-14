package command

var Backward Burst = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero},
}

var Rotate90 Burst = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, Zero},
}
