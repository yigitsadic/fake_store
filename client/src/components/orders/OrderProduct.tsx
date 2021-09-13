import React from "react";
import {OrderProduct} from "./OrderListItem";

const OrderProduct: React.FC<{product: OrderProduct}> = ({product}) => {
    return <div className="card mb-3">
        <div className="row g-0">
            <div className="col-md-4">
                <img src={product.image} className="img-fluid rounded-start" alt={product.title} />
            </div>
            <div className="col-md-8">
                <div className="card-body">
                    <h5 className="card-title">{product.title}</h5>
                    <p className="card-text">{product.description}</p>

                    <div className="d-flex justify-content-between align-items-center">
                        {product.price.toFixed(2)} EUR
                    </div>
                </div>
            </div>
        </div>
    </div>;
};

export default OrderProduct;
