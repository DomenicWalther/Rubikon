<script>
	export let type;
	let size = 24;
</script>

{#if type && size}
	{#await import(`$lib/Svg/${type}.svelte`) then module}
		{#if module.default}
			<svelte:component this={module.default} width={size} height={size} />
		{/if}
	{/await}
{/if}
