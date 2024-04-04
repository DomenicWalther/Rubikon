import { PUBLIC_BACKEND_URL } from "$env/static/public";


interface GroupData {
	name: string;
	description: string;
	isPrivate: boolean;
	imageURL: string;
}


export const processGroupCreation = async (groupData: GroupData) => {
    const response = await fetch(`${PUBLIC_BACKEND_URL}/Groups`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${await Clerk.session.getToken()}`
        },
        body: JSON.stringify(groupData)
    });
    const data = await response.json()
    return data
    }


export const processJoinGroup = async(groupId: string) => {
    const response = await fetch(`${PUBLIC_BACKEND_URL}/Groups/join`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${await Clerk.session.getToken()}`
        },
        body: JSON.stringify({group_id: groupId})
    });
    const data = await response.json()
    return data
}


export const leaveGroup = async (groupId: number) => {
	const response = await fetch(`${PUBLIC_BACKEND_URL}/Groups/leave`, {
		method: 'DELETE',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${await Clerk.session.getToken()}`
		},
		body: JSON.stringify({ group_id: groupId })
	});
	const data = await response.json();
	return data
}
