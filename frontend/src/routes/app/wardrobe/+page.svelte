<script lang="ts">
	import { RoboterSkin } from '$lib/components';
	import { activeSkin } from '$lib/store.js';
	import SkinImport from '$lib/Svg/Skins/SkinImport.svelte';
	const skin_types = ['hats', 'eyes', 'mouths'];
	export let data;
	console.log(data);

	const changeSkin = (skin) => {
		switch (skin.type) {
			case 'hats':
				$activeSkin.activeHats = skin.image;
				break;
			case 'eyes':
				$activeSkin.activeEyes = skin.image;
				break;
			case 'mouths':
				$activeSkin.activeMouths = skin.image;
				break;
			default:
				break;
		}
	};
	console.log($activeSkin);
</script>

<h1>Deine aktuellen Skins</h1>
<RoboterSkin />

<div class="flex gap-20">
	{#each skin_types as skin_type}
		{#each data.userSkins as skin}
			{#if skin.type === skin_type}
				<button on:click={() => changeSkin(skin)} class="flex flex-col items-center">
					<h1>{skin.name}</h1>
					<SkinImport type={skin.image} skin_type={skin.type} isShowcase={true} />
				</button>
			{/if}
		{/each}
	{/each}
</div>
