<script lang="ts">
	import type { GroupRowProps } from '$lib/types/groups';
	import { leaveGroup } from '$lib/utils/apiHandlers/processGroups.js';
	export let group: GroupRowProps;
</script>



<div class="flex items-center mb-5">
	<a href="/app/groups/{group.id}" class="items-center flex gap-6">
		<img src="https://via.placeholder.com/150" alt="Group" class="w-10 h-10 rounded-md" />
	<div class="flex flex-col min-w-[170px] max-w-[170px] mr-1">
		<h3 class="font-medium truncate">{group.name}</h3>
		<h4 class="truncate">Fangen wir zusammen an{group.description}</h4>
	</div>
	</a>
	<div class=" bg-mainorange w-12 py-2 text-white font-bold rounded-sm flex justify-center mr-4">
		<h3 class="truncate">{group.userCount}</h3>
	</div>
	<button class=" bg-mainorange px-5 py-2 text-white font-bold rounded-sm" on:click={() => leaveGroup(group.id)}>Verlassen</button>
</div>
