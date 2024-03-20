<script lang="ts">
	let hours = 0;
	let minutes = 0;
	let seconds = 5;
	let totalSeconds: number;
	$: totalSeconds = hours * 3600 + minutes * 60 + seconds;

	const removeSecond = () => {
		if (seconds > 0) {
			seconds--;
		} else if (minutes > 0) {
			minutes--;
			seconds = 59;
		} else if (hours > 0) {
			hours--;
			minutes = 59;
			seconds = 59;
		}
	};
</script>

<h2>Timer</h2>

<p>Countdown:</p>
<p>{hours} Stunden</p>
<p>{minutes} Minuten</p>
<p>{seconds} Sekunden</p>
<p>{totalSeconds} Totale Sekunden</p>

<button on:click={removeSecond}>Remove a second</button>

