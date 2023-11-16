import { apiSlice } from './apiSlice';
import { USERS_URL } from '../constants';
import { AUTH_URL } from '../constants';

export const userApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    login: builder.mutation({
      query: (data) => ({
        url: `${AUTH_URL}/v1/auth/login`,
        method: 'POST',
        body: data,
        credentials: "include"
      }),
    }),
    register: builder.mutation({
      query: (data) => ({
        url: `${USERS_URL}/v1/users`,
        method: 'POST',
        body: data,
      }),
    }),
    // logout: builder.mutation({
    //   query: () => ({
    //     url: `${USERS_URL}/logout`,
    //     method: 'POST',
    //   }),
    // }),
    profile: builder.mutation({
      query: ({ data, token, id }) => ({
        url: `${USERS_URL}/v1/users/${id}`,
        method: 'PUT',
        body: data,
        headers: { Authorization: `Bearer ${token}` },
      }),
    }),
    getUsers: builder.query({
      query: (token) => ({
        url: USERS_URL,
        headers: { Authorization: `Bearer ${token}` },
      }),
      providesTags: ['User'],
      keepUnusedDataFor: 5,
    }),
    deleteUser: builder.mutation({
      query: ({ userId, token }) => ({
        url: `${USERS_URL}/v1/users/${userId}`,
        method: 'DELETE',
        headers: { Authorization: `Bearer ${token}` },
      }),
    }),
    getUserDetails: builder.query({
      query: ({ id, token }) => ({
        url: `${USERS_URL}/v1/users/${id}`,
        headers: { Authorization: `Bearer ${token}` },
      }),
      keepUnusedDataFor: 5,
    }),
    updateUser: builder.mutation({
      query: ({ data, token }) => ({
        url: `${USERS_URL}/v1/users/${data.userId}`,
        method: 'PUT',
        body: data,
        headers: { Authorization: `Bearer ${token}` },
      }),
      invalidatesTags: ['User'],
    }),
  }),
});

export const {
  useLoginMutation,
  // useLogoutMutation,
  useRegisterMutation,
  useProfileMutation,
  useGetUsersQuery,
  useDeleteUserMutation,
  useUpdateUserMutation,
  useGetUserDetailsQuery,
} = userApiSlice;
