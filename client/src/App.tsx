import React, {useEffect} from "react";
import {Route, Switch} from "react-router-dom";
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
import ProductsRouteWrapper from "./components/products/ProductsRouteWrapper";


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
                        <ProductsRouteWrapper />
                    </Route>

                    <Route path="/orders">
                        <OrdersContainer />
                    </Route>

                    <Route path="/payment_failed">
                        <PaymentFailed />
                    </Route>

                    <Route path="/payment_successful">
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
