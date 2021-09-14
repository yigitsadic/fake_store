import { gql } from '@apollo/client';

export const CREATE_PAYMENT_URL_MUTATION = gql`
    mutation createPaymentLink {
        startPayment {
            url
        }
    }
`;
