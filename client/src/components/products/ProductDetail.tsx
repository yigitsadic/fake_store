import React from "react";
import {Link} from "react-router-dom";

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

                    <Link to={`/products/${product.id}`}>
                        <button className="btn btn-outline-secondary">Details</button>
                    </Link>
                </div>
            </div>
        </div>
    </div>;
}

export default ProductDetail;
