package command

var Forward Burst = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One},
}

var Backward Burst = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero},
}

var RotateRight90 Burst = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, Zero},
}

var RotateLeft90 Burst = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, One},
}