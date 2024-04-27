<script lang="ts">
	import TimerControls from './TimerControls.svelte';
	import { createEventDispatcher } from 'svelte';
	import { isTimerRunning } from '$lib/store';

	const dispatch = createEventDispatcher();

	function processUser(rubikonLength: number) {
		dispatch('processUser', rubikonLength as number);
	}

	let hours = 0;
	let minutes = 0;
	let seconds = 5;

	let totalSeconds: number;
	let intervalID: ReturnType<typeof setInterval> | undefined;
	let lastTimer: number;
	const SECONDS_IN_HOUR = 3600;
	const SECONDS_IN_MINUTE = 60;
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
		if (!$isTimerRunning) {
			$isTimerRunning = true;
			lastTimer = totalSeconds;
			intervalID = setInterval(removeSecond, 20);
		}
	};

	const resetTimer = () => {
		$isTimerRunning = false;
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
	$: if (totalSeconds === 0 && $isTimerRunning) {
		let userExperience = lastTimer - totalSeconds;
		processUser(userExperience);
		resetTimer();
	}
</script>

<div class="flex gap-10 text-3xl font-bold mt-20 items-center">
	<TimerControls
		value={minutes}
		increment={() => adjustTime('minutes', 1)}
		decrement={() => adjustTime('minutes', -1)}
	/>
	<p>:</p>
	<TimerControls
		value={seconds}
		increment={() => adjustTime('seconds', 1)}
		decrement={() => adjustTime('seconds', -1)}
	/>
</div>

<button
	class="text-2xl mt-10 border-white border-2 rounded-full py-4 px-12 font-medium"
	on:click={() => startTimer()}>STARTEN</button
>
<!--
<button class="text-2xl mt-10" on:click={() => resetTimer()}>Reset Timer</button>

//-->
