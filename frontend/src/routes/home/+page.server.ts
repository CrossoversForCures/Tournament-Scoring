import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
    const response = await fetch('http://localhost:8000/api/home');
    const data = await response.json();
    return {
        events: data.events,
        year: data.year
    };
};

