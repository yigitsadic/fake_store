import React, {useState} from "react";
import NavBar from "./components/nav-bar/NavBar";
import {BrowserRouter, Route, Switch} from "react-router-dom";
import LoginPage from "./components/login/LoginPage";
import {User} from "./user";

const App: React.FC = () => {
    const [currentUser, setCurrentUser] = useState<User>();

    return (
        <BrowserRouter>
            <NavBar currentUser={currentUser} setCurrentUser={setCurrentUser} />

            <div className="container-fluid">
                <Switch>
                    <Route path="/login">
                        <LoginPage setCurrentUser={setCurrentUser} />
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
