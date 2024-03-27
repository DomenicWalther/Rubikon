export async function load() {
    let leaderboardData;
    try {
        leaderboardData = await fetch('http://localhost:3000/Users/top-users');
        const data = await leaderboardData.json();
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