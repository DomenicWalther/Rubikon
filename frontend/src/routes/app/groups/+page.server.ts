export async function load() {
    let groups;
    try {
        groups = await fetch('http://localhost:3000/Groups');
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