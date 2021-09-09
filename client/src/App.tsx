import React from "react";
import NavBar from "./components/nav-bar/NavBar";
import {BrowserRouter, Route, Switch} from "react-router-dom";
import LoginPage from "./components/login/LoginPage";

const App: React.FC = () => {
    return (
        <BrowserRouter>
            <NavBar />

            <div className="container-fluid">
                <Switch>
                    <Route path="/login">
                        <LoginPage />
                    </Route>

                    <Route path="/products">
                        Products
                    </Route>

                    <Route path="/orders">
                        Orders
                    </Route>

                    <Route path="/cart">
                        Cart
                    </Route>

                    <Route path="/">
                        Home
                    </Route>
                </Switch>
            </div>
        </BrowserRouter>
    );
};

export default App;
