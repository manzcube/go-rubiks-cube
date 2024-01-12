package models

// Cube being an array of 6 3x3 matrices
type EmptyCube [6][3][3]int
type Cube [6][3][3]string

type Piece struct {
	Tensor []int
	Colors []string
	PieceType string
}

type CubeTensor []Piece

type CubeCombinatios [][]int
type PieceTypes []string
type PieceColors [][]string
