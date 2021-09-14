import React from "react";
import {useCreatePaymentLinkMutation} from "../../../generated/graphql";

const PaymentButton: React.FC = () => {
    const [createLink, { data, loading, error }] = useCreatePaymentLinkMutation();

    if (data && data.startPayment.url) {
        global.window.location.href = data.startPayment.url;
    }

    return <>
        <button className="btn btn-lg btn-success" disabled={loading} onClick={() => createLink()}>Proceed to payment</button>

        { error && <p className="lead text-danger">Error occurred during payment. Please try again.</p> }
    </>;
}

export default PaymentButton;
