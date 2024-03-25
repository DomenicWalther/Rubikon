<script lang="ts">
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
			<div class="font-bold text-black">{weekDayObject.label}</div>
			<div class={weekDayObject.isStreak ? 'h-8 w-8 bg-blue-500' : 'h-8 w-8 bg-slate-400'}></div>
		</div>
	{/each}
</div>
