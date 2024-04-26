<script>
	import Shop from '$lib/Svg/Shop.svelte';
	import Target from '$lib/Svg/Target.svelte';
	import Hanger from '$lib/Svg/Hanger.svelte';
	import Arrow from '$lib/Svg/Arrow.svelte';
	import { onMount } from 'svelte';

	let isIconTrayVisible = false

	function toggleIconTrayVisibility (){
		isIconTrayVisible = !isIconTrayVisible;
	}

	onMount (() =>{

	window.addEventListener ('click', handleOutSideClick);
		return() => {

	window.removeEventListener('click', handleOutSideClick);
		};
	});


	function
	handleOutSideClick(event) {
		const IconTray = document.querySelector ('.Icontray');
		const button = document.querySelector ('.button');

		if (!
	IconTray?.contains(event.target)
	&& !button?.contains (event.target))
	{
		isIconTrayVisible = false;
	}
	}


</script>



<div class="fixed left-0 flex flex-col justify-center items-center h-full">
	<button on:click={toggleIconTrayVisibility} class="button">
		<Arrow/>
	</button>
</div>


{#if isIconTrayVisible}
<div
	class="fixed left-0 flex flex-col justify-center items-center top-1/2 bg-mainblue gap-6 rounded-xl py-4 pl-2 pr-2 transform -translate-y-1/2">
	<Shop />
	<Target />
	<Hanger />
</div>
{/if}