import { gql } from '@apollo/client';

export const ADD_TO_CART_MUTATION = gql`
mutation addItemToCart($productId: ID!) {
    addToCart(productId: $productId)
}
`;
