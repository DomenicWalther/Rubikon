<script lang="ts">
	let hours = 0;
	let minutes = 59;
	let seconds = 30;
	let totalSeconds;
	$: totalSeconds = hours * 3600 + minutes * 60 + seconds;

	const removeSecond = () => {
		seconds -= 1;
	};
</script>

<h2>Timer</h2>

<p>Countdown:</p>
<p>{hours} Stunden</p>
<p>{minutes} Minuten</p>
<p>{seconds} Sekunden</p>
<p>{totalSeconds} Totale Sekunden</p>

<button on:click={removeSecond}>Remove a second</button>

