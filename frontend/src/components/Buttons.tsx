import React from "react";

import { ButtonsProps, DataProps, Piece } from "../constants/interfaces";
import { useError } from "../hooks/useError";

const Buttons: React.FC<ButtonsProps> = ({ data, setData }) => {
  const BASE_URL = "http://localhost:8080";
  const { setErrorMessage } = useError();

  const turn = async (data: Piece[] | null, movement: string) => {
    try {
      const response = await fetch(`${BASE_URL}/turn/${movement}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        console.log("THE RESPONSE", response);
        throw new Error(response.statusText);
      }

      const jsonData = await response.json();
      setData(jsonData);
      return;
    } catch (e) {
      setErrorMessage(`${e}`);
      console.log("Error sending request:", e);
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
    </div>
  );
};

export default Buttons;
