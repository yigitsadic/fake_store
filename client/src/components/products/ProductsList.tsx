import React from "react";
import ProductDetail from "./ProductDetail";

const ProductsList: React.FC = () => {
    return <div className="album py-5 bg-light">
        <div className="container">
            <div className="row row-cols-1 row-cols-sm-5 row-cols-md-5 g-3">
                <ProductDetail />
                <ProductDetail />
                <ProductDetail />
                <ProductDetail />
                <ProductDetail />
            </div>
        </div>
    </div>;
}

export default ProductsList;
