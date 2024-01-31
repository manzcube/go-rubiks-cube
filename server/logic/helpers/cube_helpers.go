package helpers

import (
	"reflect"
	"server/logic/models"
	"slices"
)

// contains checks if a slice contains a specific string.
func ContainsColor(slice []string, str string) bool {
    for _, v := range slice {
        if v == str {
            return true
        }
    }
    return false
}

// helper func to find element index
func FindIndex(wholeCube models.Cube, element []string) int {
	for i, piece := range wholeCube {
		if len(piece.Colors) == len(element) {
			match := true // Check piece type
			for j := range element {
				if !slices.Contains(piece.Colors, element[j]) { // Loop to check if slice contains all elements NO MATTER THE ORDER
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

// helper func to find element index by tensor
func FindIndexByTensor(cube models.Cube, tensor []int) int {
	for i, piece := range cube {
		if len(piece.Tensor) == len(tensor) {
			match := true // Check piece type
			for j := range piece.Tensor {
				if piece.Tensor[j] != tensor[j] { // Loop to check if is the element we search, In tensor The ORDER MATTERS
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

// Remove Nucleus piece to avoid trauma
func RemoveNucleus(combinations models.CubeCombinations) models.CubeCombinations {
	var newSlice [][]int
	nucleus := []int{1, 1, 1}
	for _, v := range combinations {
		if !reflect.DeepEqual(v, nucleus) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
} 


// Slices contains
func Contains(slices [][]int, element []int) bool {
	for _, slice := range slices {
		if len(slice) == len(element) {
			match := true
			for i, value := range slice {
				if value != element[i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}

func Index(slices [][]int, element []int) int {
	for i, slice := range slices {
		if len(slice) == len(element) {
			match := true
			for i, value := range slice {
				if value != element[i] {
					match = false
					break
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}

func GetDirectionBool(index int, direction bool) bool {
	if direction {
		return index % 2 == 0
	} else {
		return index % 2 != 0
	}
}


func IsCenter(tensor []int) bool {
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

func Sum(arr []int) int {
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

// Get subslices based on piece types
func GetColorsbyLength(colors models.PieceColors, length int) models.PieceColors {
	var newSlice models.PieceColors
	
	for _, color := range colors {
		if len(color) == length {
			newSlice = append(newSlice, color)
		}
	}

	return newSlice
}