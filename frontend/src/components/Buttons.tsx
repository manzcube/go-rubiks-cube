import React from "react";

import { ButtonsProps, DataProps, Piece } from "../constants/interfaces";
import { useError } from "../hooks/useError";
import { localEndpoint, publicEndpoint } from "../constants/interfaces";

const Buttons: React.FC<ButtonsProps> = ({ data, setData }) => {
  const { setErrorMessage } = useError();

  const turn = async (data: Piece[] | null, movement: string) => {
    try {
      const response = await fetch(`${localEndpoint}/turn/${movement}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        throw new Error(response.statusText);
      }

      const jsonData = await response.json();
      setData(jsonData);
      return;
    } catch (e) {
      setErrorMessage(`${e}`);
    }
  };

  return (
    <div className="buttons">
      <button onClick={() => turn(data, "r")}>R</button>
      <button onClick={() => turn(data, "r-prime")}>R'</button>
      <button onClick={() => turn(data, "l")}>L</button>
      <button onClick={() => turn(data, "l-prime")}>L'</button>
      <button onClick={() => turn(data, "u")}>U</button>
      <button onClick={() => turn(data, "u-prime")}>U'</button>
      <button onClick={() => turn(data, "d")}>D</button>
      <button onClick={() => turn(data, "d-prime")}>D'</button>
      <button onClick={() => turn(data, "f")}>F</button>
      <button onClick={() => turn(data, "f-prime")}>F'</button>
      <button onClick={() => turn(data, "b")}>B</button>
      <button onClick={() => turn(data, "b-prime")}>B'</button>
      <button onClick={() => turn(data, "m")}>M</button>
      <button onClick={() => turn(data, "m-prime")}>M'</button>
    </div>
  );
};

export default Buttons;
