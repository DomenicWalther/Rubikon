import { PUBLIC_BACKEND_URL } from "$env/static/public";

export async function load(event) {
 const user_id:string = event.locals.session.userId
const response = await fetch(`${PUBLIC_BACKEND_URL}/Users/${user_id}`, {
			method: 'GET',
		});
		const user_data = await response.json();
        return {
            user: user_data
        }
}