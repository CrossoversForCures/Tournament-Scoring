import type { PageServerLoad, Actions } from './$types';
import { fail, redirect } from '@sveltejs/kit';
import env from '$lib/load-env';

export const actions = {
    default: async ({ cookies, request }) => {
        const data = await request.formData();
        const username = data.get('username');
        const password = data.get('password');

        if (username === env.ADMIN_USERNAME && password === env.ADMIN_PASSWORD) {
            cookies.set('session', 'admin', { path: '/', httpOnly: true, sameSite: 'strict', maxAge: 60 * 60 * 24 * 7 });
            throw redirect(303, '/');
        }

        return fail(400, { incorrect: true });
    }
} satisfies Actions;