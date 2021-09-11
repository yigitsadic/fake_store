import { gql } from '@apollo/client';

export const CART_QUERY = gql`
query cartContent {
    cart {
        items {
            id
            title
            description
            price
            image
        }
    }
}
`;

