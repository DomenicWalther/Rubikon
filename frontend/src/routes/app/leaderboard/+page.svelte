<script lang="ts">
	import { Navigation, NavigationBottom } from '$lib/components';
	import * as Tabs from "$lib/components/shadcn/ui/tabs"

	export let data;
	let leaderboardUsers = data.body;

	
</script>

<section class="container p-0 flex flex-col">
	<div class="bg-mainyellow flex justify-center p-4 mb-1">
		<h1 class="font-medium text-3xl">Ranking</h1>
	</div>
	<Tabs.Root value="lastMonth" class="flex justify-center">
		<Tabs.List>
			<Tabs.Trigger value="lastMonth" class="border-b-2 border-black pl-6 pr-6 pt-3 pb-3"><h2 class="font-medium text-xl">Letzter Monat</h2></Tabs.Trigger>
			<Tabs.Trigger value="allMonths">Gesamtübersicht</Tabs.Trigger>
		  </Tabs.List>
		  
		<Tabs.Content value="lastMonth">
			<div class="flex flex-col items-center mt-1">
				{#each leaderboardUsers as user, index}
					<div class="flex gap-20 mt-8 justify-evenly items-center">
						<h2>Platz {index + 1}</h2>
						<img src={user.profile_image_url} alt="Profile Picture" class="w-10 h-10 rounded-full" />
						<h3 class="flex-grow">{user.username === null ? user.first_name : user.username}</h3>
						<h4 class="text-right">Erfahrung: {user.monthly_experience}</h4>
					</div>
				{/each}
			</div>
		  </Tabs.Content>
		  <Tabs.Content value="allMonths">	<div class="flex flex-col items-center mt-1">
			{#each leaderboardUsers as user, index}
				<div class="flex gap-20 mt-8 justify-evenly items-center">
					<h2>Platz {index + 1}</h2>
					<img src={user.profile_image_url} alt="Profile Picture" class="w-10 h-10 rounded-full" />
					<h3 class="flex-grow">{user.username === null ? user.first_name : user.username}</h3>
					<h4 class="text-right">Erfahrung: {user.experience}</h4>
				</div>
			{/each}
		</div></Tabs.Content>
	</Tabs.Root>
		


	<div class="bg-mainyellow p-8 mt-10">
		<img src={data.user.profile_image} alt="Profile Picture" class="w-10 h-10 rounded-full" />
		<p>{data.user.username}</p>
		<h4 class="text-right">Erfahrung: {data.user.experience}</h4>
	</div>

	


</section>

<NavigationBottom activeItem="leaderboard" />
