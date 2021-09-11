import React from "react";
import {useAppSelector} from "../../store/hooks";
import {selectedCurrentUser} from "../../store/auth/auth";
import RestrictedArea from "../restricted-area/RestrictedArea";
import CartDetail from "./CartDetail";

const CartDetailContainer: React.FC = () => {
    const { loggedIn } = useAppSelector(selectedCurrentUser);

    if (loggedIn) {
        return <CartDetail />;
    }

    return <RestrictedArea />;
}

export default CartDetailContainer;
