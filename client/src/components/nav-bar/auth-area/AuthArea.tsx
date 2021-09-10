import React from "react";
import {useAppSelector} from "../../../app/hooks";
import {selectLoggedIn} from "../../../features/auth/auth";
import AuthenticatedUser from "./AuthenticatedUser";
import Unauthenticated from "./Unauthenticated";

const AuthArea: React.FC = () => {
    const loggedIn = useAppSelector(selectLoggedIn);

    return loggedIn ? <AuthenticatedUser /> :  <Unauthenticated />;
}

export default AuthArea;
