import React from "react";

import  { Link } from "react-router-dom";
import CartLink from "./CartLink";

const Links: React.FC = () => {
    return (
        <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
                <Link to="/products" className="nav-link">
                    <button type="button" className="btn btn-sm btn-close-white position-relative">Products</button>
                </Link>
            </li>

            <li className="nav-item">
                <Link to="/orders" className="nav-link">
                    <button type="button" className="btn btn-sm btn-success position-relative">Orders</button>
                </Link>
            </li>

            <li className="nav-item">
                <Link to="/cart" className="nav-link">
                    <CartLink />
                </Link>
            </li>
        </ul>
    )
};

export default Links;