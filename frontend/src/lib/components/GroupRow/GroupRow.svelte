<script lang="ts">
	export let group;
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	async function leaveGroup(groupId: number) {
		const response = await fetch(`${PUBLIC_BACKEND_URL}/Groups/leave`, {
			method: 'DELETE',
			headers: {
				'Content-Type': 'application/json',
				Authorization: `Bearer ${await Clerk.session.getToken()}`
			},
			body: JSON.stringify({ group_id: groupId })
		});
		data = await response.json();
		console.log(data);
	}
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
