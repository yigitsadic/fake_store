import React from "react";
import {useCartCurrentItemCountQuery} from "../../../generated/graphql";

const CartLink: React.FC = () => {
    const {data} = useCartCurrentItemCountQuery();

    if (data && data.cart.itemsCount > 0) {
        return <button type="button" className="btn btn-sm btn-outline-info position-relative">
            Cart
            <span className="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">{data.cart.itemsCount}
                <span className="visually-hidden">items in cart</span>
            </span>
        </button>
    } else {
        return <button type="button" className="btn btn-sm btn-outline-info position-relative">Cart</button>
    }
}

export default CartLink;
