import React from "react";
import {selectCartCount} from "../../store/auth/auth";
import {useAppSelector} from "../../store/hooks";

const CartLink: React.FC = () => {
    const cartCount = useAppSelector(selectCartCount);

    if (cartCount > 0) {
        return <button type="button" className="btn btn-sm btn-outline-info position-relative">
            Cart
            <span className="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">{cartCount}
                <span className="visually-hidden">items in cart</span>
            </span>
        </button>
    } else {
        return <button type="button" className="btn btn-sm btn-outline-info position-relative">Cart</button>
    }
}

export default CartLink;
