import React from "react";
import emptyCartIcon from "./box.png";

const EmptyCart: React.FC = () => {
    return <div className="row">
        <div className="col-12 text-center">
            <img src={emptyCartIcon} alt="empty cart" />

            <br /><br />

            <h3>Your cart is empty.</h3>
        </div>
    </div>;
}

export default EmptyCart;
