package command

var Forward = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One},
}

var Backward = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero},
}

var RotateRight90 = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, Zero},
}

var RotateLeft90 = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, One},
}

var RotateRight45 = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, One, Mark, Zero},
}

var RotateLeft45 = Burst{
	parts: []Part{StartOfBurst, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, Zero, Mark, One, Mark, Zero, Mark, One},
}
