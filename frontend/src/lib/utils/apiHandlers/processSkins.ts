import { PUBLIC_BACKEND_URL } from "$env/static/public";

export const processUserBuySkin = async (skinID: string) => {
        const response = await fetch(`${PUBLIC_BACKEND_URL}/Skins/Shop`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${await Clerk.session.getToken()}`
            },
            body: JSON.stringify({ skinID: skinID })
        });
        const data = await response.json();
        return data
    }

