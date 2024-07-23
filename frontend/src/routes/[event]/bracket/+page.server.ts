import type { Actions } from './$types';

export const actions = {
    update: async ({ cookies, request }) => {
        const data = await request.formData();
        const teamId = data.get('teamId') as string;

        const response = await fetch("http://localhost:8000/api/update-elimination", {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ teamId: teamId })
        });
    },
} satisfies Actions;