import React from "react";
import {useAppSelector} from "../../store/hooks";
import {selectedCurrentUser} from "../../store/auth/auth";
import RestrictedArea from "../restricted-area/RestrictedArea";

const CartDetail: React.FC = () => {
    const { loggedIn } = useAppSelector(selectedCurrentUser);

    if (loggedIn) {
        return <h3>Cart Detail Page</h3>;
    }

    return <RestrictedArea />;
}

export default CartDetail;
