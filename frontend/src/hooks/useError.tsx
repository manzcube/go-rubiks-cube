import { useContext, createContext, useState } from "react";

// Constants
import { ErrorProviderProps } from "../constants/interfaces";

// Create a context
export const ErrorContext = createContext({
  errorMessage: "",
  setErrorMessage: (message: string) => {},
});

// Custom hook to use error context
export const useError = () => {
  const context = useContext(ErrorContext);
  if (!context) {
    throw new Error("useError must be used within an ErrorProvider");
  }
  return context;
};

// Create a provider component
export const ErrorProvider: React.FC<ErrorProviderProps> = ({ children }) => {
  const [errorMessage, setErrorMessage] = useState("");

  return (
    <ErrorContext.Provider value={{ errorMessage, setErrorMessage }}>
      {children}
    </ErrorContext.Provider>
  );
};
