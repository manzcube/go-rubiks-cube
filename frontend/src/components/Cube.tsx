import React from "react";

// Components
import Face from "./Face";

// Constants
import { ICube } from "../constants/interfaces";

const whiteFace: number[] = [0, 1, 2, 3, 4, 5, 6, 7, 8];
const orangeFace: number[] = [0, 1, 2, 9, 10, 11, 17, 18, 19];
const greenFace: number[] = [2, 5, 8, 11, 13, 16, 19, 22, 25];
const redFace: number[] = [6, 7, 8, 14, 15, 16, 23, 24, 25];
const yellowFace: number[] = [17, 18, 19, 20, 21, 22, 23, 24, 25];
const blueFace: number[] = [0, 3, 6, 9, 12, 14, 17, 20, 23];
const whitePositions: number[] = [0, 0, 0, 0, 0, 0, 0, 0, 0];
const orangePositions: number[] = [1, 1, 2, 0, 0, 0, 2, 1, 1];
const greenPositions: number[] = [1, 1, 2, 1, 0, 1, 2, 1, 1];
const redPositions: number[] = [2, 1, 1, 0, 0, 0, 1, 1, 2];
const yellowPositions: number[] = [0, 0, 0, 0, 0, 0, 0, 0, 0];
const bluePositions: number[] = [2, 1, 1, 1, 0, 1, 1, 1, 2];

const Cube: React.FC<ICube> = ({ data }) => {
  return data ? (
    <div className="rubiks-cube-abstraction">
      <div>
        <Face
          data={data.filter((_, i) => whiteFace.includes(i))}
          positions={whitePositions}
        />
      </div>
      <div>
        <Face
          data={data.filter((_, i) => orangeFace.includes(i))}
          positions={orangePositions}
        />
        <Face
          data={data.filter((_, i) => greenFace.includes(i))}
          positions={greenPositions}
        />
        <Face
          data={data.filter((_, i) => redFace.includes(i))}
          positions={redPositions}
        />
      </div>
      <div>
        <Face
          data={data.filter((_, i) => yellowFace.includes(i))}
          positions={yellowPositions}
        />
      </div>
      <div>
        <Face
          data={data.filter((_, i) => blueFace.includes(i))}
          positions={bluePositions}
        />
      </div>
    </div>
  ) : (
    <></>
  );
};

export default Cube;
