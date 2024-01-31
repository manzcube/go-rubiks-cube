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
	var combinations, positions models.CubeCombinations
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
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{0, 2, 2},
		{0, 2, 0},
		{2, 2, 0},
		{2, 2, 2},
	}
	var edgesToSwap = models.CubeCombinations {
		{0, 2, 1},
		{1, 2, 0},
		{2, 2, 1},
		{1, 2, 2},
	}
	// Find corner indices
	cornerindices := make([]int, len(cornersToSwap))
	for i, el := range cornersToSwap {
		cornerindices[i] = helpers.FindIndexByTensor(cube, el)
	}
	// Find edges indices
	edgesIndices := make([]int, len(edgesToSwap))
	for i, el := range edgesToSwap {
		edgesIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	if (len(cornerindices) == len(cornersToSwap)) && (len(edgesIndices) == len(edgesToSwap))  {
		// Perform the swap
		newCornersCube := utils.SwapCubeElements(cube, cornerindices)
		newEdgesCube := utils.SwapCubeElements(newCornersCube, edgesIndices)

		// Return turned cube to the client
		turnedCube = newEdgesCube
	}

	// Rotate color slices to for appropiate turn
	rotatedCube := utils.GetPiecesRotated(turnedCube, cornersToSwap, edgesToSwap, true, true)

	return rotatedCube
}

// Turn R prime
func TurnRPrime(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{2, 2, 2},
		{2, 2, 0},
		{0, 2, 0},
		{0, 2, 2},
	}
	var edgesToSwap = models.CubeCombinations {
		{1, 2, 2},
		{2, 2, 1},
		{1, 2, 0},
		{0, 2, 1},
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
		turnedCube = newEdgesCube
	}

	// Rotate color slices to for appropiate turn
	rotatedCube := utils.GetPiecesRotated(turnedCube, cornersToSwap, edgesToSwap, false, true)

	return rotatedCube
}

// Turn L
func TurnL(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{0, 0, 0},
		{0, 0, 2},
		{2, 0, 2},
		{2, 0, 0},
	}
	var edgesToSwap = models.CubeCombinations {
		{0, 0, 1},
		{1, 0, 2},
		{2, 0, 1},
		{1, 0, 0},
	}
	// Find corner indices
	cornerindices := make([]int, len(cornersToSwap))
	for i, el := range cornersToSwap {
		cornerindices[i] = helpers.FindIndexByTensor(cube, el)
	}
	// Find edges indices
	edgesIndices := make([]int, len(edgesToSwap))
	for i, el := range edgesToSwap {
		edgesIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	if (len(cornerindices) == len(cornersToSwap)) && (len(edgesIndices) == len(edgesToSwap))  {
		// Perform the swap
		newCornersCube := utils.SwapCubeElements(cube, cornerindices)
		newEdgesCube := utils.SwapCubeElements(newCornersCube, edgesIndices)

		// Return turned cube to the client
		turnedCube = newEdgesCube
	}

	// Rotate color slices to for appropiate turn
	rotatedCube := utils.GetPiecesRotated(turnedCube, cornersToSwap, edgesToSwap, true, true)

	return rotatedCube
}

// Turn L prime
func TurnLPrime(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { 
		{2, 0, 0},
		{2, 0, 2},
		{0, 0, 2},
		{0, 0, 0},
	}
	var edgesToSwap = models.CubeCombinations {
		{1, 0, 0},
		{2, 0, 1},
		{1, 0, 2},
		{0, 0, 1},
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
		turnedCube = newEdgesCube
	}

	// Rotate color slices to for appropiate turn
	rotatedCube := utils.GetPiecesRotated(turnedCube, cornersToSwap, edgesToSwap, false, true)

	return rotatedCube
}

// Turn U
func TurnU(cube models.Cube) models.Cube {
	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{0, 0, 0},
		{0, 2, 0},
		{0, 2, 2},
		{0, 0, 2},
	}
	var edgesToSwap = models.CubeCombinations {
		{0, 0, 1},
		{0, 1, 0},
		{0, 2, 1},
		{0, 1, 2},
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

// Turn U prime
func TurnUPrime(cube models.Cube) models.Cube {
	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{0, 0, 2},
		{0, 2, 2},
		{0, 2, 0},
		{0, 0, 0},
	}
	var edgesToSwap = models.CubeCombinations {
		{0, 1, 2},
		{0, 2, 1},
		{0, 1, 0},
		{0, 0, 1},		
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

// Turn D
func TurnD(cube models.Cube) models.Cube {
	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{2, 0, 2},
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 0},
	}
	var edgesToSwap = models.CubeCombinations {
		{2, 1, 2},
		{2, 2, 1},
		{2, 1, 0},
		{2, 0, 1},		
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

// Turn D prime
func TurnDPrime(cube models.Cube) models.Cube {
	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{2, 0, 0},
		{2, 2, 0},
		{2, 2, 2},
		{2, 0, 2},
	}
	var edgesToSwap = models.CubeCombinations {
		{2, 0, 1},
		{2, 1, 0},
		{2, 2, 1},
		{2, 1, 2},
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

 
// Turn F
func TurnF(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{0, 0, 2},
		{0, 2, 2},
		{2, 2, 2},
		{2, 0, 2},
	}
	var edgesToSwap = models.CubeCombinations {
		{0, 1, 2},
		{1, 2, 2},
		{2, 1, 2},
		{1, 0, 2},
	}
	// Find corner indices
	cornerindices := make([]int, len(cornersToSwap))
	for i, el := range cornersToSwap {
		cornerindices[i] = helpers.FindIndexByTensor(cube, el)
	}
	// Find edges indices
	edgesIndices := make([]int, len(edgesToSwap))
	for i, el := range edgesToSwap {
		edgesIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	if (len(cornerindices) == len(cornersToSwap)) && (len(edgesIndices) == len(edgesToSwap))  {
		// Perform the swap
		newCornersCube := utils.SwapCubeElements(cube, cornerindices)
		newEdgesCube := utils.SwapCubeElements(newCornersCube, edgesIndices)

		// Return turned cube to the client
		turnedCube = newEdgesCube
	}

	// Rotate color slices to for appropiate turn
	rotatedCube := utils.GetPiecesRotated(turnedCube, cornersToSwap, edgesToSwap, true, false)

	return rotatedCube
}

// Turn F prime
func TurnFPrime(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{2, 0, 2},
		{2, 2, 2},
		{0, 2, 2},
		{0, 0, 2},
	}
	var edgesToSwap = models.CubeCombinations {
		{1, 0, 2},
		{2, 1, 2},
		{1, 2, 2},
		{0, 1, 2},
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
		turnedCube = newEdgesCube
	}

	// Rotate color slices to for appropiate turn
	rotatedCube := utils.GetPiecesRotated(turnedCube, cornersToSwap, edgesToSwap, false, false)

	return rotatedCube
}

// Turn B
func TurnB(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{2, 0, 0},
		{2, 2, 0},
		{0, 2, 0},
		{0, 0, 0},
	}
	var edgesToSwap = models.CubeCombinations {
		{1, 0, 0},
		{2, 1, 0},
		{1, 2, 0},
		{0, 1, 0},
	}
	// Find corner indices
	cornerindices := make([]int, len(cornersToSwap))
	for i, el := range cornersToSwap {
		cornerindices[i] = helpers.FindIndexByTensor(cube, el)
	}
	// Find edges indices
	edgesIndices := make([]int, len(edgesToSwap))
	for i, el := range edgesToSwap {
		edgesIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	if (len(cornerindices) == len(cornersToSwap)) && (len(edgesIndices) == len(edgesToSwap))  {
		// Perform the swap
		newCornersCube := utils.SwapCubeElements(cube, cornerindices)
		newEdgesCube := utils.SwapCubeElements(newCornersCube, edgesIndices)

		// Return turned cube to the client
		turnedCube = newEdgesCube
	}

	// Rotate color slices to for appropiate turn
	rotatedCube := utils.GetPiecesRotated(turnedCube, cornersToSwap, edgesToSwap, true, false)

	return rotatedCube
}

// Turn B prime
func TurnBPrime(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var cornersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{0, 0, 0},
		{0, 2, 0},
		{2, 2, 0},
		{2, 0, 0},
	}
	var edgesToSwap = models.CubeCombinations {
		{0, 1, 0},
		{1, 2, 0},
		{2, 1, 0},
		{1, 0, 0},
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
		turnedCube = newEdgesCube
	}

	// Rotate color slices to for appropiate turn
	rotatedCube := utils.GetPiecesRotated(turnedCube, cornersToSwap, edgesToSwap, false, false)

	return rotatedCube
}

// Turn M
func TurnM(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var centersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{0, 1, 1},
		{1, 1, 0},
		{2, 1, 1},
		{1, 1, 2},
	}
	var edgesToSwap = models.CubeCombinations {
		{0, 1, 0},
		{2, 1, 0},
		{2, 1, 2},
		{0, 1, 2},
	}
	// Find corner indices
	centerIndices := make([]int, len(centersToSwap))
	for i, el := range centersToSwap {
		centerIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	// Find edges indices
	edgesIndices := make([]int, len(edgesToSwap))
	for i, el := range edgesToSwap {
		edgesIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	if (len(centerIndices) == len(centersToSwap)) && (len(edgesIndices) == len(edgesToSwap))  {
		// Perform the swap
		newEdgesCube := utils.SwapCubeElements(cube, edgesIndices)
		newCentersCube := utils.SwapCubeElements(newEdgesCube, centerIndices)

		// Return turned cube to the client
		turnedCube = newCentersCube
	}

	// Edge rotation
	rotatedCube := utils.GetEdgesRotated(turnedCube, edgesToSwap)

	return rotatedCube
}

// Turn M prime
func TurnMPrime(cube models.Cube) models.Cube {
	var turnedCube models.Cube

	// ElementsToSwap has to access color slices by tensor values
	var centersToSwap = models.CubeCombinations { // HEre I should automize a func to take all pieces of that layer the user wants to turn, then swap by Type
		{1, 1, 2},
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	var edgesToSwap = models.CubeCombinations {
		{0, 1, 2},
		{2, 1, 2},
		{2, 1, 0},
		{0, 1, 0},
	}
	// Find corner indices
	centerIndices := make([]int, len(centersToSwap))
	for i, el := range centersToSwap {
		centerIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	// Find edges indices
	edgesIndices := make([]int, len(edgesToSwap))
	for i, el := range edgesToSwap {
		edgesIndices[i] = helpers.FindIndexByTensor(cube, el)
	}
	if (len(centerIndices) == len(centersToSwap)) && (len(edgesIndices) == len(edgesToSwap))  {
		// Perform the swap
		newEdgesCube := utils.SwapCubeElements(cube, edgesIndices)
		newCentersCube := utils.SwapCubeElements(newEdgesCube, centerIndices)

		// Return turned cube to the client
		turnedCube = newCentersCube
	}

	// Edge rotation
	rotatedCube := utils.GetEdgesRotated(turnedCube, edgesToSwap)

	return rotatedCube
}