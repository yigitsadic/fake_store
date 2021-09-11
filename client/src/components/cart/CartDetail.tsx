import React from "react";
import {useCartContentQuery} from "../../generated/graphql";

import CartItem from "./CartItem";
import EmptyCart from "./EmptyCart";

const CartDetail: React.FC = () => {
    const { data, loading, error } = useCartContentQuery();

    if (loading) return <h3>Loading...</h3>;
    if (error) return <h3>Error occurred during fetching current cart</h3>;

    if (data && data.cart.items) {
        if (data.cart.items?.length > 0) {
            return <div className="row">
                <div className="col-2">&nbsp;</div>

                <div className="col-4">
                    {data.cart.items.map(item => <CartItem key={item.id} item={item} /> )}
                </div>

                <div className="col-1">&nbsp;</div>

                <div className="col-3">
                    <button className="btn btn-lg btn-success" disabled={true}>Proceed to payment</button>
                </div>
            </div>
        } else {
            return <EmptyCart />;
        }
    } else {
        return <EmptyCart />;
    }
}

export default CartDetail;
