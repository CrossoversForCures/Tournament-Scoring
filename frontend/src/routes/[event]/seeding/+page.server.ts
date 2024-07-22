import type { Actions } from './$types';

export const actions = {
    // default: async ({ cookies, request }) => {
    //     const data = await request.formData();
    //     const gameId = data.get('gameId') as string;
    //     const team1Score = Number(data.get('team1Score'));
    //     const team2Score = Number(data.get('team2Score'));

    //     const response = await fetch("http://localhost:8000/api/update-pool", {
    //         method: "POST",
    //         headers: {
    //             'Content-Type': 'application/json',
    //         },
    //         body: JSON.stringify({ gameId: gameId, team1Score: team1Score, team2Score: team2Score })
    //     });
    // },
    start: async ({ params }) => {
        const response = await fetch(`http://localhost:8000/api/${params.event}/start-elimination`, {
            method: "POST",
        });
    }
} satisfies Actions;