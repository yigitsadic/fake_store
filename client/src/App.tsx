import React from "react";
import NavBar from "./components/nav-bar/NavBar";
import {BrowserRouter, Route, Switch} from "react-router-dom";

const App: React.FC = () => {
    return (
        <BrowserRouter>
            <NavBar />

            <div className="container-fluid">
                <Switch>
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
