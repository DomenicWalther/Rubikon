<script>
	import { onMount } from 'svelte';
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
	import Globus from '$lib/svg/globus.svg';
	import People from '$lib/svg/people.svg';
	import Line from '$lib/svg/line.svg';
	import Ham from '$lib/svg/ham.svg';
	import Coin from '$lib/svg/coin.svg';
	import Premium from '$lib/svg/premium.svg';
	import Einstellungen from '$lib/svg/einstellungen.svg';
	import Shop from '$lib/svg/shop.svg';
	import Gruppen from '$lib/svg/gruppen.svg';
	import Statistiken from '$lib/svg/statistiken.svg';
	import Bewerten from '$lib/svg/bewerten.svg';
	import Teilen from '$lib/svg/teilen.svg';
	import Hilfe from '$lib/svg/hilfe.svg';
	import Über_uns from '$lib/svg/über_uns.svg';
	import Abmelden from '$lib/svg/abmelden.svg';

	let isMenuOpen = false;

	function toggleMenu() {
		isMenuOpen = !isMenuOpen;
	}

	function closeMenüOnClickOutside(event) {
		if (!document.getElementById('navigation')?.contains(event.target)) {
			isMenuOpen = false;
		}
	}

	onMount(async () => {
		document.addEventListener('click', closeMenüOnClickOutside);
	});
</script>

<nav class="nav-top flex justify-between items-center w-full p-3" id="navigation">
	<button class="md:hidden" on:click={toggleMenu}><img src={Ham} alt="Text" /></button>
	<div class="mitte flex justify-center items-center">
		<img src={People} alt="Text" />
		<div class="px-2.5">
			<img src={Line} alt="Text" />
		</div>
		<img src={Globus} alt="Text" />
	</div>
	<div>
		<img src={Coin} alt="Text" />
	</div>
</nav>

{#if isMenuOpen}
	<nav class="ham py-1" id="navigation">
		<ul class="items-center font-medium">
			<li class="py-5 px-5 bg-gray-300">
				<a href="/" class="font-medium">Rubikon</a>
				<a href="/" class="font-medium"><br />7 Tage Probezeit</a>
			</li>
			<li class="flex items-center mt-7 ml-4 mb-8">
				<div class="rounded bg-red-500 p-1.5 mr-3"><img src={Premium} alt="text" /></div>
				<a href="/">Auf Premium upgraden</a>
			</li>
			<li class="flex items-center ml-4 mb-6">
				<div class="rounded bg-mainblue p-1.5 mr-3"><img src={Einstellungen} alt="text" /></div>
				<a href="/">Einstellungen</a>
			</li>
			<li class="flex items-center ml-4 mb-6">
				<div class="rounded bg-mainblue p-1.5 mr-3"><img src={Shop} alt="text" /></div>
				<a href="/">Shop</a>
			</li>
			<li class="flex items-center ml-4 mb-6">
				<div class="rounded bg-mainblue p-1 mr-3">
					<img src={Gruppen} alt="text" class="w-6 h-6" />
				</div>
				<a href="/">Gruppen</a>
			</li>
			<li class="flex items-center ml-4 mb-6">
				<div class="rounded bg-mainblue p-1.5 mr-3"><img src={Statistiken} alt="text" /></div>
				<a href="/">Statistiken</a>
			</li>
			<li class="flex items-center ml-4 mb-6">
				<div class="rounded bg-mainblue p-1.5 mr-3"><img src={Bewerten} alt="text" /></div>
				<a href="/">Bewerten Sie die App</a>
			</li>
			<li class="flex items-center ml-4 mb-6">
				<div class="rounded bg-mainblue p-1.5 mr-3"><img src={Teilen} alt="text" /></div>
				<a href="/">Teilen</a>
			</li>
			<li class="flex items-center ml-4 mb-6">
				<div class="rounded bg-mainblue p-1.5 mr-3"><img src={Hilfe} alt="text" /></div>
				<a href="/">Hilfe</a>
			</li>
			<li class="flex items-center ml-4 mb-6">
				<div class="rounded bg-mainblue p-1.5 mr-3"><img src={Über_uns} alt="text" /></div>
				<a href="/">Über uns</a>
			</li>
			<li class="flex items-center px-4 py-20">
				<div class="p-1.5 mr-3"><img src={Abmelden} alt="text" /></div>
				<a href="/">Abmelden</a>
			</li>
		</ul>
	</nav>
{/if}

<div class="bottom_navigatiion">
	<nav
		class="bg-mainblue min-w-[400px] max-w-prose text-white font-bold py-2 rounded-full relativ bottom-10 m-auto left-0 right-0"
	>
		<ul class="flex justify-around items-center">
			<UserButton afterSignOutUrl="/" />
			<li class=" py-5 rounded-full px-8"><a href="/app">Home</a></li>
			<li class="py-5 rounded-full px-8"><a href="/app/progress">Fortschritt</a></li>
			<li class="py-5 rounded-full px-8 bg-gray-900"><a href="/app/leaderboard">Rangliste</a></li>
		</ul>
	</nav>
</div>
