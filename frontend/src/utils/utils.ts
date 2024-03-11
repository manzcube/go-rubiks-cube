import { IFace, Piece } from "../constants/interfaces";
import * as THREE from "three";

export const colorMapping: Record<string, number> = {
  White: 0xffffff, // White
  Yellow: 0xffff31, // Yellow
  Red: 0xff4646, // Red
  Orange: 0xffbe45, // Orange
  Blue: 0x4141ff, // Blue
  Green: 0x47e747, // Green
};

// export const getFacePieces: (n: number[], data: IFace) => IFace = (
//   facePieces: number[],
//   data: IFace
// ) => {
//   let newData: IFace = data;

//   newData = data.map(piece => piece)

//   return newData;
// };

export const getColoredFaces = (piece: Piece) => {
  // This array will contain 6 MeshBasicMaterial for 6 single cube faces
  let pieceFaces = [];

  // Centers
  if (piece.PieceType === "center") {
    for (let i = 0; i < 6; i++) {
      pieceFaces[i] = new THREE.MeshBasicMaterial({
        color: colorMapping[piece.Colors[0]],
      });
    }
    return pieceFaces;

    // Corners, which have combination [0, 2, 5] for THREEjs to display the correct color order (top, front, left)
  } else if (piece.PieceType === "corner") {
    const cornerCombination = [0, 2, 5];
    for (let i = 0; i < 3; i++) {
      pieceFaces[cornerCombination[i]] = new THREE.MeshBasicMaterial({
        color: colorMapping[piece.Colors[i]],
      });
    }
    return pieceFaces;

    // Edges, just need [0, 2] representing (top, front) of Threejs cube
  } else {
    for (let i = 0; i < 2; i++) {
      pieceFaces[i * 2] = new THREE.MeshBasicMaterial({
        color: colorMapping[piece.Colors[i]],
      });
    }
    return pieceFaces;
  }
};

export const getRotationFromTensor = (tensor: number[], type: string) => {
  // CORNERS
  if (type === "corner") {
    if (tensor[0] === 0) {
      // With white
      const middleTwo = tensor[1] === 2;
      let rotation =
        tensor[2] === 0
          ? middleTwo
            ? Math.PI / (2 / 3)
            : Math.PI
          : middleTwo
          ? 0
          : Math.PI / 2;
      return { x: rotation, y: Math.PI, z: 0 };
    } else {
      // With yellow
      const middleTwo = tensor[1] === 2;
      let rotation =
        tensor[2] === 0
          ? middleTwo
            ? 0
            : Math.PI / (2 / 3)
          : middleTwo
          ? Math.PI / 2
          : Math.PI;
      return { x: rotation, y: 0, z: 0 };
    }
    // EDGES
  } else {
    if (tensor[0] === 0) {
      // With white
      let rotation =
        tensor[2] === 1
          ? tensor[1] === 0
            ? Math.PI
            : 0
          : tensor[2] === 0
          ? Math.PI / (2 / 3)
          : Math.PI / 2;
      return { x: rotation, y: Math.PI, z: 0 };
    } else if (tensor[0] === 2) {
      // With yellow
      let rotation =
        tensor[2] === 1
          ? tensor[1] === 0
            ? Math.PI
            : 0
          : tensor[2] === 0
          ? Math.PI / (2 / 3)
          : Math.PI / 2;
      return { x: rotation, y: 0, z: 0 };
    } else {
      // 4 middle edges
      return {
        x: tensor[1] === 2 ? 0 : Math.PI,
        y:
          tensor[2] === 2
            ? tensor[1] === 0
              ? Math.PI / (2 / 3)
              : Math.PI / 2
            : tensor[1] === 0
            ? Math.PI / 2
            : Math.PI / (2 / 3),
        z: Math.PI / 2,
      };
    }
  }
};
