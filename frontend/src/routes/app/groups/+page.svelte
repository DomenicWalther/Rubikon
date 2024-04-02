<script lang="ts">
	import { Navigation } from '$lib/components';
	import * as Popover from '$lib/components/shadcn/ui/popover';
	import { processGroupCreation, processJoinGroup } from '$lib/utils/apiHandlers/processGroups.js';
	export let data;
	let groups = data.body;
	console.log(groups);

	let newGroupImageUrl = 'https://via.placeholder.com/150';
	let newGroupName: string;
	let newGroupIsPrivate: boolean;
	let newGroupDescription: string;

	async function createGroup() {
		const group = {
			name: newGroupName,
			description: newGroupDescription,
			isPrivate: newGroupIsPrivate,
			imageURL: newGroupImageUrl
		};
		await processGroupCreation(group);
	}
</script>

<Navigation />
<section class="container">
	<h1 class="font-bold text-4xl text-center">Gruppen</h1>
	<Popover.Root>
		<Popover.Trigger>Gruppe erstellen</Popover.Trigger>
		<Popover.Content class="p-5 w-80 flex gap-5 flex-col rounded-lg shadow-xl">
			<div class="flex gap-20">
				<img src={newGroupImageUrl} alt="Group Picture" class="w-10 h-10 rounded-full" />
				<input
					type="text"
					bind:value={newGroupName}
					placeholder="Name der Gruppe"
					class="w-40 focus:ring-0 focus:outline-none border-b-[1px] border-gray-300"
				/>
			</div>
			<div class="flex justify-between">
				<label for="private">Privat</label>
				<input type="checkbox" id="private" bind:checked={newGroupIsPrivate} name="private" />
			</div>
			<div class="flex flex-col">
				<label for="description">Beschreibung</label>
				<input
					type="textarea"
					class=" h-40 border-[1px] border-gray-300 mt-3"
					bind:value={newGroupDescription}
				/>
			</div>
			<button
				class="text-white font-bold bg-mainorange rounded-full py-2 text-center"
				on:click={() => createGroup()}>Gruppe hinzufügen</button
			>
		</Popover.Content>
	</Popover.Root>

	<div>
		<p class="text-2xl font-bold mt-10">Deine Gruppen</p>
		{#each groups as group}
			{#if group.is_member}
				<div class="flex gap-10 mt-5 w-96 content-around items-center">
					<img src="https://via.placeholder.com/150" alt="Group" class="w-10 h-10 rounded-full" />
					<div>
						<h3>{group.name}</h3>
						<h4>{group.description}</h4>
					</div>
				</div>
			{/if}
		{/each}
		<p class="text-2xl font-bold mt-10">Gruppe beitreten</p>
		{#each groups as group}
			{#if !group.is_member}
				<div class="flex gap-10 mt-5 w-96 content-around items-center">
					<img src="https://via.placeholder.com/150" alt="Group" class="w-10 h-10 rounded-full" />
					<div>
						<h3>{group.name}</h3>
						<h4>{group.description}</h4>
					</div>
					<button
						class="bg-mainorange text-white font-bold px-5 py-2"
						on:click={() => processJoinGroup(group.id)}>Beitreten</button
					>
				</div>
			{/if}
		{/each}
	</div>
</section>
