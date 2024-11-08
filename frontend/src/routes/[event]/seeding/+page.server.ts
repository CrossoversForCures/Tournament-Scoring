import type { Actions } from './$types';
const API_URL = import.meta.env.VITE_API_URL

export const actions = {
    start: async ({ params }) => {
        const response = await fetch(`${API_URL}/api/${params.event}/start-elimination`, {
            method: "POST",
        });
    }
} satisfies Actions;