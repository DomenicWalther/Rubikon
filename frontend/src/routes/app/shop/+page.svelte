<script lang="ts">
	import * as Card from '$lib/components/shadcn/ui/card';
	import SkinImport from '$lib/Svg/Skins/SkinImport.svelte';
	const skin_types = ['hats', 'eyes', 'mouths'];
	export let data;
	console.log(data);
</script>

{#each skin_types as skin_type}
	<h2 class="text-2xl font-bold text-black">{skin_type}</h2>
	<div class="flex gap-20">
		{#each data.skins as skin}
			{#if skin.type === skin_type}
				<div class="flex flex-col items-center">
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
				</div>
			{/if}
		{/each}
	</div>
{/each}
