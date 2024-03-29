<script lang="ts">
	import { Navigation } from '$lib/components';

	export let data;
	let users = data.body;
</script>

<Navigation />
<section class="container">
	<h1 class="font-bold text-4xl text-center">Ranking</h1>
	<div class="flex flex-col px-[410px] mt-20">
		{#each users as user, index}
			<div class="flex gap-20 justify-between">
				<h2>Platz {index + 1}</h2>
				<h3 class="flex-grow">{user.ID}</h3>
				<h4 class="text-right">Erfahrung: {user.experience}</h4>
			</div>
		{/each}
	</div>
</section>
