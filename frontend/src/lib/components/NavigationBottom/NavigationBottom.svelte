<script lang="ts">
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
	import NavigationBottomItem from './NavigationBottomItem.svelte';
	export let activeItem: "leaderboard" | "progress" | "home" = "leaderboard"
</script>


	<nav class="bg-mainblue min-w-[400px] max-w-prose text-white font-bold py-2 fixed m-auto left-0 right-0 bottom-0">
		<ul class="flex justify-around items-center py-5 relativ">
			<NavigationBottomItem label="Home" link="/app" svgType="Home" isActive={"home" === activeItem} />

			<NavigationBottomItem label="Fortschritt" link="/app/progress" svgType="Progress" isActive={"progress" === activeItem} />

			<NavigationBottomItem label="Rangliste" link="/app/leaderboard" svgType="Leaderboard" isActive={"leaderboard" === activeItem} />


		</ul>
	</nav>


