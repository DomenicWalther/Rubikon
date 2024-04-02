export async function load(event) {
	const user_id: string = event.locals.session.userId
	try {
		const groups = await fetch(`http://localhost:3000/Groups/${user_id}`);
		const data = await groups.json();
		return {
			body: data
		}
	} catch (e) {
		return {
			status: 500,
			body: {
				message: 'Failed to fetch Group data'
			}
		};
	}
}
