import { gql } from '@apollo/client';

export const REMOVE_FROM_CART_MUTATION = gql`
    mutation removeFromCart($productId: ID!) {
        removeFromCart(productId: $productId) {
            itemsCount
        }
    }
`;
