<script>
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
</script>

<div>
	<nav
		class="bg-mainblue min-w-[400px] max-w-prose text-white font-bold py-2 rounded-full fixed bottom-10 m-auto left-0 right-0"
	>
		<ul class="flex justify-around items-center">
			<UserButton afterSignOutUrl="/" />
			<li class=" py-5 rounded-full px-8"><a href="/app">Home</a></li>
			<li class="py-5 rounded-full px-8"><a href="/app/progress">Fortschritt</a></li>
			<li class="py-5 rounded-full px-8 bg-gray-900"><a href="/app/leaderboard">Rangliste</a></li>
		</ul>
	</nav>
</div>
