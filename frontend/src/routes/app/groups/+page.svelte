<script lang="ts">
	import { Navigation } from '$lib/components';
	import * as Popover from '$lib/components/shadcn/ui/popover';
</script>

<Navigation />
<section class="container">
	<h1 class="font-bold text-4xl text-center">Gruppen</h1>
	<Popover.Root>
		<Popover.Trigger>Gruppe erstellen</Popover.Trigger>
		<Popover.Content class="p-5 w-80 flex gap-5 flex-col rounded-lg shadow-xl">
			<div class="flex gap-20">
				<img
					src="https://via.placeholder.com/150"
					alt="Group Picture"
					class="w-10 h-10 rounded-full"
				/>
				<input
					type="text"
					placeholder="Name der Gruppe"
					class="w-40 focus:ring-0 focus:outline-none border-b-[1px] border-gray-300"
				/>
			</div>
			<div class="flex justify-between">
				<label for="private">Privat</label>
				<input type="checkbox" id="private" name="private" value="private" />
			</div>
			<div class="flex flex-col">
				<label for="description">Beschreibung</label>
				<input type="textarea" class=" h-40 border-[1px] border-gray-300 mt-3" />
			</div>
			<button class="text-white font-bold bg-mainorange rounded-full py-2 text-center"
				>Gruppe hinzufügen</button
			>
		</Popover.Content>
	</Popover.Root>
</section>
