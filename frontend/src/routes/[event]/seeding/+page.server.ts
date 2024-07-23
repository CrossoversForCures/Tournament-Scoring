import type { Actions } from './$types';

export const actions = {
    start: async ({ params }) => {
        const response = await fetch(`http://localhost:8000/api/${params.event}/start-elimination`, {
            method: "POST",
        });
    }
} satisfies Actions;