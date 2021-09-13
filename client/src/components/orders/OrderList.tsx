import React from "react";
import {useListOrdersQuery} from "../../generated/graphql";
import OrderListItem from "./OrderListItem";
import {nanoid} from "nanoid";

const OrderList: React.FC = () => {
    const {data, loading, error} = useListOrdersQuery();

    if (loading) return <h3>Loading...</h3>;
    if (error) return <h3>Error occurred during listing orders...</h3>;

    if (data && data.orders) {
        return <div>
            {data.orders.map(order => <OrderListItem key={nanoid()} order={order} />)}
        </div>
    }

    return <h3>Loading...</h3>;
}

export default OrderList;
