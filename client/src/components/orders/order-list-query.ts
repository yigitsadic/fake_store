import { gql } from '@apollo/client';

export const ORDER_LIST_QUERY = gql`
    query listOrders {
        orders {
            paymentAmount
            createdAt
            orderItems {
                id
                title
                description
                price
                image
            }
        }
    }
`;
