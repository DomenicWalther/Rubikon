<script>
	import { onMount } from 'svelte';
	import Globus from '$lib/svg/globus.svg';
	import Personen from '$lib/svg/people.svg';
	import Linie from '$lib/svg/line.svg';
	import Coin from '$lib/svg/coin.svg';
	import { isSidebarOpen } from '$lib/store';
	import { NavigationSidebar } from '$lib/components';

	export let userCurrency = 0;

	let navigation, sidebar;

	function toggleMenu() {
		$isSidebarOpen = !$isSidebarOpen;
	}

	function handleOutsideClick(event) {
		if (
			navigation &&
			!navigation.contains(event.target) &&
			sidebar &&
			!sidebar.contains(event.target)
		) {
			$isSidebarOpen = false;
		}
	}

	onMount(() => {
		document.addEventListener('click', handleOutsideClick);
		return () => {
			document.removeEventListener('click', handleOutsideClick);
		};
	});
</script>

<nav class="nav-top flex justify-between items-center w-full p-3" bind:this={navigation}>
	<button class="md:hidden hamb {$isSidebarOpen ? 'active' : ''}" on:click={toggleMenu}
		><svg class="ham" viewBox="0 0 100 100"
			><path
				class="line top"
				d="m 30,33 h 40 c 3.722839,0 7.5,3.126468 7.5,8.578427 0,5.451959 -2.727029,8.421573 -7.5,8.421573 h -20"
			></path><path class="line middle" d="m 30,50 h 40"></path><path
				class="line bottom"
				d="m 70,67 h -40 c 0,0 -7.5,-0.802118 -7.5,-8.365747 0,-7.563629 7.5,-8.634253 7.5,-8.634253 h 20"
			></path></svg
		></button
	>
	<div class="flex justify-center items-center">
		<a href="/app/groups"><img src={Personen} alt="Text" /></a>
		<div class="px-2.5">
			<img src={Linie} alt="Text" />
		</div>
		<img src={Globus} alt="Text" />
	</div>
	<div>
		<div class="rounded-full w-24 gap-2 bg-mainblue h-10 flex items-center">
			<img src={Coin} alt="Text" class="h-10" />
			<p class="text-white">{userCurrency}</p>
		</div>
	</div>
</nav>

<NavigationSidebar bind:this={sidebar} />

<style>
	.hamb {
		position: relative;
		margin-right: -0.625rem;
		border-width: 0px;
		background-color: transparent;
		padding: 0;
	}
	.sr-only {
		position: absolute;
		width: 1px;
		height: 1px;
		padding: 0;
		margin: -1px;
		overflow: hidden;
		clip: rect(0, 0, 0, 0);
		white-space: nowrap;
		border-width: 0;
	}
	.hamb .ham {
		-webkit-tap-highlight-color: transparent;
		-webkit-user-select: none;
		-moz-user-select: none;
		user-select: none;
	}
	.hamb .ham {
		height: 60px;
		width: 60px;
		cursor: pointer;
		transition-duration: 0.3s;
	}
	.hamb .ham .top {
		stroke-dasharray: 40 160;
	}
	.hamb .ham .middle {
		transform-origin: 50%;
		stroke-dasharray: 40 142;
	}
	.hamb .ham .bottom {
		transform-origin: 50%;
		stroke-dasharray: 40 85;
	}
	.hamb .line {
		fill: none;
		stroke: black;
		stroke-width: 5;
		transition-duration: 0.3s;
		stroke-linecap: round;
	}
	.hamb.active svg {
		--tw-rotate: 45deg;
		transform: rotate(var(--tw-rotate));
	}
	.hamb.active svg .top {
		stroke-dashoffset: -64px;
	}
	.hamb.active svg .middle {
		--tw-rotate: 90deg;
		transform: rotate(var(--tw-rotate));
	}
	.hamb.active svg .bottom {
		stroke-dashoffset: -64px;
	}
</style>
