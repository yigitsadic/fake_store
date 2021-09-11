import React from "react";

const ProductDetail: React.FC = () => {
    return <div className="col">
        <div className="card shadow-sm">
            <img src="https://via.placeholder.com/150" />

            <div className="card-body">
                <p className="card-text">
                    Product description.
                </p>
                <div className="d-flex justify-content-between align-items-center">
                    <div className="btn-group">
                        <button type="button" className="btn btn-sm btn-outline-success">Add to Cart</button>
                    </div>
                </div>
            </div>
        </div>
    </div>;
}

export default ProductDetail;
