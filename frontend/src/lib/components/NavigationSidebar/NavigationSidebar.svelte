<script lang="ts">
	import Einstellungen from '$lib/svg/einstellungen.svg';
	import Shop from '$lib/svg/shop.svg';
	import Gruppen from '$lib/svg/gruppen.svg';
	import Statistiken from '$lib/svg/statistiken.svg';
	import Bewerten from '$lib/svg/bewerten.svg';
	import Teilen from '$lib/svg/teilen.svg';
	import Hilfe from '$lib/svg/hilfe.svg';
	import Über_uns from '$lib/svg/über_uns.svg';
	import Abmelden from '$lib/svg/abmelden.svg';
	import Premium from '$lib/svg/premium.svg';
	import { isSidebarOpen } from '$lib/store';
	import SidebarItem from './SidebarItem.svelte';
	import SignOutButton from 'clerk-sveltekit/client/SignOutButton.svelte';
	import { goto } from '$app/navigation';

	const sidebarItems: Array<any> = [
		{ label: 'Auf Premium upgraden', svg: Premium, backgroundColor: 'bg-red-500' },
		{ label: 'Einstellungen', svg: Einstellungen },
		{ label: 'Shop', svg: Shop },
		{ label: 'Gruppen', svg: Gruppen },
		{ label: 'Statistiken', svg: Statistiken },
		{ label: 'Bewerten Sie die App', svg: Bewerten },
		{ label: 'Teilen', svg: Teilen },
		{ label: 'Hilfe', svg: Hilfe },
		{ label: 'Über uns', svg: Über_uns }
	];
	const handleSignOut = () => {
		goto('/');
	};
</script>

<nav
	class="ham py-1 fixed z-10 bg-white w-full h-full {$isSidebarOpen
		? 'left-0'
		: '-left-full'} transition-[left] duration-300 ease-in-out"
	id="sidebar"
>
	<ul class="items-center font-medium">
		<li class="py-5 px-5 bg-gray-300">
			<a href="/" class="font-medium">Rubikon</a>
			<a href="/" class="font-medium"><br />7 Tage Probezeit</a>
		</li>
		<div class="flex gap-6 flex-col mt-7">
			{#each sidebarItems as { label, svg, backgroundColor }}
				<SidebarItem {label} {svg} {backgroundColor} />
			{/each}
		</div>
		<li class="flex items-center px-4 py-20">
			<div class="p-1.5 mr-3"><img src={Abmelden} alt="text" /></div>
			<SignOutButton signOutCallback={handleSignOut}><button>Abmelden</button></SignOutButton>
		</li>
	</ul>
</nav>
