import type { PageLoad } from './$types';
const API_URL = import.meta.env.VITE_API_URL

export const load: PageLoad = async ({ params }) => {
    const response = await fetch(`${API_URL}/api/${params.event}/teams`);
    if (response.status == 404) {
        return {
            teams: null,
        };
    }
    const data = await response.json();
    return {
        teams: data,
    };
};