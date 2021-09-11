import React from "react";
import ProductDetail from "./ProductDetail";
import {useListProductsQuery} from "../../generated/graphql";

const ProductsList: React.FC = () => {
    const {data, loading, error} = useListProductsQuery();

    if (loading) return <h3>Loading...</h3>;
    if (error) return <h3>Error occurred during displaying products. Please try again.</h3>

    if (data && data.products) {
        return <div className="album py-5 bg-light">
            <div className="container">
                <div className="row row-cols-1 row-cols-sm-5 row-cols-md-5 g-3">
                    {data.products.map(product => <ProductDetail key={product.id} product={product} /> )}
                </div>
            </div>
        </div>;
    }

    return <h3>Error occurred during displaying products. Please try again.</h3>
}

export default ProductsList;
