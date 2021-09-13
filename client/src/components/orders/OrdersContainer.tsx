import React from "react";
import {useAppSelector} from "../../store/hooks";
import {selectedCurrentUser} from "../../store/auth/auth";
import RestrictedArea from "../restricted-area/RestrictedArea";
import OrderList from "./OrderList";

const OrdersContainer: React.FC = () => {
    const { loggedIn } = useAppSelector(selectedCurrentUser);

    if (!loggedIn) return <RestrictedArea />;


    return <div className="row">
        <div className="col-3">&nbsp;</div>
        <div className="col-6"><OrderList /></div>
    </div>
}

export default OrdersContainer;
