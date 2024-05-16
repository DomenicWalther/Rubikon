<script lang="ts">
	import { NavigationBottom } from '$lib/components';
	import * as Tabs from "$lib/components/shadcn/ui/tabs"

	export let data;
	let leaderboardUsers = data.body;

	const findCurrentUserInLeaderboard = (array, stringToFind) => {
		return array.findIndex(obj => obj.username === stringToFind)
	}

	let currentUserPosition = findCurrentUserInLeaderboard(leaderboardUsers, data.user.username)


	
</script>

<section class="container p-0 flex flex-col">
	<div class="bg-mainyellow flex justify-center p-5 mb-1">
		<h1 class="font-medium text-3xl">Ranking</h1>
	</div>
	<Tabs.Root value="lastMonth" class="flex flex-col justify-evenly items-center mt-3">
		<Tabs.List>
			<Tabs.Trigger value="lastMonth" class="font-medium text-xl">Letzter Monat</Tabs.Trigger>
			<Tabs.Trigger value="allMonths" class="font-medium text-xl text-opacity-50 text-black">Gesamtübersicht</Tabs.Trigger>
		  </Tabs.List>
		  
		<Tabs.Content value="lastMonth">
				{#each leaderboardUsers as user, index}
					<div class="flex gap-16 mt-8 items-center">
						<h2>Platz {index + 1}</h2>
						<img src={user.profile_image_url} alt="Profile Picture" class="w-10 h-10 rounded-full" />
						<h3 class="flex-grow">{user.username === null ? user.first_name : user.username}</h3>
						<h4 class="text-right">Erfahrung: {user.monthly_experience}</h4>
					</div>
				{/each}
		  </Tabs.Content>

		  <Tabs.Content value="allMonths">
			{#each leaderboardUsers as user, index}
				<div class="flex gap-16 mt-8 justify-evenly items-center">
					<h2>Platz {index + 1}</h2>
					<img src={user.profile_image_url} alt="Profile Picture" class="w-10 h-10 rounded-full" />
					<h3 class="flex-grow">{user.username === null ? user.first_name : user.username}</h3>
					<h4 class="text-right">Erfahrung: {user.experience}</h4>
				</div>
			{/each}
			</Tabs.Content>
	</Tabs.Root>
		

	<div class="bg-mainyellow p-6 mt-10 flex justify-center">
		<div class="flex items-center gap-16">
			<h2>Platz {currentUserPosition + 1}</h2>
			<img src={data.user.profile_image} alt="Profile Picture" class="w-10 h-10 rounded-full" />
			<p>{data.user.username}</p>
			<h4 class="text-right">Erfahrung: {data.user.experience}</h4>
		</div>
	</div>

	<NavigationBottom activeItem="leaderboard" />

</section>

