<script lang="ts">
	import { NavigationBottom } from '$lib/components';
	import  ProgressRank  from './ProgressRank.svelte';

	export let data
	console.log(data.user)
	import Lightning from '$lib/Svg/Lightning.svelte';

</script>

<main>
	<section class="">
		<div class="bg-mainlightblue flex justify-center p-5 mb-1">
			<h1 class="text-white font-medium text-3xl">Fortschritt</h1>
		</div>
		<div class="mt-8">
			<h2 class="flex justify-center text-3xl font-medium ">{data.user.username.toUpperCase()}'s Fortschritt</h2>
		</div>
		<div class="mt-10">
			<ProgressRank currentRank="Rank_1"/>
		</div>
	</section>
	<NavigationBottom/>
</main>	