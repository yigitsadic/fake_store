import React from "react";
import {useCreatePaymentLinkMutation} from "../../../generated/graphql";

const PaymentButton: React.FC = () => {
    const [createLink, { data, loading, error }] = useCreatePaymentLinkMutation();

    if (data && data.startPayment.url) {
        global.window.location.href = data.startPayment.url;
    }

    return <>
        <button className="btn btn-lg btn-success" disabled={loading || !!error} onClick={() => createLink()}>Proceed to payment</button>
    </>;
}

export default PaymentButton;
