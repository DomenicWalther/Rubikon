
export const processUserUpdates = async (userExperience: number) => {
		const response = await fetch('http://localhost:3000/Users/manage-daily-streak', {
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