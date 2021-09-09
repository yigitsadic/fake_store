import React from "react";
import Links from "./Links";
import { Link } from "react-router-dom";
import AuthArea from "./AuthArea";

const NavBar: React.FC = () => {
    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
                <Link to="/" className="navbar-brand">Fake Store</Link>

                <button className="navbar-toggler" type="button" data-bs-toggle="collapse"
                        data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false"
                        aria-label="Toggle navigation">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <div className="collapse navbar-collapse" id="navbarCollapse">
                    <Links />

                    <AuthArea />
                </div>
            </div>
        </nav>
    );
};

export default NavBar;
