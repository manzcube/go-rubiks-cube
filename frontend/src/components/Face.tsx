import React from "react";
import Sticker from "./Sticker";
import { Piece } from "../constants/interfaces";

interface Face {
  face: Piece[];
}

const Face: React.FC<Face> = ({ face }) => {
  if (face.length) {
    return (
      <div className="face">
        {/* Map all pieces for the ones having White or equivalent 1st tensor el = 0 */}
        {face.map((piece) => {
          if (piece.Tensor[0] === 0) {
            return <Sticker color={piece.Colors[0]} />;
          }
        })}
      </div>
    );
  } else {
    return <div></div>;
  }
};

export default Face;
