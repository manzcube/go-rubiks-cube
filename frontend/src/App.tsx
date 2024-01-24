// Lib and Hooks
import React, { useEffect, useState } from "react";
import { useError } from "./hooks/useError";

// Style
import "./App.css";

// Components
import Cube from "./components/Cube";
import Buttons from "./components/Buttons";
import ErrorBox from "./components/ErrorBox";
import RubiksCube from "./components/RubiksCube";
import { Piece } from "./constants/interfaces";

function App() {
  const [data, setData] = useState<Piece[] | null>(null);
  const { setErrorMessage } = useError();

  useEffect(() => {
    setErrorMessage("");
    const getAbstraction = async () => {
      try {
        const response = await fetch("http://localhost:8080/");
        if (!response.ok) {
          throw new Error("Network response was not ok..");
        }

        const jsonData = await response.json();
        setData(jsonData);
      } catch (e) {
        setErrorMessage(`${e}`);
      }
    };

    getAbstraction();
  }, []);

  return (
    <div className="App">
      <ErrorBox />
      <div id="left-box">
        <header className="App-header">Rubik's Cube Simulation</header>
        <Buttons data={data} setData={setData} />
        {/* <Cube data={data} /> */}
      </div>
      <RubiksCube cubeData={data} />
    </div>
  );
}

export default App;
