import React from "react";

import { Link } from "react-router-dom";

const AuthArea: React.FC = () => {
    return <>
        <Link to="/login" className="btn btn-outline-success">Login</Link>
    </>
}

export default AuthArea;

