<script lang="ts">
	import { Navigation, NavigationBottom } from '$lib/components';

	export let data;
	let users = data.body;
</script>

<section class="container">
	<h1 class="font-bold text-4xl text-center">Ranking</h1>
	<div class="flex flex-col px-[410px] mt-20">
		{#each users as user, index}
			<div class="flex gap-20 mt-5 w-96 content-around">
				<h2>Platz {index + 1}</h2>
				<img src={user.profile_image_url} alt="Profile Picture" class="w-10 h-10 rounded-full" />
				<h3 class="flex-grow">{user.username === null ? user.first_name : user.username}</h3>
				<h4 class="text-right">Erfahrung: {user.experience}</h4>
			</div>
		{/each}
	</div>
</section>

<NavigationBottom />
