import React from "react";

// Components
import Face from "./Face";

// Constants
import { DataProps } from "../constants/interfaces";

const Cube: React.FC<DataProps> = ({ data }) => {
  return data ? (
    <div className="rubiks-cube-abstraction">
      <div>
        <Face face={data} />
      </div>
      {/*<div>
        <Face face={data[1]} />
        <Face face={data[2]} />
        <Face face={data[3]} />
      </div>
      <div>
        <Face face={data[4]} />
      </div>
      <div>
        <Face face={data[5]} />
      </div> */}
    </div>
  ) : (
    <></>
  );
};

export default Cube;
