import React from "react";
import {logout, selectedCurrentUser} from "../../../store/auth/auth";
import {useAppDispatch, useAppSelector} from "../../../store/hooks";

const AuthenticatedUser: React.FC = () => {
    const dispatch = useAppDispatch();
    const currentUser = useAppSelector(selectedCurrentUser);

    return <>
        <button type="button" className="btn btn-primary position-relative">
            <img src={currentUser?.avatar} width="20px" height="20px" /> &nbsp;&nbsp;
            {currentUser?.fullName}
        </button>

        <button type="button"
                className="btn btn-danger position-relative"
                onClick={() => dispatch(logout())}>
            Logout
        </button>
    </>;
}

export default AuthenticatedUser;
