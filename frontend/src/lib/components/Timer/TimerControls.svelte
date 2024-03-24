<script lang="ts">
	export let label: string;
	export let value: number;
	export let increment: () => void;
	export let decrement: () => void;
</script>

<div class="flex flex-col justify-center items-center">
	<button on:click={increment}>+</button>
	<p>{value} {label}</p>
	<button on:click={decrement}>-</button>
</div>
