import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
    const response = await fetch(`http://localhost:8000/api/${params.event}/bracket`);
    const data = await response.json();
    return {
        root: data.root
    };
};

