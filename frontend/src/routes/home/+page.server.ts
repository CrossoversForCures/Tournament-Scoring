import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';

export const actions = {
    default: async ({ cookies, request }) => {
        const data = await request.formData();
        const username = data.get('username');
        const password = data.get('password');

        if (username === 'admin' && password === 'password') {
            cookies.set('session', 'admin', { path: '/', httpOnly: true, sameSite: 'strict', maxAge: 60 * 60 * 24 * 7 });
            throw redirect(303, '/');
        }

        return fail(400, { incorrect: true });
    }
} satisfies Actions;