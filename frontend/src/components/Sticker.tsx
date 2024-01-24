import React from "react";
const getColor = (symbol: string) => {
  switch (symbol) {
    case "W":
      return "bg-white";
    case "Y":
      return "bg-yellow";
    case "G":
      return "bg-green";
    case "B":
      return "bg-blue";
    case "R":
      return "bg-red";
    case "O":
      return "bg-orange";
  }
};

interface Sticker {
  color: string;
}

const Sticker: React.FC<Sticker> = ({ color }) => {
  return <div className={`sticker ${getColor(color)}`}></div>;
};

export default Sticker;
