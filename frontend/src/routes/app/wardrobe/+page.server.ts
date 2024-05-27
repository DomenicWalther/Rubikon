import { PUBLIC_BACKEND_URL } from "$env/static/public";

export async function load(event) {
    const user_id: string = event.locals.session.userId
    try {
        const userSkins = await fetch(`${PUBLIC_BACKEND_URL}/Skins/Owned/${user_id}`);
        const data = await userSkins.json();
        return {
            userSkins: data
        }
    } catch (e) {
        console.log(e)
    }
}