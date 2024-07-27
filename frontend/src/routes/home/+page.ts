import type { PageLoad } from './$types';
const API_URL = import.meta.env.VITE_API_URL

export const load: PageLoad = async ({ params }) => {
    const response = await fetch(`${API_URL}/api/home`);
    // const response = await fetch(`http://localhost:8000/api/home`);
    const data = await response.json();
    return {
        events: data.events,
        year: data.year
    };
};

