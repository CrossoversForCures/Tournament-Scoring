import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
    const response = await fetch(`http://localhost:8000/api/${params.event}/teams`);
    const data = await response.json();
    return {
        teams: data,
    };
};