<script>
	import { Navigation, NavigationBottom} from '$lib/components';
	import Roboter from '$lib/svg/roboter_lobby.svg';
	import Blitz from '$lib/svg/blitz_lobby.svg';

	export let data;
	let userExperience = data.user.experience;

	let userLevel = Math.floor(Math.sqrt(userExperience / 10) + 1);
	let currentLevelXP = userExperience - Math.pow(userLevel - 1, 2) * 10;
	let nextLevelXP = Math.pow(userLevel, 2) * 10 - Math.pow(userLevel - 1, 2) * 10;
</script>

<Navigation />
<NavigationBottom/>
<main class="flex items-center flex-col pt-10 h-screen">
	<div class="mt-10 mb-10"><h1 class="font-bold text-2xl">Was machen wir Tobias?</h1></div>
	<div class="flex items-center">
		<img src={Roboter} alt="text" class="w-80 h-80 mb-8" />
		<img src={Blitz} alt="text" class="w-15 h-15" />
	</div>
		<div class="flex items-center h-5 w-80 bg-mainblue rounded-full m-5"></div>
		<p class="font-medium text-xl">{currentLevelXP}/{nextLevelXP} Erfahrungspunkte</p>
		<p class="font-medium text-xl">Aktuelles Level:{userLevel}</p>
	<a href="/app/timer" class="mt-16 bg-mainblue text-white font-bold rounded-full px-5 py-3"
		>Reise starten</a
	>
</main>
