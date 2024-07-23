import type { LayoutLoad } from './$types';
import { isAdmin } from '$lib/stores/admin';

export const load: LayoutLoad = async ({ data }) => {
    isAdmin.set(data.isAdmin);
};