<script>
	import { onMount } from 'svelte';
	import Globus from '$lib/svg/globus.svg';
	import Personen from '$lib/svg/people.svg';
	import Linie from '$lib/svg/line.svg';
	import Ham from '$lib/svg/ham.svg';
	import Coin from '$lib/svg/coin.svg';
	import { isSidebarOpen } from '$lib/store';
	import { NavigationSidebar } from '$lib/components';

	$isSidebarOpen;

	function toggleMenu() {
		$isSidebarOpen = !$isSidebarOpen;
	}

	function closeSidebarOnClickOutside(event) {
		const navigation = document.getElementById('navigation');
		const sidebar = document.getElementById('sidebar');
		if (
			navigation &&
			!navigation.contains(event.target) &&
			sidebar &&
			!sidebar.contains(event.target)
		) {
			$isSidebarOpen = false;
		}
	}

	onMount(async () => {
		document.addEventListener('click', closeSidebarOnClickOutside);
	});
</script>

<nav class="nav-top flex justify-between items-center w-full p-3" id="navigation">
	<button class="md:hidden" on:click={toggleMenu}><img src={Ham} alt="Text" /></button>
	<div class="mitte flex justify-center items-center">
		<img src={Personen} alt="Text" />
		<div class="px-2.5">
			<img src={Linie} alt="Text" />
		</div>
		<img src={Globus} alt="Text" />
	</div>
	<div>
		<div class="rounded w-10 to-blue-600 h-6"></div>
		<img src={Coin} alt="Text" />
	</div>
</nav>

<NavigationSidebar />
