package utils

import (
	"server/logic/helpers"
	"server/logic/models"
	"slices"
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


// Get all valid color types for every piece
func GeneratePieceTypes(combinations models.CubeCombinatios) models.PieceTypes {
	var result []string
	for _, tensor := range combinations {
		if helpers.Sum(tensor) % 2 != 0 {
			result = append(result, "edge")
		} else if helpers.IsCenter(tensor) {
			result = append(result, "center")
		} else {
			result = append(result, "corner")
		}
	}
	return result
}



// Generate all possible color combinations for respective pieces
func GenerateColors() [][]string {
    colors := []string{"White", "Orange", "Blue", "Green", "Red", "Yellow"}
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


// Now we need to assign the colors properly
func OrderColors(positions models.CubeCombinatios, types models.PieceTypes, colors models.PieceColors) models.PieceColors {
	// Empty colors slices array
	var orderedColors models.PieceColors
	centerColors := helpers.GetColorsbyLength(colors, 1) 
	edgeColors := helpers.GetColorsbyLength(colors, 2) 
	cornerColors := helpers.GetColorsbyLength(colors, 3) 
	
	// Assign 
	for _, t := range types {
		if t == "center" {
			orderedColors = append(orderedColors, centerColors[0])
			slices.Delete(centerColors, 0, 1)
		} else if t == "corner" {
			orderedColors = append(orderedColors, cornerColors[0])
			slices.Delete(cornerColors, 0, 1)
		} else {
			orderedColors = append(orderedColors, edgeColors[0])
			slices.Delete(edgeColors, 0, 1)
		}
	} 

	return orderedColors
}

// function to swap elements in color tensors slice
func SwapElements(tensorSlice models.PieceColors, elementsToSwap models.PieceColors) models.PieceColors {
	// helper func to find element index
	findIndex := func(slice models.PieceColors, element []string) int {
		for i, el := range slice {
			if len(el) == len(element) {
				match := true // Check piece type
				for j := range el {
					if el[j] != element[j] { // Loop to check if is the element we search
						match = false
						break
					}
					
				}
				if match {
					return i // If is the element return the index of it
				}
			}
		}
		return -1 // Or not found
	}

	// Find its indices
	indices := make([]int, len(elementsToSwap))
	for i, el := range elementsToSwap {
		indices[i] = findIndex(tensorSlice, el)
	}


	// Perform the swaps
	if len(indices) == 4 {
		// Swap in circle
		temp := tensorSlice[indices[3]]
		tensorSlice[indices[3]] = tensorSlice[indices[2]]
        tensorSlice[indices[2]] = tensorSlice[indices[1]]
        tensorSlice[indices[1]] = tensorSlice[indices[0]]
        tensorSlice[indices[0]] = temp
	}

	return tensorSlice
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



