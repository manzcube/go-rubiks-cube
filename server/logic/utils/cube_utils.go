package utils

import (
	"reflect"
	"server/logic/helpers"
	"server/logic/models"
)

// Tensor combinations
func GenerateCombinations(n int) [][]int {
	// If dimension is 0 or lower, return empty tensor
	if n <= 0 {
		return [][]int{}
	}

	// Create reursive function to generate combinations
	var generate func(int, []int, [][]int) [][]int
	generate = func(dim int, vec []int, result [][]int) [][]int {
		if dim == n {
			// Make a copy of vec to avoid slice referencing issues
			combo := make([]int, n)
			copy(combo, vec)
			result = append(result, combo)
			return result
		}

		// Recursive part when the function is used within itself
		for i := 0; i < n; i++ {
			vec[dim] = i
			result = generate(dim+1, vec, result)
		}


		
		return result
	}

	return generate(0, make([]int, n), [][]int{})
}

// Remove Nucleus piece to avoid trauma
func RemoveNucleus(combinations models.CubeCombinatios) models.CubeCombinatios {
	var newSlice [][]int
	nucleus := []int{1, 1, 1}
	for _, v := range combinations {
		if !reflect.DeepEqual(v, nucleus) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
} 

func GenerateColors() [][]string {
    colors := []string{"White", "Yellow", "Blue", "Red", "Orange", "Green"}
    var combinations [][]string

    var generate func(int, []string)
    generate = func(start int, currentCombination []string) {
        if len(currentCombination) > 0 && len(currentCombination) <= 3 {
            combinations = append(combinations, append([]string(nil), currentCombination...))
        }
        if len(currentCombination) == 3 {
            return // Stop if the current combination is already of length 3
        }

        for i := start; i < len(colors); i++ {
            if helpers.ContainsColor(currentCombination, "White") && colors[i] == "Yellow" || helpers.ContainsColor(currentCombination, "Yellow") && colors[i] == "White" {
                continue // Skip if adding this color violates the restriction
            } else if helpers.ContainsColor(currentCombination, "Green") && colors[i] == "Blue" || helpers.ContainsColor(currentCombination, "Blue") && colors[i] == "Green" {
                continue // Skip if adding this color violates the restriction
            } else if helpers.ContainsColor(currentCombination, "Red") && colors[i] == "Orange" || helpers.ContainsColor(currentCombination, "Orange") && colors[i] == "Red" {
                continue // Skip if adding this color violates the restriction
            }
            generate(i+1, append(currentCombination, colors[i]))
        }
    }

    generate(0, []string{})
    return combinations
}




// Get all valid color types for every piece
func GeneratePieceTypes(combinations models.CubeCombinatios) models.PieceTypes {
	var result []string
	for _, tensor := range combinations {
		if sum(tensor) % 2 != 0 {
			result = append(result, "edge")
		} else if isCenter(tensor) {
			result = append(result, "center")
		} else {
			result = append(result, "corner")
		}
	}
	return result
}

// func isOppositeColor(firstColor string, secondColor string) bool {

// }

func isCenter(tensor []int) bool {
	n := 0
	for _, v := range tensor {
		if v == 1 {
			n++
		}
	}
	if n == 2 {
		return true
	} else {
		return false
	}
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Get Color by piece code num
func GetColor(code int) string {
	colorMap := map[int]string {
		0: "W",
		1: "O",
		2: "G",
		3: "R",
		4: "Y",
		5: "B",
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



