import React from "react";
import {User} from "../../user";
import {useHistory} from "react-router-dom";
import {AuthInterface} from "../nav-bar/auth-interface";

const LoginPage: React.FC<AuthInterface> = ({setCurrentUser}) => {
    const history = useHistory();

    const user = {
        fullName: "Yiğit Sadıç",
        id: "12312",
        avatar: "https://avatars.dicebear.com/api/human/c5608998bd4d4e969b1e4b558833199c.svg"
    }

    const handleLogin = () => {
        if (setCurrentUser) {
            setCurrentUser(user);
        }

        history.push("/")
    }

    return <div className="container-fluid">
        <div className="row">
            <div className="col-12 text-center">
                <button className="btn btn-lg btn-success"
                        onClick={() => handleLogin()}>
                    Login
                </button>
            </div>
        </div>
    </div>;
};

export default LoginPage;
