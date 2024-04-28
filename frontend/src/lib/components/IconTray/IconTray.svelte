<script>
	import Shop from '$lib/Svg/Shop.svelte';
	import Target from '$lib/Svg/Target.svelte';
	import Hanger from '$lib/Svg/Hanger.svelte';
	import Arrow from '$lib/Svg/Arrow.svelte';
	import { onMount } from 'svelte';

	let isIconTrayVisible = false;
	let iconTrayButton, iconTray;

	function toggleIconTrayVisibility() {
		isIconTrayVisible = !isIconTrayVisible;
	}

	onMount(() => {
		window.addEventListener('click', handleOutSideClick);
		return () => {
			window.removeEventListener('click', handleOutSideClick);
		};
	});

	function handleOutSideClick(event) {
		if (!iconTrayButton.contains(event.target) && !iconTray.contains(event.target)) {
			isIconTrayVisible = false;
		}
	}
</script>

<div class="absolute Icontray">
	{#if isIconTrayVisible}
		<div
			bind:this={iconTray}
			class="flex flex-col justify-center items-center bg-mainblue gap-6 rounded-xl py-4 pl-2 pr-2"
		>
			<Shop />
			<Target />
			<Hanger />
		</div>
	{:else}
		<div>
			<button on:click={toggleIconTrayVisibility}>
				<Arrow />
			</button>
		</div>
	{/if}
</div>
