package utils

// Get Color by piece code num
func GetColor(code int) string {
	colorMap := map[int]string {
		0: "W",
		1: "Y",
		2: "G",
		3: "B",
		4: "R",
		5: "O",
	}
	return colorMap[code]
}

// get correct face for Vertical turn clockwise
func GetFaceForVerticalTurnClockwise(face int) int {
	switch face {
	case 0:
		return 2
	case 2:
		return 4
	case 4:
		return 5
	case 5:
		return 0
	default:
		return face
	}
}

// get correct face for Vertical turn counterclockwise
func GetFaceForVerticalTurnCounterClockwise(face int) int {

	switch face {
	case 0:
		return 5
	case 2:
		return 0
	case 4:
		return 2
	case 5:
		return 4
	default:
		return face
	}
}

// get correct face for Horizontal turn clockwise
func GetFaceForHorizontalTurnClockwise(face int) int {
	switch face {
	case 1:
		return 5
	case 2:
		return 1
	case 3:
		return 2
	case 5:
		return 3
	default:
		return face
	}
}

// get correct face for Horizontal turn counterclockwise
func GetFaceForHorizontalTurnCounterClockwise(face int) int {
	switch face {
	case 1:
		return 2
	case 2:
		return 3
	case 3:
		return 5
	case 5:
		return 1
	default:
		return face
	}
}



