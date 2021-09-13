import React from "react";
import {Maybe} from "../../generated/graphql";
import OrderProduct from "./OrderProduct";

export interface OrderProduct {
    id: string;
    title: string;
    description: string;
    price: number;
    image: string;
}

export interface Order {
    paymentAmount: number;
    createdAt: string;
    orderItems?: Maybe<OrderProduct[]>;
}

const OrderListItem: React.FC<{order: Order}> = ({order}) => {
    return <div>
        <h3>Order Total: {order.paymentAmount.toFixed(2)} EUR</h3>

        <p>
            <small>{order.createdAt}</small>
        </p>

        <div className="row">
            <div className="col-7">
                {order.orderItems?.map(product => <OrderProduct product={product} key={product.id} />)}
            </div>
        </div>
    </div>;
}

export default OrderListItem;
