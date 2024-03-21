<script lang="ts">
	let hours = 0;
	let minutes = 0;
	let seconds = 5;
	let totalSeconds: number;
	let intervalID: ReturnType<typeof setInterval> | undefined;
	let lastTimer: number;
	let isRunning: boolean = false;
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
		if (!isRunning) {
			isRunning = true;
			lastTimer = totalSeconds;
			intervalID = setInterval(removeSecond, 1000);
		}
	};

	const resetTimer = () => {
		isRunning = false;
		clearInterval(intervalID);
		hours = Math.floor(lastTimer / 3600);
		minutes = Math.floor((lastTimer % 3600) / 60);
		seconds = lastTimer % 60;
	};

	const adjustTime = (timeType: string, value: number) => {
		switch (timeType) {
			case 'hours':
				if (hours + value >= 0) hours += value;
				break;
			case 'minutes':
				if (minutes + value >= 0) minutes += value;
				break;
			case 'seconds':
				if (seconds + value >= 0) seconds += value;
				break;
		}
	};
	$: if (totalSeconds === 0) {
		isRunning = false;
		clearInterval(intervalID);
	}
</script>

<section class="flex items-center justify-center flex-col">
	<h2 class=" font-bold text-9xl">Timer</h2>

	<p>Countdown:</p>
	<div class="flex gap-20 text-3xl font-bold mt-20">
		<div class="flex flex-col justify-center items-center">
			<button on:click={() => adjustTime('hours', 1)}>+</button>
			<p>{hours} Stunden</p>
			<button on:click={() => adjustTime('hours', -1)}>-</button>
		</div>
		<div class="flex flex-col justify-center items-center">
			<button on:click={() => adjustTime('minutes', 1)}>+</button>
			<p>{minutes} Minuten</p>
			<button on:click={() => adjustTime('minutes', -1)}>-</button>
		</div>
		<div class="flex flex-col justify-center items-center">
			<button on:click={() => adjustTime('seconds', 1)}>+</button>
			<p>{seconds} Sekunden</p>
			<button on:click={() => adjustTime('seconds', -1)}>-</button>
		</div>
	</div>

	<button class="text-2xl mt-10" on:click={() => startTimer()}>Start Timer</button>
	<button class="text-2xl mt-10" on:click={() => resetTimer()}>Reset Timer</button>
</section>
