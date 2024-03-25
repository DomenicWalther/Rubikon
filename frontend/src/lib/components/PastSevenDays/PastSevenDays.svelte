<script lang="ts">
	interface weekDayObject {
		label: string;
		isStreak: boolean;
	}

	export let streakLength: number = Math.min(4, 7);
	const today = new Date();
	const weekdayNames = [
		'Sonntag',
		'Montag',
		'Dienstag',
		'Mittwoch',
		'Donnerstag',
		'Freitag',
		'Samstag'
	];
	let allWeekDays: weekDayObject[] = [];
	const todaysIndex = today.getDay();
	for (let i = 0; i < 7; i++) {
		allWeekDays.push({ label: weekdayNames[(todaysIndex + i) % 7].slice(0, 2), isStreak: false });
	}

	allWeekDays.push(allWeekDays.shift() as weekDayObject);
	for (let i = 0; i < streakLength; i++) {
		allWeekDays[6 - i].isStreak = true;
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
