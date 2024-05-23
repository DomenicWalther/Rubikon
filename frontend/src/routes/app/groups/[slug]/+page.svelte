<script lang="ts">
	export let data;
	import { createNewGroupMessage } from '$lib/utils/apiHandlers/processGroupMessages';
	import { currentGroupChatMessages } from '$lib/store.js';
	import Pusher from 'pusher-js';

	let newMessage: string;

	let pusher = new Pusher('4b5ad963b27dd5a727d3', {
		cluster: 'eu'
	});

	const channelID = 'group-chat-' + data.groupID;
	var channel = pusher.subscribe(channelID);
	channel.bind('new-message', function (data) {
		$currentGroupChatMessages = [
			...$currentGroupChatMessages,
			{
				message: data.message,
				group_id: data.group_id
			}
		];
	});
	$currentGroupChatMessages = data.messages;
	const submitMessage = () => {
		createNewGroupMessage({
			message: newMessage,
			group_id: data.groupID,
			channelID
		});
		newMessage = '';
	};
</script>

<section class="bg-gray-300 h-screen mt-0">
	<div class="bg-white flex items-center mx-10 p-10 gap-5">
		<img src="https://picsum.photos/200" alt="Group" class="w-20 rounded-lg h-20" />
		<div class="flex flex-col gap-2">
			<h2 class="font-semibold text-xl">Ziel: 20h in der Woche</h2>
			<span class="text-gray-400 text-lg">Fangen wir zusammen an</span>
			<div class="flex gap-5">
				<span class="bg-mainorange px-2 py-1 text-sm text-white font-bold rounded-sm text-center"
					>1328</span
				>
				<button class="px-2 py-1 bg-gray-200">Verlassen</button>
			</div>
		</div>
	</div>

	<div class="bg-white mt-10 pt-2">
		<div class=" flex justify-around mx-60">
			<p class="text-2xl w-[250px] text-gray-400">Livestream</p>
			<p class="text-2xl border-b-2 border-mainorange w-[250px] text-center">Nachrichten</p>
		</div>
	</div>
	<div class="mx-40 flex flex-col gap-5 mt-5 overflow-y-auto flex-col-reverse h-[600px]">
		<div>
			{#each $currentGroupChatMessages as message}
				<div class="flex items-center gap-5">
					<img src="https://picsum.photos/50" alt="Group" class="w-10 rounded-full h-10" />
					<p class="text-xl">{message.message}</p>
				</div>
			{/each}
		</div>
	</div>
	<div class="fixed bottom-0 w-screen flex justify-center bg-white py-5">
		<input
			bind:value={newMessage}
			type="text"
			placeholder="Schreibe eine Nachricht"
			class="bg-gray-200 w-80 py-4 px-8 rounded-3xl text-black placeholder-black placeholder:font-semibold"
		/>
		<button on:click={() => submitMessage()}>Send Message</button>
	</div>
</section>
