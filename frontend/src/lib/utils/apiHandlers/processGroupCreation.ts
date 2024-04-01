interface GroupData {
    name: string;
    description: string;
    isPrivate: boolean;
    imageURL: string;
}


export const processGroupCreation = async (groupData: GroupData) => {
    const response = await fetch('http://localhost:3000/Groups', {
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