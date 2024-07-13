// export const BASE_URL =
//   process.env.NODE_ENV === 'develeopment' ? 'http://localhost:5000' : '';
const DOMAIN = '192.168.0.110';
console.log(DOMAIN);
export const BASE_URL = `http://${DOMAIN}:8000`; // If using proxy
export const PRODUCTS_URL = '/products';
export const USERS_URL = '/users';
export const AUTH_URL = '/auth';
export const ORDERS_URL = '/orders';
export const PAYPAL_URL = '/auth/v1/config/paypal';
