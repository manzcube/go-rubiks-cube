package utils

import (
	"log"
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

// Now we need to assign the colors properly
func AssignPieceColors(positions models.CubeCombinatios, types models.PieceTypes) models.PieceColors {
	// Empty colors slice
	var orderedColors models.PieceColors

	// Counters
	var centerCounter, cornerCounter, edgeCounter int = 0, 0, 0

	// Taking default feed
	centerColors := models.CentersColors 
	edgeColors := models.EdgesColors 
	cornerColors := models.CornersColors
	
	// Assigning 
	for _, t := range types {
		if t == "center" {
			orderedColors = append(orderedColors, centerColors[centerCounter])
			centerCounter++
		} else if t == "corner" {
			orderedColors = append(orderedColors, cornerColors[cornerCounter])
			cornerCounter++
		} else {
			orderedColors = append(orderedColors, edgeColors[edgeCounter])
			edgeCounter++
		}
	} 

	// slices.Delete(centerColors, 0, 1) // This func updates the original var in models. Which causes to color problem

	return orderedColors
}

// function to swap elements in color tensors slice
func SwapCubeElements(cube models.Cube, indices []int) models.Cube {

	// Swap in circle
	temp := cube[indices[len(indices) - 1]].Colors

	// Loop over the middle elements
	log.Println("INDICES ", indices)
	for i := len(indices) - 1; i > 0; i-- {
		cube[indices[i]].Colors = cube[indices[i - 1]].Colors				 
	}

	cube[indices[0]].Colors = temp

	// Return new cube
	return cube
}





















// // Generate all possible color combinations for respective pieces
// func GenerateColors() [][]string {
//     colors := []string{"White", "Orange", "Blue", "Green", "Red", "Yellow"}
//     var combinations [][]string

//     var generate func(int, []string)
//     generate = func(start int, currentCombination []string) {
//         if len(currentCombination) > 0 && len(currentCombination) <= 3 {
//             combinations = append(combinations, append([]string(nil), currentCombination...))
//         }
//         if len(currentCombination) == 3 {
//             return // Stop if the current combination is already of length 3
//         }

//         for i := start; i < len(colors); i++ {
//             if helpers.ContainsColor(currentCombination, "White") && colors[i] == "Yellow" || helpers.ContainsColor(currentCombination, "Yellow") && colors[i] == "White" {
//                 continue // Skip if adding this color violates the restriction
//             } else if helpers.ContainsColor(currentCombination, "Green") && colors[i] == "Blue" || helpers.ContainsColor(currentCombination, "Blue") && colors[i] == "Green" {
//                 continue // Skip if adding this color violates the restriction
//             } else if helpers.ContainsColor(currentCombination, "Red") && colors[i] == "Orange" || helpers.ContainsColor(currentCombination, "Orange") && colors[i] == "Red" {
//                 continue // Skip if adding this color violates the restriction
//             }
//             generate(i+1, append(currentCombination, colors[i]))
//         }
//     }

//     generate(0, []string{})
//     return combinations
// }
// func OrderColorSlicesElements(colors models.PieceColors, positions models.CubeCombinatios, types models.PieceTypes) models.PieceColors {
// 	// This function is meant to loop over all color slices and order each element by their value. 
// 	// Top, Right, Left (corners), Top, Bottom (edges), and default for centers
	
// 	var ColorMapping3 = map[string]int{ // This preserves the 2 dimensional visualization order
// 		"White": 0, 
// 		"Yellow": 1,
// 		"Green": 2, 
// 		"Blue": 3, 
// 		"Orange": 4, 
// 		"Red": 5, 
// 	}
//     var result [][]string

//     for _, el := range colors {
// 		// Empty slice
// 		mappingSlice := []string{}
// 		newSlice := []string{}
// 		for idx, v := range el { // First we create a mapping slice based on colormapping order
// 			orderedIndex := ColorMapping3[v]
// 			if len(mappingSlice) > 0 {
// 				if orderedIndex < ColorMapping3[mappingSlice[idx - 1]] {
// 					n := make([]string, len(mappingSlice))
// 					n = append(n, v)
// 					mappingSlice = append(n, mappingSlice...)
// 				} else {
// 					mappingSlice = append(mappingSlice, v)
// 				}
// 			} else {
// 				mappingSlice = append(mappingSlice, v)
// 			}
// 			log.Println("mapping slice should be len 3", mappingSlice)
// 		}
// 		for i, val := range mappingSlice { // Now we reorder with even index elements first
// 			if i % 2 == 0 {
// 				newSlice = append(newSlice, val)
// 			}
// 		}
// 		for i, val := range mappingSlice { // Now we reorder with odd index elements secondly
// 			if i % 2 != 0 {
// 				newSlice = append(newSlice, val)
// 			}
// 		}
// 		log.Println("ANd this is NEWSLICE", newSlice)

// 		result = append(result, newSlice)
// 	}

//     return result
// }



// function to swap elements in color tensors slice
// func SwapElements(tensorSlice models.PieceColors, elementsToSwap models.PieceColors) models.PieceColors {
// 	// helper func to find element index
// 	findIndex := func(slice models.PieceColors, element []string) int {
// 		for i, el := range slice {
// 			if len(el) == len(element) {
// 				match := true // Check piece type
// 				for j := range el {
// 					if el[j] != element[j] { // Loop to check if is the element we search
// 						match = false
// 						break
// 					}
					
// 				}

// 				if match {
// 					return i // If is the element return the index of it
// 				}
// 			}
// 		}
// 		return -1 // Or not found
// 	}

// 	// Find its indices
// 	indices := make([]int, len(elementsToSwap))
// 	for i, el := range elementsToSwap {
// 		indices[i] = findIndex(tensorSlice, el)
// 	}


// 	// Perform the swaps
// 	if len(indices) == len(elementsToSwap) {
// 		// Swap in circle
// 		temp := tensorSlice[indices[len(elementsToSwap) - 1]]

// 		// Loop over the middle elements
// 		for i := len(elementsToSwap) - 1; i > 0; i-- {
// 			tensorSlice[indices[i]] = tensorSlice[indices[i - 1]]				 
// 		}

// 		tensorSlice[indices[0]] = temp
// 	}

// 	return tensorSlice
// }












