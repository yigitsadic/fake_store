import { gql } from '@apollo/client';

export const REMOVE_FROM_CART_MUTATION = gql`
    mutation removeFromCart($cartItemId: ID!) {
        removeFromCart(cartItemId: $cartItemId)
    }
`;
