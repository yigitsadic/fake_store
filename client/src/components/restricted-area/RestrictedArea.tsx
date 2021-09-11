import React from "react";
import restrictedImage from "./restricted-image.png";

const RestrictedArea: React.FC = () => {
    return <div className="row">
        <div className="col-12 text-center">
            <img src={restrictedImage} alt="Restricted access" />

            <h2>Please login to access this page.</h2>
        </div>
    </div>;
}

export default RestrictedArea;
