<script lang="ts">
	let hours = 0;
	let minutes = 0;
	let seconds = 5;
	let totalSeconds: number;
	let intervalID: number | undefined;
	let lastTimer: number;
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

	const startTimer = () => {
		lastTimer = totalSeconds;
		intervalID = setInterval(removeSecond, 1000);
	};

	const resetTimer = () => {
		clearInterval(intervalID);
		hours = Math.floor(lastTimer / 3600);
		minutes = Math.floor((lastTimer % 3600) / 60);
		seconds = lastTimer % 60;
	};

	$: if (totalSeconds === 0) {
		clearInterval(intervalID);
	}
</script>

<section class="flex items-center justify-center flex-col">
	<h2>Timer</h2>

	<p>Countdown:</p>
	<div class="flex gap-20 text-3xl font-bold mt-20">
		<div class="flex flex-col justify-center items-center">
			<button on:click={() => (hours += 1)}>+</button>
			<p>{hours} Stunden</p>
			<button on:click={() => (hours -= 1)}>-</button>
		</div>
		<div class="flex flex-col justify-center items-center">
			<button on:click={() => (minutes += 1)}>+</button>
			<p>{minutes} Minuten</p>
			<button on:click={() => (minutes -= 1)}>-</button>
		</div>
		<div class="flex flex-col justify-center items-center">
			<button on:click={() => (seconds += 1)}>+</button>
			<p>{seconds} Sekunden</p>
			<button on:click={() => (seconds -= 1)}>-</button>
		</div>
	</div>

	<button class="text-2xl mt-10" on:click={() => startTimer()}>Start Timer</button>
	<button class="text-2xl mt-10" on:click={() => resetTimer()}>Reset Timer</button>
</section>
