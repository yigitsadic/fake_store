import React from "react";

import  { Link } from "react-router-dom";

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
                    <button type="button" className="btn btn-sm btn-outline-info position-relative">
                        Cart
                        <span className="position-absolute top-0 start-100 translate-middle badge rounded-pill bg-danger">5
                            <span className="visually-hidden">unread messages</span>
                        </span>
                    </button>
                </Link>
            </li>
        </ul>
    )
};

export default Links;