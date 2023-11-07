import { apiSlice } from './apiSlice';
import { ORDERS_URL, PAYPAL_URL } from '../constants';

export const orderApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    createOrder: builder.mutation({
      query: ({ order, token }) => ({
        url: `${ORDERS_URL}/v1/orders`,
        method: 'POST',
        body: order,
        headers: { Authorization: `Bearer ${token}` },
      }),
    }),
    getOrderDetails: builder.query({
      query: ({ orderId, token }) => ({
        url: `${ORDERS_URL}/v1/orders/${orderId}`,
        headers: { Authorization: `Bearer ${token}` },
      }),
      keepUnusedDataFor: 5,
    }),
    payOrder: builder.mutation({
      query: ({ orderId, details, token }) => ({
        url: `${ORDERS_URL}/v1/orders/${orderId}/pay`,
        method: 'PUT',
        body: details,
        headers: { Authorization: `Bearer ${token}` },
      }),
    }),
    getPaypalClientId: builder.query({
      query: () => ({
        url: PAYPAL_URL,
      }),
      keepUnusedDataFor: 5,
    }),
    getMyOrders: builder.query({
      query: (token) => ({
        url: `${ORDERS_URL}/v1/users/orders`,
        headers: { Authorization: `Bearer ${token}` },
      }),
      keepUnusedDataFor: 5,
    }),
    getOrders: builder.query({
      query: (token) => ({
        url: `${ORDERS_URL}/v1/orders`,
        headers: { Authorization: `Bearer ${token}` },
      }),
      keepUnusedDataFor: 5,
    }),
    deliverOrder: builder.mutation({
      query: ({ orderId, token }) => ({
        url: `${ORDERS_URL}/v1/orders/${orderId}/deliver`,
        method: 'PUT',
        headers: { Authorization: `Bearer ${token}` },
      }),
    }),
  }),
});

export const {
  useCreateOrderMutation,
  useGetOrderDetailsQuery,
  usePayOrderMutation,
  useGetPaypalClientIdQuery,
  useGetMyOrdersQuery,
  useGetOrdersQuery,
  useDeliverOrderMutation,
} = orderApiSlice;
