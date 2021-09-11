import React from "react";
import {useAppDispatch, useAppSelector} from "../../store/hooks";
import {selectedCurrentUser, updateCartCount} from "../../store/auth/auth";
import {useAddItemToCartMutation} from "../../generated/graphql";

interface ProductDetailProps {
    product: {
        id: string,
        title: string,
        description: string,
        price: number,
        image: string
    };
}

const ProductDetail: React.FC<ProductDetailProps> = ({ product }: ProductDetailProps) => {
    const dispatch = useAppDispatch();
    const currentUser = useAppSelector(selectedCurrentUser);

    const [addToCartFn, {loading, error}] = useAddItemToCartMutation();

    const handleAddToCart = () => {
        addToCartFn({variables: {productId: product.id}})
            .then(result => {
                const itemsCount = result.data?.addToCart.itemsCount;

                if (itemsCount) {
                    dispatch(updateCartCount(itemsCount));
                }
            });
    }

    return <div className="col">
        <div className="card shadow-sm">
            <img src={product.image} alt={product.title} />

            <div className="card-body">
                <p className="card-title">{product.title}</p>

                <p className="card-text">
                    {product.description}
                </p>

                <div className="d-flex justify-content-between align-items-center">
                    <b className="text-muted">
                        {product.price.toFixed(2)} EUR
                    </b>

                    <button type="button"
                            className="btn btn-sm btn-outline-success"
                            disabled={!currentUser.loggedIn}
                            onClick={() => handleAddToCart()}
                    >
                        {error ? "Try again..." : (loading ? "Working..." : "Add to Cart")}
                    </button>
                </div>
            </div>
        </div>
    </div>;
}

export default ProductDetail;
