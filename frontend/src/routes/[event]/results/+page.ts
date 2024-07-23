import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
    const response = await fetch(`http://localhost:8000/api/${params.event}/results`);
    if (response.status == 404) {
        return {
            results: null,
        };
    }
    const data = await response.json();
    return {
        results: data,
    };
};
