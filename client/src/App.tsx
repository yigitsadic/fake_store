import React from "react";
import NavBar from "./components/nav-bar/NavBar";

const App: React.FC = () => {
    return (
        <>
            <NavBar />

            <div className="container-fluid">
                Hello World
            </div>
        </>
    );
};

export default App;