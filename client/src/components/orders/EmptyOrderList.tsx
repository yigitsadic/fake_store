import React from "react";
import emptyOrderIcon from "./box.png";

const EmptyOrderList: React.FC = () => {
    return <div className="row">
        <div className="col-12 text-center">
            <img src={emptyOrderIcon} alt="empty order list" />

            <br /><br />

            <h3>You have no orders yet.</h3>
        </div>
    </div>;
}

export default EmptyOrderList;
