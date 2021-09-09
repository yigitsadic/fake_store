import React, {useEffect} from "react";

import { Link } from "react-router-dom";
import {User} from "../../user";
import {AuthInterface} from "./auth-interface";

const AuthArea: React.FC<AuthInterface> = ({currentUser, setCurrentUser}) => {
    const handleLogout = () => {
        if (setCurrentUser) {
            setCurrentUser({});
        }
    }

    if (currentUser?.id) {
        return <>
            <button type="button" className="btn btn-primary position-relative">
                <img src={currentUser.avatar} width="20px" height="20px" /> &nbsp;&nbsp;
                {currentUser.fullName}
            </button>

            <button type="button"
                    className="btn btn-danger position-relative"
                    onClick={() => handleLogout()}>
                Logout
            </button>
        </>
    } else {
        return <>
            <Link to="/login" className="btn btn-outline-success">Login</Link>
        </>
    }
}

export default AuthArea;

