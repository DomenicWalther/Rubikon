<script>
	import { Navigation, NavigationBottom } from '$lib/components';
	import Roboter from '$lib/svg/roboter_lobby.svg';
	import Blitz from '$lib/svg/blitz_lobby.svg';
	import Pfeil from '$lib/svg/pfeil.svg';

	export let data;
	let userExperience = data.user.experience;

	let userLevel = Math.floor(Math.sqrt(userExperience / 10) + 1);
	let currentLevelXP = userExperience - Math.pow(userLevel - 1, 2) * 10;
	let nextLevelXP = Math.pow(userLevel, 2) * 10 - Math.pow(userLevel - 1, 2) * 10;
</script>

<Navigation />
<main class="flex items-center flex-col h-fit bg-maingrey">
	<div class="mt-12 mb-10"><h1 class="font-medium text-3xl text-center">Was machen wir <br> Tobias?</h1></div>
	<div class="flex items-center">
		<img src={Pfeil} alt="text" class="w-14 h-8" />
		<img src={Roboter} alt="text" class="w-40 h-70 mb-8 mr-6 ml-6" />
		<img src={Blitz} alt="text" class="w-14" />
	</div>
	<progress
		value={currentLevelXP}
		max={nextLevelXP}
		class="[&::-webkit-progress-bar]:rounded-lg w-80 [&::-webkit-progress-value]:rounded-lg [&::-webkit-progress-bar]:bg-slate-300 [&::-webkit-progress-value]:bg-mainblue [&::-moz-progress-bar]:bg-mainblue"
	></progress>
	<p class="mt-2 font-medium text-base">{currentLevelXP}/{nextLevelXP} Erfahrungspunkte</p>
	<p class="mt-4 font-medium text-xl">Aktuelles Level:{userLevel}</p>
	<button class="bg-mainblue text-white font-medium rounded-full px-10 py-3 m-20 mx-auto">
		<a href="/app/timer" class="text-xl">Reise starten</a>
	</button>
</main>

<NavigationBottom />
