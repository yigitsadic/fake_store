import React from "react";

import  { Link } from "react-router-dom";

const Links: React.FC = () => {
    return (
        <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
                <Link to="/products" className="nav-link">Products</Link>
            </li>

            <li className="nav-item">
                <Link to="/orders" className="nav-link">Orders</Link>
            </li>

            <li className="nav-item">
                <Link to="/cart" className="nav-link">Cart</Link>
            </li>
        </ul>
    )
};

export default Links;