
import { PUBLIC_BACKEND_URL } from "$env/static/public";

export async function load() {
    try {
        const skins = await fetch(`${PUBLIC_BACKEND_URL}/Skins/Shop`);
        const data = await skins.json();
        return {
            skins: data
        }
    } catch(e) {
        return {
            status: 500,
            body: {
                message: 'Failed to fetch Skin data'
            }
        };
    }
}