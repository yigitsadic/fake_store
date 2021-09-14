import React, {useEffect} from "react";
import {Redirect, Route, Switch} from "react-router-dom";
import jwt_decode from "jwt-decode";

import NavBar from "./components/nav-bar/NavBar";
import ProductsList from "./components/products/ProductsList";
import CartDetailContainer from "./components/cart/CartDetailContainer";
import {TokenPayload} from "./store/auth/token-payload";
import {useAppDispatch} from "./store/hooks";
import {login} from "./store/auth/auth";
import OrdersContainer from "./components/orders/OrdersContainer";
import PaymentFailed from "./components/cart/payment-failed/PaymentFailed";
import PaymentSuccessful from "./components/cart/payment-successful/PaymentSuccessful";


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

    const urlSearchParams = new URLSearchParams(window.location.search);
    const payment_status = urlSearchParams.get("payment_status")

    if (payment_status === "successful") {
        return <Redirect to="/cart/payment-successful" />
    } else if (payment_status === "payment_failed") {
        return <Redirect to="/cart/payment-failed" />
    }

    return (
        <>
            <NavBar />

            <div className="container-fluid">
                <Switch>
                    <Route path="/products">
                        <ProductsList />
                    </Route>

                    <Route path="/orders">
                        <OrdersContainer />
                    </Route>

                    <Route path="/cart/payment-failed">
                        <PaymentFailed />
                    </Route>

                    <Route path="/cart/payment-successful">
                        <PaymentSuccessful />
                    </Route>

                    <Route path="/cart">
                        <CartDetailContainer />
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
