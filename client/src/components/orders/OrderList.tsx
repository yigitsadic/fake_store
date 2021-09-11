import React from "react";
import {useAppSelector} from "../../store/hooks";
import {selectedCurrentUser} from "../../store/auth/auth";
import RestrictedArea from "../restricted-area/RestrictedArea";

const OrderList: React.FC = () => {
    const { loggedIn } = useAppSelector(selectedCurrentUser);

    if (loggedIn) {
        return <h3>Order List Page</h3>;
    }

    return <RestrictedArea />;
}

export default OrderList;
