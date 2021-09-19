import React from "react";
import {Route, useRouteMatch, Link} from "react-router-dom";
import ProductShow from "./ProductShow";

const ProductsRouteWrapper: React.FC = () => {
    const { url } = useRouteMatch();

    return <>
        <Route path={`${url}/:productId`} component={ProductShow} />
    </>;
}

export default ProductsRouteWrapper;
