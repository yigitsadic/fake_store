import React from "react";

const Links: React.FC = () => {
    return (
        <ul className="navbar-nav me-auto mb-2 mb-md-0">
            <li className="nav-item">
                <a className="nav-link active" aria-current="page" href="#">Products</a>
            </li>

            <li className="nav-item">
                <a className="nav-link" href="#">Orders</a>
            </li>

            <li className="nav-item">
                <a className="nav-link" href="#">Cart</a>
            </li>
        </ul>
    )
};

export default Links;