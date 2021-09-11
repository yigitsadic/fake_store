import React from "react";
import {useAppSelector} from "../../store/hooks";
import {selectedCurrentUser} from "../../store/auth/auth";

interface ProductProps {
    id: string,
    title: string,
    description: string,
    price: number,
    image: string
}
interface ProductDetailProps {
    product: ProductProps;
}

const ProductDetail: React.FC<ProductDetailProps> = ({ product }: ProductDetailProps) => {
    const currentUser = useAppSelector(selectedCurrentUser);

    return <div className="col">
        <div className="card shadow-sm">
            <img src={product.image} alt={product.title} />

            <div className="card-body">
                <p className="card-title">{product.title}</p>

                <p className="card-text">
                    {product.description}
                </p>
                <div className="d-flex justify-content-between align-items-center">
                    <div className="btn-group">
                        <button type="button"
                                className="btn btn-sm btn-outline-success"
                                disabled={!currentUser.loggedIn}>
                            Add to Cart
                        </button>
                    </div>

                    <small className="text-muted">{product.price} EUR</small>
                </div>
            </div>
        </div>
    </div>;
}

export default ProductDetail;
