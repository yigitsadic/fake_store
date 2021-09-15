import React from "react";
import failedIcon from "./payment-terminal.png";

const PaymentFailed: React.FC = () => {
    return <div className="row">
        <div className="col-12 text-center">
            <img src={failedIcon} alt="payment failed" />

            <br /><br />

            <h3>We cannot complete your order.</h3>
        </div>
    </div>;
}

export default PaymentFailed;
