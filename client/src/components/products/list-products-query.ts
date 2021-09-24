import { gql } from '@apollo/client';

export const PRODUCTS_QUERY = gql`
query listProducts {
    products {
        id
        title
        description
        price
        image
    }
}
`;
