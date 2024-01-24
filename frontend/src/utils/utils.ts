import { Piece } from "../constants/interfaces";
import * as THREE from "three";

export const colorMapping: Record<string, number> = {
  White: 0xffffff, // White
  Yellow: 0xffff31, // Yellow
  Red: 0xff4646, // Red
  Orange: 0xffbe45, // Orange
  Blue: 0x4141ff, // Blue
  Green: 0x47e747, // Green
};

const getPieceOrderedColorFaces = (piece: Piece) => {
  // This array will contain 6 MeshBasicMaterial for 6 single cube faces
  let pieceFaces = [];

  // Fill with default color
  for (let i = 0; i < 6; i++) {
    pieceFaces.push(new THREE.MeshBasicMaterial({ color: 0x000000 }));
  }

  // HERE IS WHERE THE MAGIC OCCURS
  if (piece.PieceType === "center") {
    // For centers
    for (let i = 0; i < 6; i++) {
      pieceFaces[i] = new THREE.MeshBasicMaterial({
        color: colorMapping[piece.Colors[0]],
      });
    }
    return pieceFaces;
  } else if (piece.PieceType === "corner") {
    // For corners
    let newTensor;
    if (piece.Tensor[0] === 0) {
      newTensor =
        piece.Tensor[1] + piece.Tensor[2] === 0
          ? [1, 3, 5]
          : piece.Tensor[1] + piece.Tensor[2] === 4
          ? [1, 2, 4]
          : piece.Tensor[1] === 2
          ? [1, 5, 2]
          : [1, 4, 3];
    } else {
      newTensor =
        piece.Tensor[1] + piece.Tensor[2] === 0
          ? [0, 5, 3]
          : piece.Tensor[1] + piece.Tensor[2] === 4
          ? [0, 4, 2]
          : piece.Tensor[1] === 2
          ? [0, 2, 5]
          : [0, 3, 4];
    }

    for (let i = 0; i < 3; i++) {
      pieceFaces[newTensor[i]] = new THREE.MeshBasicMaterial({
        color: colorMapping[piece.Colors[i]],
      });
    }
    return pieceFaces;
  } else {
    // For edges
    let newTensor;
    if (piece.Tensor[0] === 0) {
      newTensor =
        piece.Tensor[0] + piece.Tensor[1] === 0
          ? [1, 3]
          : piece.Tensor[0] + piece.Tensor[1] === 2
          ? [1, 2]
          : piece.Tensor[2] === 2
          ? [1, 4]
          : [1, 5];
    } else if (piece.Tensor[0] === 1) {
      newTensor =
        piece.Tensor[1] + piece.Tensor[2] === 0
          ? [3, 5]
          : piece.Tensor[1] + piece.Tensor[2] === 4
          ? [2, 4]
          : piece.Tensor[1] === 2
          ? [2, 5]
          : [3, 4];
    } else {
      newTensor =
        piece.Tensor[0] + piece.Tensor[1] === 2
          ? [0, 3]
          : piece.Tensor[0] + piece.Tensor[1] === 4
          ? [0, 2]
          : piece.Tensor[2] === 2
          ? [0, 4]
          : [0, 5];
    }

    for (let i = 0; i < 2; i++) {
      pieceFaces[newTensor[i]] = new THREE.MeshBasicMaterial({
        color: colorMapping[piece.Colors[i]],
      });
    }
    return pieceFaces;
  }
};

export default getPieceOrderedColorFaces;
