import React from "react";
import {Link, useParams} from "react-router-dom";
import {useProductDetailQuery} from "../../generated/graphql";

const ProductShow: React.FC = () => {
    const {productId} = useParams< {productId: string} >();

    const {data, loading, error} = useProductDetailQuery({variables: {id: productId}});

    if (loading) {
        return <h3>Loading...</h3>;
    }

    if (error) {
        return <h3>Error occurred</h3>;
    }

    if (data?.product) {
        const { product } = data;

        return <>
            <div className="row">
                <div className="col-3">

                </div>

                <div className="col-6">
                    <div className="card shadow-sm">
                        <img src={product.image} alt={product.title} width={150} height={150} />

                        <div className="card-body">
                            <p className="card-title">{product.title}</p>

                            <p className="card-text">
                                {product.description}
                            </p>

                            <div className="d-flex justify-content-between align-items-center">
                                <b className="text-muted">
                                    {product.price.toFixed(2)} EUR
                                </b>

                                <div className="btn-group" role="group" aria-label="Basic example">
                                    <button className="btn btn-outline-danger">❤️ Add to favourites</button>
                                    <button className="btn btn-outline-success">🛒 Add to cart</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    }

    return <>
        Loading...
    </>;
}

export default ProductShow;
