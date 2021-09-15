import React from "react";
import successIcon from "./confirm.png";

const PaymentSuccessful: React.FC = () => {
    return <div className="row">
        <div className="col-12 text-center">
            <img src={successIcon} alt="payment succeeded" />

            <br /><br />

            <h2 className="text-success">We start to prepare your order.</h2>

            <h3>
                <small>You'll be notified when shipment is en route.</small>
            </h3>
        </div>
    </div>;
}

export default PaymentSuccessful;
