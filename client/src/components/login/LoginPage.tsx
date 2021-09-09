import React from "react";
import {Redirect} from "react-router-dom";
import {AuthInterface} from "../nav-bar/auth-interface";
import {useLoginMutation} from "../../generated/graphql";

const LoginPage: React.FC<AuthInterface> = ({setCurrentUser}) => {
    const [loginUser, {data,loading, error}] = useLoginMutation();

    if (loading) {
        return <div>Loading...</div>;
    }

    if (error) {
        return <div>Unable to send request due to {error.message}</div>;
    }

    if (data && setCurrentUser) {
        setCurrentUser({
            id: data.login.id,
            fullName: data.login.fullName,
            avatar: data.login.avatar,
        })

        return <Redirect to="/" />
    }

    const handleLogin = () => {
        loginUser();
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
