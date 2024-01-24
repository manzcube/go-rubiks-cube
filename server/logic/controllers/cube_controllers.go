package controllers

import (
	"server/logic/helpers"
	"server/logic/models"
	"server/logic/utils"
)

// // GET controllers

// Render default Cube Data Structure
func RenderCubeTensor() models.Cube {
	// Render all combinatios pieces positions
	var combinations, positions models.CubeCombinatios
	combinations = utils.GenerateCombinations(3) 
	positions = helpers.RemoveNucleus(combinations)

	// Render piece types 
	var pieceTypes models.PieceTypes 
	pieceTypes = utils.GeneratePieceTypes(positions)

	// Assign all color combinations to each piece
	orderedColors := utils.AssignPieceColors(positions, pieceTypes)

	// Create Cube Data Sctructure
	var cube models.Cube
	for i, position := range positions {
		var piece models.Piece
		piece.Tensor = position
		piece.PieceType = pieceTypes[i]
		piece.Colors = orderedColors[i]		

		cube = append(cube, piece)
	}
	
	return cube // Return default cube state
}



// // POST controllers

// Turn R
func TurnR(cube models.Cube) models.Cube {
	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinatios { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{0, 2, 2},
		{0, 2, 0},
		{2, 2, 0},
		{2, 2, 2},
	}
	var edgesToSwap = models.CubeCombinatios {
		{0, 2, 1},
		{1, 2, 0},
		{2, 2, 1},
		{1, 2, 2},
	}
	// Find its corner indices
	cornerindices := make([]int, len(cornersToSwap))
	for i, el := range cornersToSwap {
		cornerindices[i] = helpers.FindIndexByTensor(cube, el)
	}
	// Find its edges indices
	edgesIndices := make([]int, len(edgesToSwap))
	for i, el := range edgesToSwap {
		edgesIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	if (len(cornerindices) == len(cornersToSwap)) && (len(edgesIndices) == len(edgesToSwap))  {
		// Perform the swap
		newCornersCube := utils.SwapCubeElements(cube, cornerindices)
		newEdgesCube := utils.SwapCubeElements(newCornersCube, edgesIndices)

		// Return turned  cube to the client
		return newEdgesCube
	}

	return cube
}
 
// // Turn R Prime
// func TurnRPrime(cube models.Cube) models.Cube {
// 	// We create a new Cube object
// 	newCube := cube

// 	// Set colors
// 	for face := 0; face < 6; face++ {
// 		if face != 1 && face != 3 {
// 			for row := 0; row < 3; row++ {
// 				newCube[face][row][2] = cube[utils.GetFaceForVerticalTurnCounterClockwise(face)][row][2]
// 			}
// 		}
// 	}
// 	return newCube
// }

// // Turn L
// func TurnL(cube models.Cube) models.Cube {
// 	// We create a new Cube object instance
// 	newCube := cube

// 	// Perfom the turn
// 	for face := 0; face < 6; face++ {
// 		if face != 1 && face != 3 {
// 			for row := 0; row < 3; row++ {
// 				newCube[face][row][0] = cube[utils.GetFaceForVerticalTurnCounterClockwise(face)][row][0]
// 			}
// 		}
// 	}
// 	return newCube
// }

// // Turn L Prime
// func TurnLPrime(cube models.Cube) models.Cube {
// 	// We create a new Cube object
// 	newCube := cube

// 	// Set colors
// 	for face := 0; face < 6; face++ {
// 		if face != 1 && face != 3 {
// 			for row := 0; row < 3; row++ {
// 				newCube[face][row][0] = cube[utils.GetFaceForVerticalTurnClockwise(face)][row][0]
// 			}
// 		}
// 	}
// 	return newCube
// }

// // Turn U
// func TurnU(cube models.Cube) models.Cube {
// 	// We create a new Cube object instance
// 	newCube := cube

// 	// Perfom the turn
// 	for face := 0; face < 6; face++ {
// 		if face != 0 && face != 4 {
// 			for row := 0; row < 3; row++ {
// 				newCube[face][0][row] = cube[utils.GetFaceForHorizontalTurnCounterClockwise(face)][0][row]
// 			}
// 		}
// 	}
// 	return newCube
// }

// // Turn U Prime
// func TurnUPrime(cube models.Cube) models.Cube {
// 	// We create a new Cube object
// 	newCube := cube

// 	// Set colors
// 	for face := 0; face < 6; face++ {
// 		if face != 0 && face != 4 {
// 			for row := 0; row < 3; row++ {
// 				newCube[face][0][row] = cube[utils.GetFaceForHorizontalTurnClockwise(face)][0][row]
// 			}
// 		}
// 	}
// 	return newCube
// }
