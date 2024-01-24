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

export interface Piece {
  Tensor: number[];
  Colors: string[];
  PieceType: string;
}
