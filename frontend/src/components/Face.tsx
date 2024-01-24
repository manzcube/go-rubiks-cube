import React from "react";
import Sticker from "./Sticker";

interface Face {
  face: string[];
}

const Face: React.FC<Face> = ({ face }) => {
  if (face.length) {
    return (
      <div className="face">
        {face.map((layer, index) => (
          <div className="face-layer" key={index}>
            <Sticker color={layer[0]} />
            <Sticker color={layer[1]} />
            <Sticker color={layer[2]} />
          </div>
        ))}
      </div>
    );
  } else {
    return <div></div>;
  }
};

export default Face;
