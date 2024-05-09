<script lang="ts">
	import { Navigation, NavigationBottom } from '$lib/components';
</script>

<Navigation />

<main class="flex flex-col items-center bg-maingrey h-screen">
	<section class="pt-10">
		<h2 class="text-5xl font-bold">Sam's Fortschritt</h2>
	</section>

	<div class="">
		<NavigationBottom activeItem="progress"/>
	</div>
</main>
