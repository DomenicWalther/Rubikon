<script lang="ts">
	import type { GroupRowProps } from '$lib/types/groups';
	import { leaveGroup } from '$lib/utils/apiHandlers/processGroups.js';
	export let group: GroupRowProps;
</script>

<div class="flex w-90 gap-7 mt-5 items-center">
	<a href="/app/groups/{group.id}" class="flex gap-5 items-center">
		<img src="https://via.placeholder.com/150" alt="Group" class="w-10 h-10 rounded-md" />
		<div class="whitespace-nowrap overflow-hidden text-ellipsis">
			<h3 class="font-bold">{group.name}</h3>
			<h4 class="">{group.description}</h4>
		</div>
	</a>
	<div class="bg-mainorange px-5 py-2 text-white font-bold rounded-sm">
		{group.userCount}
	</div>
	<div class="bg-mainorange px-5 py-2 text-white font-bold rounded-sm">
		<button on:click={() => leaveGroup(group.id)}>Verlassen</button>
	</div>
</div>
