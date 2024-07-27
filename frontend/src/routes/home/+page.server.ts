import type { Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';

const ADMIN_USERNAME = import.meta.env.VITE_ADMIN_USERNAME
const ADMIN_PASSWORD = import.meta.env.VITE_ADMIN_PASSWORD

console.log('Admin username is set:', !!ADMIN_USERNAME);
console.log('Admin password is set:', !!ADMIN_PASSWORD);

export const actions = {
    default: async ({ cookies, request }) => {
        console.log('Login attempt initiated');

        const data = await request.formData();
        const username = data.get('username');
        const password = data.get('password');

        console.log('Received username:', username);
        console.log('Password received:', !!password);

        if (username && password && username === ADMIN_USERNAME && password === ADMIN_PASSWORD) {
            console.log('Credentials match, setting cookie');
            
            cookies.set('session', 'admin', { 
                path: '/', 
                httpOnly: false, 
                sameSite: 'strict', 
                maxAge: 60 * 60 * 24 * 7 
            });
            
            console.log('Cookie set, redirecting');
            throw redirect(303, '/');
        }

        console.log('Login failed, returning error');
        return fail(400, { incorrect: true });
    }
} satisfies Actions;