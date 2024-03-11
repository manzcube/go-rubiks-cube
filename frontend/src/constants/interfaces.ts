import { strict } from "assert";
import { ReactNode } from "react";

export interface DataProps {
  data: Piece[] | null;
}

export interface ButtonsProps extends DataProps {
  setData: (state: Piece[] | null) => void;
}

export interface ErrorProviderProps {
  children: ReactNode;
}

export interface ICube {
  data: Piece[] | null;
}

export interface IFace {
  data: Piece[];
  positions: number[];
}

export interface Piece {
  Tensor: number[];
  Colors: string[];
  PieceType: string;
}

// Endpoint
export const publicEndpoint: string =
  "https://rubiks-cube-server-wpcp.onrender.com";
export const localEndpoint: string = "http://localhost:8080";
