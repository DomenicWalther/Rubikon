<script lang="ts">
	import TimerComponent from './TimerComponent.svelte';
	import { PastSevenDays } from '$lib/components';
	import { processUserUpdates } from '$lib/utils/apiHandlers/processUserUpdates';
	import * as Dialog from '$lib/components/shadcn/ui/dialog';
	import { Toaster } from '$lib/components/shadcn/ui/sonner';
	import { toast } from 'svelte-sonner';

	let dialogOpen = false;
	let streakLength = 0;

	async function processUserUpdatesAndShowDialog(rubikonLength: number) {
		const response = await processUserUpdates(rubikonLength);
		console.log(response);
		let userExperience: number = response.gainedExperience;
		toast('Du hast ' + userExperience + ' Erfahrungspunkte erhalten.');
		streakLength = response.streakLength;
		if (response != 'You already increased your streak today') {
			dialogOpen = true;
		}
	}
</script>

<Toaster />
<section class="flex items-center justify-center flex-col">
	<TimerComponent
		on:processUser={({ detail: userExperience }) => processUserUpdatesAndShowDialog(userExperience)}
	/>
	<Dialog.Root bind:open={dialogOpen}>
		<Dialog.Content>
			<Dialog.Header>
				<Dialog.Title>Herzlichen Glückwunsch</Dialog.Title>
				<Dialog.Description
					>Du hast {streakLength}
					{streakLength > 1 ? 'Tage' : 'Tag'} produktiv gearbeitet. Mach weiter so.<PastSevenDays
						{streakLength}
					/></Dialog.Description
				>
			</Dialog.Header>
		</Dialog.Content>
	</Dialog.Root>
</section>
