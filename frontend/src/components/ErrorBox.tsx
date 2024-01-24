import React, { useState, useEffect, useContext } from "react";
import { ErrorContext } from "../hooks/useError";

export const ErrorBox: React.FC = () => {
  const { errorMessage, setErrorMessage } = useContext(ErrorContext);
  const [showError, setShowError] = useState(false);

  useEffect(() => {
    if (errorMessage) {
      setShowError(true);
      const timer = setTimeout(() => {
        setShowError(false);
        setErrorMessage(""); // Reset the error message in the context after hiding
      }, 3000);

      return () => clearTimeout(timer); // Clear timeout if the component unmounts
    }
  }, [errorMessage, setErrorMessage]);

  return showError ? (
    <div id="error-message-component">{errorMessage}</div>
  ) : null;
};

export default ErrorBox;
