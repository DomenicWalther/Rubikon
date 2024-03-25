<script lang="ts">
	import TimerControls from './TimerControls.svelte';
	import * as Dialog from '$lib/components/shadcn/ui/dialog';

	let hours = 0;
	let minutes = 0;
	let seconds = 5;
	let totalSeconds: number;
	let intervalID: ReturnType<typeof setInterval> | undefined;
	let lastTimer: number;
	let dialogOpen = false;
	const SECONDS_IN_HOUR = 3600;
	const SECONDS_IN_MINUTE = 60;
	let isRunning: boolean = false;
	$: totalSeconds = hours * SECONDS_IN_HOUR + minutes * SECONDS_IN_MINUTE + seconds;

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
			intervalID = setInterval(removeSecond, 100);
		}
	};

	const resetTimer = () => {
		isRunning = false;
		clearInterval(intervalID);
		totalSeconds = lastTimer;
		hours = Math.floor(lastTimer / SECONDS_IN_HOUR);
		minutes = Math.floor((lastTimer % SECONDS_IN_HOUR) / SECONDS_IN_MINUTE);
		seconds = lastTimer % SECONDS_IN_MINUTE;
		lastTimer = totalSeconds;
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
	$: if (totalSeconds === 0 && isRunning) {
		handleUserStreakAfterTimer();
		resetTimer();
	}

	const handleUserStreakAfterTimer = async () => {
		const response = await fetch('http://localhost:3000/Users/manage-daily-streak', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${await Clerk.session.getToken()}`
			}
		});
		const data = await response.json();
		dialogOpen = true;
		console.log(data);
	};
</script>

<section class="flex items-center justify-center flex-col">
	<h2 class=" font-bold text-7xl">Timer</h2>

	<p>Countdown:</p>
	<div class="flex gap-20 text-3xl font-bold mt-20">
		<TimerControls
			label="Stunden"
			value={hours}
			increment={() => adjustTime('hours', 1)}
			decrement={() => adjustTime('hours', -1)}
		/>
		<TimerControls
			label="Minuten"
			value={minutes}
			increment={() => adjustTime('minutes', 1)}
			decrement={() => adjustTime('minutes', -1)}
		/>
		<TimerControls
			label="Sekunden"
			value={seconds}
			increment={() => adjustTime('seconds', 1)}
			decrement={() => adjustTime('seconds', -1)}
		/>
	</div>

	<button class="text-2xl mt-10" on:click={() => startTimer()}>Start Timer</button>
	<button class="text-2xl mt-10" on:click={() => resetTimer()}>Reset Timer</button>
	<Dialog.Root bind:open={dialogOpen}>
		<Dialog.Content>
			<Dialog.Header>
				<Dialog.Title>Herzlichen Glückwunsch</Dialog.Title>
				<Dialog.Description>Deine Streak ist jetzt bei 1. Weiter so!</Dialog.Description>
			</Dialog.Header>
		</Dialog.Content>
	</Dialog.Root>
</section>
