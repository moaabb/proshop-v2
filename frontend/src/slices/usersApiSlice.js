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
                credentials: 'include',
            }),
        }),
        register: builder.mutation({
            query: (data) => ({
                url: `${USERS_URL}/v1/users`,
                method: 'POST',
                body: data,
            }),
        }),
        logout: builder.mutation({
            query: () => ({
                url: `${AUTH_URL}/v1/auth/logout`,
                method: 'POST',
                credentials: 'include',
            }),
        }),
        profile: builder.mutation({
            query: ({ data, id }) => ({
                url: `${USERS_URL}/v1/users/${id}`,
                method: 'PUT',
                body: data,
            }),
        }),
        getUsers: builder.query({
            query: () => ({
                url: `${USERS_URL}/v1/users`,
            }),
            providesTags: ['User'],
            keepUnusedDataFor: 5,
        }),
        deleteUser: builder.mutation({
            query: (userId) => ({
                url: `${USERS_URL}/v1/users/${userId}`,
                method: 'DELETE',
            }),
        }),
        updateUser: builder.mutation({
            query: (data) => ({
                url: `${USERS_URL}/v1/users/${data.userId}`,
                method: 'PUT',
                body: data,
            }),
            invalidatesTags: ['User'],
        }),
        getUserDetails: builder.query({
            query: (id) => ({
                url: `${USERS_URL}/v1/users/${id}`,
            }),
            keepUnusedDataFor: 5,
        }),
    }),
});

export const {
    useLoginMutation,
    useLogoutMutation,
    useRegisterMutation,
    useProfileMutation,
    useGetUsersQuery,
    useDeleteUserMutation,
    useUpdateUserMutation,
    useGetUserDetailsQuery,
} = userApiSlice;
