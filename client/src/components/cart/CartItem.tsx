import React from "react";
import {useRemoveFromCartMutation} from "../../generated/graphql";
import {CART_CONTENT_QUERY} from "./cart-content-query";
import {CART_ITEM_COUNT_QUERY} from "../nav-bar/cart-link/cart-item-count-query";

interface CartItemProps {
    item: {
        id: string,
        title: string,
        description: string,
        price: number,
        image: string
    },
}

const CartItem:React.FC<CartItemProps> = ({ item }: CartItemProps) => {
    const [removeFromCartFn, { loading, error }] = useRemoveFromCartMutation();

    const handleRemoveFromCart = () => {
        removeFromCartFn({
            variables: {
                productId: item.id,
            },
            refetchQueries: [{ query: CART_CONTENT_QUERY }, {query: CART_ITEM_COUNT_QUERY}],
        });
    }

    return <>
        <div className="card mb-3">
            <div className="row g-0">
                <div className="col-md-4">
                    <img src={item.image} className="img-fluid rounded-start" alt="..." />
                </div>
                <div className="col-md-8">
                    <div className="card-body">
                        <h5 className="card-title">{item.title}</h5>
                        <p className="card-text">{item.description}</p>

                        <div className="d-flex justify-content-between align-items-center">
                            {item.price.toFixed(2)} EUR

                            <div className="btn-group">
                                <button type="button"
                                        className="btn btn-sm btn-outline-danger"
                                        onClick={() => handleRemoveFromCart()}
                                        disabled={loading}>
                                    {error ? "Try again..." : (loading ? "Working..." : "Remove from Cart")}
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </>;
}

export default CartItem;
