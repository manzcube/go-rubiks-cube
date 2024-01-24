package models

// Piece struct
type Piece struct {
	Tensor []int
	Colors []string
	PieceType string
}

// Cube Data Strcuture formed by an slice of Pieces
type Cube []Piece

// Slices for each prop of Piece on its generation
type CubeCombinatios [][]int
type PieceTypes []string
type PieceColors [][]string


// Default color slices for corners and edges
var CornersColors = PieceColors{
    {"White", "Orange", "Blue"},
    {"White", "Green", "Orange"},
    {"White", "Blue", "Red"},
    {"White", "Red", "Green"},
    {"Yellow", "Blue", "Orange"},
    {"Yellow", "Orange", "Green"},
    {"Yellow", "Red", "Blue"},
    {"Yellow", "Green", "Red"},
}

var EdgesColors = PieceColors{
    {"White", "Orange"},
    {"White", "Blue"},
    {"White", "Green"},
    {"White", "Red"},
    {"Orange", "Blue"},
    {"Orange", "Green"},
    {"Red", "Blue"},
    {"Red", "Green"},
    {"Yellow", "Orange"},
    {"Yellow", "Blue"},    
    {"Yellow", "Green"},
    {"Yellow", "Red"},   
}

var CentersColors = PieceColors{
    {"White"},
    {"Orange"},
    {"Blue"},
    {"Green"},    
    {"Red"},
    {"Yellow"},
}

