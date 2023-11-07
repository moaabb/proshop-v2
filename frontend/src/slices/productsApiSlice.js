import { PRODUCTS_URL } from '../constants';
import { apiSlice } from './apiSlice';

export const productsApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    getProducts: builder.query({
      query: ({ keyword, pageNumber }) => ({
        url: `${PRODUCTS_URL}/v1/products`,
        params: { keyword, pageNumber },
      }),
      keepUnusedDataFor: 5,
      providesTags: ['Products'],
    }),
    getProductDetails: builder.query({
      query: (productId) => ({
        url: `${PRODUCTS_URL}/v1/products/${productId}`,
      }),
      keepUnusedDataFor: 5,
    }),
    createProduct: builder.mutation({
      query: (product, token) => ({
        url: `${PRODUCTS_URL}/v1/products`,
        method: 'POST',
        body: product,
        headers: { Authorization: `Bearer ${token}` },
      }),
      invalidatesTags: ['Product'],
    }),
    updateProduct: builder.mutation({
      query: (data, token) => ({
        url: `${PRODUCTS_URL}/v1/products/${data.productId}`,
        method: 'PUT',
        body: data,
        headers: { Authorization: `Bearer ${token}` },
      }),
      invalidatesTags: ['Products'],
    }),
    uploadProductImage: builder.mutation({
      query: (data, token) => ({
        url: `/api/upload`,
        method: 'POST',
        body: data,
        headers: { Authorization: `Bearer ${token}` },
      }),
    }),
    deleteProduct: builder.mutation({
      query: (productId, token) => ({
        url: `${PRODUCTS_URL}/v1/products/${productId}`,
        method: 'DELETE',
        headers: { Authorization: `Bearer ${token}` },
      }),
      providesTags: ['Product'],
    }),
    createReview: builder.mutation({
      query: (data, token) => ({
        url: `${PRODUCTS_URL}/v1/products/${data.productId}/reviews`,
        method: 'POST',
        body: data,
        headers: { Authorization: `Bearer ${token}` },
      }),
      invalidatesTags: ['Product'],
    }),
    getTopProducts: builder.query({
      query: () => `${PRODUCTS_URL}/v1/products/top`,
      keepUnusedDataFor: 5,
    }),
  }),
});

export const {
  useGetProductsQuery,
  useGetProductDetailsQuery,
  useCreateProductMutation,
  useUpdateProductMutation,
  useUploadProductImageMutation,
  useDeleteProductMutation,
  useCreateReviewMutation,
  useGetTopProductsQuery,
} = productsApiSlice;
