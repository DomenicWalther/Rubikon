<script lang="ts">
	import { GroupRow } from '$lib/components';
	import { currentGroups } from '$lib/store';
	import * as Popover from '$lib/components/shadcn/ui/popover';
	import { processGroupCreation, processJoinGroup } from '$lib/utils/apiHandlers/processGroups.js';
	export let data;
	$currentGroups = data.body;
	console.log(data.body);

	let newGroupImageUrl = 'https://via.placeholder.com/150';
	let newGroupName: string;
	let newGroupIsPrivate: boolean;
	let newGroupDescription: string;
	const hasMember = $currentGroups.some((group) => group.is_member);

	async function createGroup() {
		const group = {
			name: newGroupName,
			description: newGroupDescription,
			isPrivate: newGroupIsPrivate,
			imageURL: newGroupImageUrl
		};
		await processGroupCreation(group);
		$currentGroups = [...$currentGroups, group];
	}
</script>

<section class="container">
	<h1 class="font-medium text-3xl text-center">Gruppen</h1>
	<Popover.Root>
		<Popover.Trigger>Gruppe erstellen</Popover.Trigger>
		<Popover.Content class="p-5 w-80 flex gap-5 flex-col rounded-lg shadow-xl">
			<div class="flex gap-7">
				<div class="rounded-sm border-2 border-dashed border-mainorange w-10 h-10">
					<img src={newGroupImageUrl} alt="Group Picture" class="w-10 h-10 invisible" />
				</div>
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

		<p class="text-2xl font-medium mt-10 mb-4">Meine Gruppen</p>
		{#if hasMember}
			{#each $currentGroups as group}
				{#if group.is_member}
					<GroupRow {group} />
				{/if}
			{/each}
		{:else}
			<p>Noch keine Gruppe erstellt oder beigetreten</p>
		{/if}

		<p class="text-2xl font-medium mt-10">Weitere Gruppen</p>
		<input type="text" placeholder="Suchen" class="mb-1 font-medium" />

		{#each $currentGroups as group}
		{#if !group.is_member}
			<div class="flex items-center mb-6 mt-1">
				<div class="flex items-center gap-6">
					<img src="https://via.placeholder.com/150" alt="Group" class="w-10 h-10 rounded-md" />
					<div class="flex flex-col min-w-[170px] max-w-[170px] mr-1">
						<h3 class="font-medium truncate">{group.name}</h3>
						<h4 class="truncate">{group.description}</h4>
					</div>
				</div>
				<div class=" bg-mainorange w-12 py-2 text-white font-bold rounded-sm flex justify-center mr-4">
					<h3 class="truncate">{group.userCount}</h3>
				</div>
					<button class="bg-mainorange text-white font-bold px-5 py-2 rounded-sm flex"
					on:click={() => processJoinGroup(group.id)}>Beitreten</button>
			</div>
		{/if}
		{/each}
</section>


