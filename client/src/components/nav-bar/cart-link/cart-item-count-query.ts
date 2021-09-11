import { gql } from '@apollo/client';

export const CART_ITEM_COUNT_QUERY = gql`
    query cartCurrentItemCount {
        cart {
            itemsCount
        }
    }
`;
