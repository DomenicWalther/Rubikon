import { PUBLIC_BACKEND_URL } from "$env/static/public";


interface MessageData {
    message: string;
    group_id: string;
    channelID: string;
    username: string,
    profile_image_url: string,
}

export const createNewGroupMessage = async(messageData: MessageData) => {
    
        const response = await fetch(`${PUBLIC_BACKEND_URL}/Groups/chat/addMessagePusher`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${await Clerk.session.getToken()}`
        },
        body: JSON.stringify(messageData)
    });
    const data = await response.json()
    return data
    }
