<script lang="ts">
	import RobotBase from './robot_base.svelte';
	import SkinImport from '$lib/Svg/Skins/SkinImport.svelte';
	import { activeSkin } from '$lib/store';
</script>

<div class="relative h-full w-full mt-14 mx-6">
	<SkinImport type={$activeSkin.activeHats} skin_type="hats" />
	<RobotBase />
	<SkinImport type={$activeSkin.activeEyes} skin_type="eyes" />
	<SkinImport type={$activeSkin.activeMouths} skin_type="mouths" />
</div>
