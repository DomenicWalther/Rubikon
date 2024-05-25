<script lang="ts">
	import RobotBase from './robot_base.svelte';
	import SkinImport from '$lib/Svg/Skins/SkinImport.svelte';
	let activeHat = 'robot_hat_01';
	let activeMouth = 'robot_mouth_01';
	let activeEye = 'robot_eyes_01';
	const hats = ['robot_hat_01', 'robot_hat_02', 'robot_hat_03'];
	const eyes = ['robot_eyes_01', 'robot_eyes_02', 'robot_eyes_03'];
	const mouths = ['robot_mouth_01', 'robot_mouth_02'];
</script>

<h1>Skin Playground</h1>

<div class="relative bg-pink-500 h-full w-full mt-14">
	<SkinImport type={activeHat} skin_type="hats" />
	<RobotBase />
	<SkinImport type={activeEye} skin_type="eyes" />
	<SkinImport type={activeMouth} skin_type="mouths" />
</div>

<h3 class="text-4xl font-bold">Available Skins</h3>
<h4 class="text-2xl font-bold">Hats</h4>
<div class="flex flex-col gap-5">
	<div class="relative flex gap-20">
		{#each hats as hat}
			<button
				on:click={() => (activeHat = hat)}
				class="w-20 h-20 bg-gray-300 justify-center items-center flex rounded-xl p-10"
			>
				<SkinImport type={hat} skin_type="hats" isShowcase={true} />
			</button>
		{/each}
	</div>
	<div class="relative flex gap-20">
		{#each eyes as eye}
			<button
				on:click={() => (activeEye = eye)}
				class="w-21 h-20 bg-gray-300 justify-center items-center flex rounded-xl p-10"
			>
				<SkinImport type={eye} skin_type="eyes" isShowcase={true} />
			</button>
		{/each}
	</div>
	<div class="relative flex gap-20">
		{#each mouths as mouth}
			<button
				on:click={() => (activeMouth = mouth)}
				class="w-23 h-20 bg-gray-300 justify-center items-center flex rounded-xl p-10"
			>
				<div class="scale-[300%]">
					<SkinImport type={mouth} skin_type="mouths" isShowcase={true} />
				</div>
			</button>
		{/each}
	</div>
</div>
