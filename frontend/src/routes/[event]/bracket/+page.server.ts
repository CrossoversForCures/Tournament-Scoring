import type { Actions } from './$types';
const API_URL = import.meta.env.VITE_API_URL

export const actions = {
    update: async ({ cookies, request }) => {
        const data = await request.formData();
        const teamId = data.get('teamId') as string;

        const response = await fetch(`${API_URL}/api/update-elimination`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ teamId: teamId })
        });
    },
} satisfies Actions;