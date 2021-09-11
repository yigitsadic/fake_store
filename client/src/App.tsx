import React, {useEffect, useState} from "react";
import {Route, Switch} from "react-router-dom";
import jwt_decode from "jwt-decode";

import NavBar from "./components/nav-bar/NavBar";
import ProductsList from "./components/products/ProductsList";
import CartDetail from "./components/cart/CartDetail";
import OrderList from "./components/orders/OrderList";
import {TokenPayload} from "./store/auth/token-payload";
import {useAppDispatch} from "./store/hooks";
import {login} from "./store/auth/auth";


const App: React.FC = () => {
    const dispatch = useAppDispatch();
    let token = localStorage.getItem("fake_store_token");

    useEffect(() => {
        if (token) {
            const payload = jwt_decode(token) as TokenPayload;

            dispatch(login({
                id: payload.jti,
                fullName: payload.fullName,
                avatar: payload.avatar,
                loggedIn: true,
            }));
        }
    }, [token]);

    return (
        <>
            <NavBar />

            <div className="container-fluid">
                <Switch>
                    <Route path="/products">
                        <ProductsList />
                    </Route>

                    <Route path="/orders">
                        <OrderList />
                    </Route>

                    <Route path="/cart">
                        <CartDetail />
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
