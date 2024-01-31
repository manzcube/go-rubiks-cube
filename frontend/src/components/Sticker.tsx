import React from "react";

const getColor = (symbol: string) => {
  switch (symbol) {
    case "White":
      return "bg-white";
    case "Yellow":
      return "bg-yellow";
    case "Green":
      return "bg-green";
    case "Blue":
      return "bg-blue";
    case "Red":
      return "bg-red";
    case "Orange":
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
