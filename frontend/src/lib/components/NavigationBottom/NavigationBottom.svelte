<script>
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
	import Home from '$lib/Svg/Home.svelte';
	import Progress from '$lib/Svg/Progress.svelte';
	import Leaderboard from '$lib/Svg/Leaderboard.svelte';
</script>

<div>
	<nav
		class="bg-mainblue min-w-[400px] max-w-prose text-white font-bold py-2 rounded-full fixed m-auto left-0 right-0"
	>
		<ul class="flex justify-around items-center">
			<UserButton afterSignOutUrl="/" />
			<li class="flex flex-col items-center py-5 rounded-full px-8">
				<a href="/app" class="flex items-center flex-col"><Home />Home</a>
			</li>
			<li class="py-5 rounded-full px-8">
				<a href="/app/progress" class="flex items-center flex-col"><Progress />Fortschritt</a>
			</li>
			<li class=" py-5 rounded-full px-8 bg-gray-900">
				<a href="/app/leaderboard" class="flex items-center flex-col"><Leaderboard />Rangliste</a>
			</li>
		</ul>
	</nav>
</div>
