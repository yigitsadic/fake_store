import React from "react";
import AuthenticatedUser from "./AuthenticatedUser";
import Unauthenticated from "./Unauthenticated";
import {useAppSelector} from "../../../store/hooks";
import {selectedCurrentUser} from "../../../store/auth/auth";

const AuthArea: React.FC = () => {
    const { loggedIn } = useAppSelector(selectedCurrentUser);

    return loggedIn ? <AuthenticatedUser /> :  <Unauthenticated />;
}

export default AuthArea;
