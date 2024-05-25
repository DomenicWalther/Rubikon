<script lang="ts">
	export let type: string | null = null;
	export let skin_type: string;
	export let isShowcase = false;
</script>

{#if type}
	{#await import(`./${skin_type}/${type}.svelte`) then module}
		{#if module.default}
			<svelte:component this={module.default} {isShowcase} />
		{/if}
	{/await}
{/if}
