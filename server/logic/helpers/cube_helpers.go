package helpers

import (
	"reflect"
	"server/logic/models"
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