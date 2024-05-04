<script lang="ts">
	export let type: string;
	export let color: string = "white";
	export let isActiveItem: boolean;
	let size = 24;
</script>

{#if type && size && color}
	{#await import(`$lib/Svg/${type}.svelte`) then module}
		{#if module.default}
			<svelte:component this={module.default} width={size} height={size} color={color} {isActiveItem} />
		{/if}
	{/await}
{/if}
