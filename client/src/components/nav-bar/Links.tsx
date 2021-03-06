import React from "react";

import  { Link } from "react-router-dom";

const Links: React.FC = () => {
    return (
        <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
                <Link to="/orders" className="nav-link">
                    <button type="button" className="btn btn-sm btn-success position-relative">Orders</button>
                </Link>
            </li>

            <li className="nav-item">
                <Link to="/cart" className="nav-link">
                    <button type="button" className="btn btn-sm btn-outline-info position-relative">Cart</button>
                </Link>
            </li>
        </ul>
    )
};

export default Links;