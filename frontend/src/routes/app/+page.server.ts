export async function load(event) {
 const user_id:string = event.locals.session.userId
const response = await fetch(`http://localhost:3000/Users/${user_id}`, {
			method: 'GET',
		});
		const user_data = await response.json();
        console.log(user_data)
        return {
            user: user_data
        }
}