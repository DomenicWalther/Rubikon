<script lang="ts">
	export let value: number;
	export let increment: () => void;
	export let decrement: () => void;
	const zeroPad = (num, places) => String(num).padStart(places, '0');
</script>

<div class="flex flex-col justify-center items-center gap-5 font-normal font-digitalNumbers">
	<button on:click={increment}>+</button>
	<p class="font-digitalNumbers font-normal text-5xl">{zeroPad(value, 2)}</p>
	<button on:click={decrement}>-</button>
</div>
