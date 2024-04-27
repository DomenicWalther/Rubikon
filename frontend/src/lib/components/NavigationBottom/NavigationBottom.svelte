<script>
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
	import SvgIcon from '../SVGIcon.svelte';
</script>

<div>
	<nav
		class="bg-mainblue min-w-[400px] max-w-prose text-white font-bold py-2 rounded-full fixed m-auto left-0 right-0"
	>
		<ul class="flex justify-around items-center">
			<UserButton afterSignOutUrl="/" />
			<li class="flex flex-col items-center py-5 rounded-full px-8">
				<a href="/app" class="flex items-center flex-col"><SvgIcon type="Home" size={24} />Home</a>
			</li>
			<li class="py-5 rounded-full px-8">
				<a href="/app/progress" class="flex items-center flex-col"
					><SvgIcon type="Progress" />Fortschritt</a
				>
			</li>
			<li class=" py-5 rounded-full px-8 bg-gray-900">
				<a href="/app/leaderboard" class="flex items-center flex-col"
					><SvgIcon type="Leaderboard" />Rangliste</a
				>
			</li>
		</ul>
	</nav>
</div>
