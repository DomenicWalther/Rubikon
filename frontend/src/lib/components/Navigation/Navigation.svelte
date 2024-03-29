<script>
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
</script>

<nav
	class="bg-mainblue min-w-[400px] max-w-prose text-white font-bold py-5 rounded-full absolute bottom-10 m-auto left-0 right-0"
>
	<ul class="flex justify-around items-center">
		<UserButton afterSignOutUrl="/" />
		<li><a href="/">Home</a></li>
		<li><a href="/">Fortschritt</a></li>
		<li><a href="/app/leaderboard">Rangliste</a></li>
	</ul>
</nav>
