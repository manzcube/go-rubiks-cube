import React from "react";
import Sticker from "./Sticker";

import { IFace } from "../constants/interfaces";

const Face: React.FC<IFace> = ({ data, positions }) => {
  if (data.length) {
    return (
      <div className="face">
        {data.map((piece, i) => (
          <Sticker key={i} color={piece.Colors[positions[i]]} />
        ))}
      </div>
    );
  } else {
    return <div></div>;
  }
};

export default Face;
