<script lang="ts">
	import type { GroupRowProps } from '$lib/types/groups';
	import { leaveGroup } from '$lib/utils/apiHandlers/processGroups.js';
	export let group: GroupRowProps;
</script>

<div class="flex gap-10 mt-5 w-80 justify-between items-center">
	<div class="flex gap-6">
		<img src="https://via.placeholder.com/150" alt="Group" class="w-10 h-10 rounded-full" />
		<div>
			<h3>{group.name}</h3>
			<h4>{group.description}</h4>
		</div>
	</div>
	<div class="bg-mainorange px-5 py-2 text-white font-bold">
		{group.userCount}
	</div>
	<div class="bg-mainorange px-5 py-2 text-white font-bold">
		<button on:click={() => leaveGroup(group.id)}>Verlassen</button>
	</div>
</div>
