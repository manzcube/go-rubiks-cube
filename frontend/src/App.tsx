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
import MobileResponse from "./components/MobileResponse";

// Types
import { Piece } from "./constants/interfaces";

// Images
const go = require("./images/go-lang.jpg");
const cube = require("./images/cube.png");

function App() {
  let isMobile = window.innerWidth / window.innerHeight < 1.6;
  const [data, setData] = useState<Piece[] | null>(null);
  const { setErrorMessage } = useError();

  useEffect(() => {
    setErrorMessage("");
    const getAbstraction = async () => {
      try {
        const response = await fetch(
          `${process.env.REACT_APP_PUBLIC_ENDPOINT}`
        );
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

  return isMobile ? (
    <MobileResponse />
  ) : (
    <div className="App">
      <ErrorBox />
      <div className="column">
        <header className="header">
          Rubik's Cube Go Simulation
          <img src={cube} alt="cube icon" />
          <img src={go} alt="go icon" />
        </header>
        <div>
          <h5>Turn the Cube</h5>
          <Buttons data={data} setData={setData} />
        </div>
        <div>
          <h5>2D Representation</h5>
          <Cube data={data} />
        </div>
      </div>
      <div className="column">
        <RubiksCube cubeData={data} />
      </div>
    </div>
  );
}

export default App;
