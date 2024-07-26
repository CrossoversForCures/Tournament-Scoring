import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
    const response = await fetch('http://localhost:8000/api/home');
    console.log(response)
    const data = await response.json();
    return {
        events: data.events,
        year: data.year
    };
};

