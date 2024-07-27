import type { Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
const ADMIN_USERNAME = import.meta.env.VITE_ADMIN_USERNAME
const ADMIN_PASSWORD = import.meta.env.VITE_ADMIN_PASSWORD

export const actions = {
    default: async ({ cookies, request }) => {
        const data = await request.formData();
        const username = data.get('username');
        const password = data.get('password');

        if (username && password && username === ADMIN_USERNAME && password === ADMIN_PASSWORD) {
            cookies.set('session', 'admin', { path: '/', httpOnly: false, sameSite: 'strict', secure: false, maxAge: 60 * 60 * 24 * 7 });
            throw redirect(303, '/');
        }
        return fail(400, { incorrect: true });
    }
} satisfies Actions;