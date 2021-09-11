import React from "react";
import NavBar from "./components/nav-bar/NavBar";
import {Route, Switch} from "react-router-dom";
import ProductsList from "./components/products/ProductsList";

const App: React.FC = () => {
    return (
        <>
            <NavBar />

            <div className="container-fluid">
                <Switch>
                    <Route path="/products">
                        <ProductsList />
                    </Route>

                    <Route path="/orders">
                        Orders
                    </Route>

                    <Route path="/cart">
                        Cart
                    </Route>

                    <Route path="/">
                        <ProductsList />
                    </Route>
                </Switch>
            </div>
        </>
    );
};

export default App;
