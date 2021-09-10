import React from "react";
import {useLoginMutation} from "../../../generated/graphql";
import {login} from "../../../features/auth/auth";
import {useAppDispatch} from "../../../app/hooks";

const Unauthenticated: React.FC = () => {
    const dispatch = useAppDispatch();
    const [loginUser, {data,loading, error}] = useLoginMutation();

    if (data) {
        dispatch(login({
            id: data.login.id,
            fullName: data.login.fullName,
            avatar: data.login.avatar,
        }));
    }

    return <>
        <button
            type="button"
            className="btn btn-outline-success"
            onClick={() => loginUser()}
            disabled={loading}
        >
            {error ? "Error occurred - Retry" : (loading ? "Loading..." : "Login")}
        </button>
    </>;
};

export default Unauthenticated;
