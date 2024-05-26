<script lang="ts">
	import * as Card from '$lib/components/shadcn/ui/card';
	import { processUserBuySkin } from '$lib/utils/apiHandlers/processSkins.js';
	import SkinImport from '$lib/Svg/Skins/SkinImport.svelte';
	import { Toaster } from '$lib/components/shadcn/ui/sonner';
	import { toast } from 'svelte-sonner';

	const skin_types = ['hats', 'eyes', 'mouths'];
	export let data;
	console.log(data.skins);

	const buySkin = async (skin_id) => {
		const response = await processUserBuySkin(skin_id);
		if (response.message !== undefined) {
			toast.error(response.message);
		} else {
			toast.success('Skin gekauft');
		}
		data.skins = data.skins.filter((skin) => skin.ID !== skin_id);
	};
</script>

<Toaster richColors />
{#each skin_types as skin_type}
	<h2 class="text-2xl font-bold text-black">{skin_type}</h2>
	<div class="flex gap-20">
		{#each data.skins as skin}
			{#if skin.type === skin_type}
				<button on:click={() => buySkin(skin.ID)} class="flex flex-col items-center">
					<Card.Root>
						<Card.Header>
							<Card.Title>Benötigtes Level: {skin.required_level}</Card.Title>
						</Card.Header>
						<Card.Content class="flex items-center gap-5  flex-col">
							<h1>{skin.name}</h1>
							<SkinImport type={skin.image} skin_type={skin.type} isShowcase={true} />
						</Card.Content>
					</Card.Root>
					<p class="text-lg">{skin.price}</p>
				</button>
			{/if}
		{/each}
	</div>
{/each}
