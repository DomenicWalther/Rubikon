<script lang="ts">
	import { Navigation, NavigationBottom } from '$lib/components';

	export let data;
	let users = data.body;
</script>

<section class="container p-0 flex flex-col">
	<div class="bg-mainyellow flex justify-center p-4 mb-1">
		<h1 class="font-medium text-3xl">Ranking</h1>
	</div>
	<div class="flex justify-center">
		<div class="border-b-2 border-black pl-6 pr-6 pt-3 pb-3">
			<h2 class="font-medium text-xl">Letzter Monat</h2>
		</div>
		<div class=" border-black pl-6 pr-6 pt-3 pb-3">
			<h2 class="font-medium text-xl text-black text-opacity-50">Gesamtübersicht</h2>
		</div>
		
	</div>

	<div class="flex flex-col items-center mt-1">
		{#each users as user, index}
			<div class="flex gap-20 mt-8 justify-evenly items-center">
				<h2>Platz {index + 1}</h2>
				<img src={user.profile_image_url} alt="Profile Picture" class="w-10 h-10 rounded-full" />
				<h3 class="flex-grow">{user.username === null ? user.first_name : user.username}</h3>
				<h4 class="text-right">Erfahrung: {user.experience}</h4>
			</div>
		{/each}
	</div>

</section>

<NavigationBottom activeItem="leaderboard" />
