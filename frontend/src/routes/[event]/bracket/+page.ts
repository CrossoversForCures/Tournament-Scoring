import type { PageLoad } from './$types';

export const load: PageLoad = async ({ params }) => {
    const response = await fetch(`http://localhost:8000/api/${params.event}/bracket`);
    if (response.status == 404) {
        return {
            games: null
        };
    }
    const data = await response.json();
    return {
        event: data.event,
        rounds: data.rounds,
        courts: data.courts,
        root: data.root,
    };
};