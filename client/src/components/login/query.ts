import { gql } from '@apollo/client';

export const LOGIN_MUTATION = gql`
    mutation login {
        login {
            id
            avatar
            fullName
            token
        }
    }
`;
