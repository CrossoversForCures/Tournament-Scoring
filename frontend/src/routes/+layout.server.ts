import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ cookies }) => {
    const isAdmin = cookies.get('session') === 'admin'
    console.log(isAdmin)
    return { isAdmin }
};