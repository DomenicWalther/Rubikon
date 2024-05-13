import { PUBLIC_BACKEND_URL } from "$env/static/public";
import type { PageServerLoad } from "../$types";

export const load: PageServerLoad = async({params}) => {
    const groupID = params.slug

const chatMessages = await fetch(`${PUBLIC_BACKEND_URL}/Groups/chat/${groupID}`, {
			method: 'GET',
		});
		const chatMessage = await chatMessages.json();
        return {
            messages: chatMessage,
            groupID
        }
}