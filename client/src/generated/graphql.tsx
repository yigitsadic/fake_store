import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions =  {}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Cart = {
  __typename?: 'Cart';
  items?: Maybe<Array<CartItem>>;
  itemsCount: Scalars['Int'];
};

export type CartItem = {
  __typename?: 'CartItem';
  description: Scalars['String'];
  id: Scalars['ID'];
  image: Scalars['String'];
  price: Scalars['Float'];
  productId: Scalars['ID'];
  title: Scalars['String'];
};

export type LoginResponse = {
  __typename?: 'LoginResponse';
  avatar: Scalars['String'];
  fullName: Scalars['String'];
  id: Scalars['ID'];
  token: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  addToCart: Cart;
  login: LoginResponse;
  removeFromCart: Cart;
  startPayment: PaymentStartResponse;
};


export type MutationAddToCartArgs = {
  productId: Scalars['ID'];
};


export type MutationRemoveFromCartArgs = {
  cartItemId: Scalars['ID'];
};

export type Order = {
  __typename?: 'Order';
  createdAt: Scalars['String'];
  orderItems?: Maybe<Array<Product>>;
  paymentAmount: Scalars['Float'];
};

export type PaymentStartResponse = {
  __typename?: 'PaymentStartResponse';
  url: Scalars['String'];
};

export type Product = {
  __typename?: 'Product';
  description: Scalars['String'];
  id: Scalars['ID'];
  image: Scalars['String'];
  price: Scalars['Float'];
  title: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  cart: Cart;
  orders?: Maybe<Array<Order>>;
  product: Product;
  products?: Maybe<Array<Product>>;
};


export type QueryProductArgs = {
  ID: Scalars['ID'];
};

export type CartContentQueryVariables = Exact<{ [key: string]: never; }>;


export type CartContentQuery = { __typename?: 'Query', cart: { __typename?: 'Cart', items?: Maybe<Array<{ __typename?: 'CartItem', id: string, title: string, description: string, price: number, image: string }>> } };

export type CreatePaymentLinkMutationVariables = Exact<{ [key: string]: never; }>;


export type CreatePaymentLinkMutation = { __typename?: 'Mutation', startPayment: { __typename?: 'PaymentStartResponse', url: string } };

export type RemoveFromCartMutationVariables = Exact<{
  cartItemId: Scalars['ID'];
}>;


export type RemoveFromCartMutation = { __typename?: 'Mutation', removeFromCart: { __typename?: 'Cart', itemsCount: number } };

export type LoginMutationVariables = Exact<{ [key: string]: never; }>;


export type LoginMutation = { __typename?: 'Mutation', login: { __typename?: 'LoginResponse', id: string, avatar: string, fullName: string, token: string } };

export type ListOrdersQueryVariables = Exact<{ [key: string]: never; }>;


export type ListOrdersQuery = { __typename?: 'Query', orders?: Maybe<Array<{ __typename?: 'Order', paymentAmount: number, createdAt: string, orderItems?: Maybe<Array<{ __typename?: 'Product', id: string, title: string, description: string, price: number, image: string }>> }>> };

export type AddItemToCartMutationVariables = Exact<{
  productId: Scalars['ID'];
}>;


export type AddItemToCartMutation = { __typename?: 'Mutation', addToCart: { __typename?: 'Cart', itemsCount: number } };

export type ProductDetailQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type ProductDetailQuery = { __typename?: 'Query', product: { __typename?: 'Product', id: string, title: string, description: string, price: number, image: string } };

export type ListProductsQueryVariables = Exact<{ [key: string]: never; }>;


export type ListProductsQuery = { __typename?: 'Query', products?: Maybe<Array<{ __typename?: 'Product', id: string, title: string, description: string, price: number, image: string }>> };


export const CartContentDocument = gql`
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

/**
 * __useCartContentQuery__
 *
 * To run a query within a React component, call `useCartContentQuery` and pass it any options that fit your needs.
 * When your component renders, `useCartContentQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCartContentQuery({
 *   variables: {
 *   },
 * });
 */
export function useCartContentQuery(baseOptions?: Apollo.QueryHookOptions<CartContentQuery, CartContentQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<CartContentQuery, CartContentQueryVariables>(CartContentDocument, options);
      }
export function useCartContentLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<CartContentQuery, CartContentQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<CartContentQuery, CartContentQueryVariables>(CartContentDocument, options);
        }
export type CartContentQueryHookResult = ReturnType<typeof useCartContentQuery>;
export type CartContentLazyQueryHookResult = ReturnType<typeof useCartContentLazyQuery>;
export type CartContentQueryResult = Apollo.QueryResult<CartContentQuery, CartContentQueryVariables>;
export const CreatePaymentLinkDocument = gql`
    mutation createPaymentLink {
  startPayment {
    url
  }
}
    `;
export type CreatePaymentLinkMutationFn = Apollo.MutationFunction<CreatePaymentLinkMutation, CreatePaymentLinkMutationVariables>;

/**
 * __useCreatePaymentLinkMutation__
 *
 * To run a mutation, you first call `useCreatePaymentLinkMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreatePaymentLinkMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createPaymentLinkMutation, { data, loading, error }] = useCreatePaymentLinkMutation({
 *   variables: {
 *   },
 * });
 */
export function useCreatePaymentLinkMutation(baseOptions?: Apollo.MutationHookOptions<CreatePaymentLinkMutation, CreatePaymentLinkMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreatePaymentLinkMutation, CreatePaymentLinkMutationVariables>(CreatePaymentLinkDocument, options);
      }
export type CreatePaymentLinkMutationHookResult = ReturnType<typeof useCreatePaymentLinkMutation>;
export type CreatePaymentLinkMutationResult = Apollo.MutationResult<CreatePaymentLinkMutation>;
export type CreatePaymentLinkMutationOptions = Apollo.BaseMutationOptions<CreatePaymentLinkMutation, CreatePaymentLinkMutationVariables>;
export const RemoveFromCartDocument = gql`
    mutation removeFromCart($cartItemId: ID!) {
  removeFromCart(cartItemId: $cartItemId) {
    itemsCount
  }
}
    `;
export type RemoveFromCartMutationFn = Apollo.MutationFunction<RemoveFromCartMutation, RemoveFromCartMutationVariables>;

/**
 * __useRemoveFromCartMutation__
 *
 * To run a mutation, you first call `useRemoveFromCartMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useRemoveFromCartMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [removeFromCartMutation, { data, loading, error }] = useRemoveFromCartMutation({
 *   variables: {
 *      cartItemId: // value for 'cartItemId'
 *   },
 * });
 */
export function useRemoveFromCartMutation(baseOptions?: Apollo.MutationHookOptions<RemoveFromCartMutation, RemoveFromCartMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<RemoveFromCartMutation, RemoveFromCartMutationVariables>(RemoveFromCartDocument, options);
      }
export type RemoveFromCartMutationHookResult = ReturnType<typeof useRemoveFromCartMutation>;
export type RemoveFromCartMutationResult = Apollo.MutationResult<RemoveFromCartMutation>;
export type RemoveFromCartMutationOptions = Apollo.BaseMutationOptions<RemoveFromCartMutation, RemoveFromCartMutationVariables>;
export const LoginDocument = gql`
    mutation login {
  login {
    id
    avatar
    fullName
    token
  }
}
    `;
export type LoginMutationFn = Apollo.MutationFunction<LoginMutation, LoginMutationVariables>;

/**
 * __useLoginMutation__
 *
 * To run a mutation, you first call `useLoginMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useLoginMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [loginMutation, { data, loading, error }] = useLoginMutation({
 *   variables: {
 *   },
 * });
 */
export function useLoginMutation(baseOptions?: Apollo.MutationHookOptions<LoginMutation, LoginMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<LoginMutation, LoginMutationVariables>(LoginDocument, options);
      }
export type LoginMutationHookResult = ReturnType<typeof useLoginMutation>;
export type LoginMutationResult = Apollo.MutationResult<LoginMutation>;
export type LoginMutationOptions = Apollo.BaseMutationOptions<LoginMutation, LoginMutationVariables>;
export const ListOrdersDocument = gql`
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

/**
 * __useListOrdersQuery__
 *
 * To run a query within a React component, call `useListOrdersQuery` and pass it any options that fit your needs.
 * When your component renders, `useListOrdersQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useListOrdersQuery({
 *   variables: {
 *   },
 * });
 */
export function useListOrdersQuery(baseOptions?: Apollo.QueryHookOptions<ListOrdersQuery, ListOrdersQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ListOrdersQuery, ListOrdersQueryVariables>(ListOrdersDocument, options);
      }
export function useListOrdersLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ListOrdersQuery, ListOrdersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ListOrdersQuery, ListOrdersQueryVariables>(ListOrdersDocument, options);
        }
export type ListOrdersQueryHookResult = ReturnType<typeof useListOrdersQuery>;
export type ListOrdersLazyQueryHookResult = ReturnType<typeof useListOrdersLazyQuery>;
export type ListOrdersQueryResult = Apollo.QueryResult<ListOrdersQuery, ListOrdersQueryVariables>;
export const AddItemToCartDocument = gql`
    mutation addItemToCart($productId: ID!) {
  addToCart(productId: $productId) {
    itemsCount
  }
}
    `;
export type AddItemToCartMutationFn = Apollo.MutationFunction<AddItemToCartMutation, AddItemToCartMutationVariables>;

/**
 * __useAddItemToCartMutation__
 *
 * To run a mutation, you first call `useAddItemToCartMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddItemToCartMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addItemToCartMutation, { data, loading, error }] = useAddItemToCartMutation({
 *   variables: {
 *      productId: // value for 'productId'
 *   },
 * });
 */
export function useAddItemToCartMutation(baseOptions?: Apollo.MutationHookOptions<AddItemToCartMutation, AddItemToCartMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddItemToCartMutation, AddItemToCartMutationVariables>(AddItemToCartDocument, options);
      }
export type AddItemToCartMutationHookResult = ReturnType<typeof useAddItemToCartMutation>;
export type AddItemToCartMutationResult = Apollo.MutationResult<AddItemToCartMutation>;
export type AddItemToCartMutationOptions = Apollo.BaseMutationOptions<AddItemToCartMutation, AddItemToCartMutationVariables>;
export const ProductDetailDocument = gql`
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

/**
 * __useProductDetailQuery__
 *
 * To run a query within a React component, call `useProductDetailQuery` and pass it any options that fit your needs.
 * When your component renders, `useProductDetailQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useProductDetailQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useProductDetailQuery(baseOptions: Apollo.QueryHookOptions<ProductDetailQuery, ProductDetailQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ProductDetailQuery, ProductDetailQueryVariables>(ProductDetailDocument, options);
      }
export function useProductDetailLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ProductDetailQuery, ProductDetailQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ProductDetailQuery, ProductDetailQueryVariables>(ProductDetailDocument, options);
        }
export type ProductDetailQueryHookResult = ReturnType<typeof useProductDetailQuery>;
export type ProductDetailLazyQueryHookResult = ReturnType<typeof useProductDetailLazyQuery>;
export type ProductDetailQueryResult = Apollo.QueryResult<ProductDetailQuery, ProductDetailQueryVariables>;
export const ListProductsDocument = gql`
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

/**
 * __useListProductsQuery__
 *
 * To run a query within a React component, call `useListProductsQuery` and pass it any options that fit your needs.
 * When your component renders, `useListProductsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useListProductsQuery({
 *   variables: {
 *   },
 * });
 */
export function useListProductsQuery(baseOptions?: Apollo.QueryHookOptions<ListProductsQuery, ListProductsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ListProductsQuery, ListProductsQueryVariables>(ListProductsDocument, options);
      }
export function useListProductsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ListProductsQuery, ListProductsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ListProductsQuery, ListProductsQueryVariables>(ListProductsDocument, options);
        }
export type ListProductsQueryHookResult = ReturnType<typeof useListProductsQuery>;
export type ListProductsLazyQueryHookResult = ReturnType<typeof useListProductsLazyQuery>;
export type ListProductsQueryResult = Apollo.QueryResult<ListProductsQuery, ListProductsQueryVariables>;