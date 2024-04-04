import { PUBLIC_BACKEND_URL } from "$env/static/public";

export const processUserUpdates = async (userExperience: number) => {
		const response = await fetch(`${PUBLIC_BACKEND_URL}/Users/process-user-progress`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${await Clerk.session.getToken()}`
			},
            body: JSON.stringify({ userExperience: userExperience })
		});
		const data = await response.json();
        return data
	};