export async function load() {
    let leaderboardData;
    try {
        leaderboardData = await fetch('http://localhost:3000/Users/top-users');
        const data = await leaderboardData.json();
        data.sort((a, b) => b.experience - a.experience);
        for (const obj of data) {
            delete obj.id
        }
        return {
            body: data
        }
    } catch (e) {
        return {
            status: 500,
            body: {
                message: 'Failed to fetch leaderboard data'
            }
        };
    }
}