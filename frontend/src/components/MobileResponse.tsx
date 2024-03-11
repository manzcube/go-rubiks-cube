import React from "react";

const monitor = require("../images/monitor.png");

const MobileResponse: React.FC = () => {
  return (
    <div className="mobile-response-page">
      Please, visit this website on a bigger Device
      <img src={monitor} alt="monitor-icon" />
    </div>
  );
};

export default MobileResponse;
