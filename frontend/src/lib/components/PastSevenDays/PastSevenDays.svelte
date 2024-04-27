<script lang="ts">

	import Checkmark from '$lib/Svg/Checkmark.svelte';

	interface weekDayObject {
		label: string;
		isStreak: boolean;
	}

	export let streakLength: number = 1;
	const weekdayNames = ['So', 'Mo', 'Di', 'Mi', 'Do', 'Fr', 'Sa'];
	let allWeekDays: weekDayObject[] = [];
	const todaysIndex = new Date().getDay() - 1;
	for (let i = 0; i < 7; i++) {
		allWeekDays.push({ label: weekdayNames[(i + 1) % 7], isStreak: false });
	}

	for (let i = todaysIndex; i >= 0; i--) {
		if (streakLength > 0) {
			allWeekDays[i].isStreak = true;
			streakLength -= 1;
		}
	}
</script>

<div class="flex gap-5">
	{#each allWeekDays as weekDayObject}
		<div class="flex items-center flex-col">
			<div class="font-medium text-black m-1">{weekDayObject.label}</div>
			<div class="h-8 w-8 rounded-md{weekDayObject.isStreak ? ' bg-mainlightblue  flex items-center justify-center' : ' bg-slate-400'}">
				{#if weekDayObject.isStreak}
				
				<Checkmark />

				{/if}
			</div>
		</div>
	{/each}
</div>




