import { gql } from '@apollo/client';

export const PRODUCT_DETAIL_QUERY = gql`
    query productDetail($id: ID!) {
        product(ID: $id) {
            id
            title
            description
            price
            image
        }
    }
`;
