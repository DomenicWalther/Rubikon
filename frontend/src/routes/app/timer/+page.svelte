<script>
	import { Timer } from '$lib/components';
	import Roboter from '$lib/svg/roboter_lobby.svg';
</script>

<div
	class="bg-gradient-to-b from-[#232744] from-[54%] to-[#424673] to-[86%] h-screen w-screen flex items-center justify-center flex-col"
>
	<a href="/app" class="  absolute top-1/2 left-20"><div class="arrow left"></div></a>
	<div class="flex flex-col items-center gap-20">
		<p class="text-white text-2xl">Legen wir zusammen los?</p>
		<div>
			<img src={Roboter} alt="text" class="w-60" />
		</div>
	</div>
	<Timer />
</div>

<style>
	.arrow {
		border: solid white;
		border-width: 0 7px 7px 0;
		display: inline-block;
		padding: 7px;
	}

	.left {
		transform: rotate(135deg);
		-webkit-transform: rotate(135deg);
	}
</style>
