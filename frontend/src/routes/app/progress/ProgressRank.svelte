<script lang="ts">
    import ProgressRankItems from './ProgressRankItems.svelte';
	export let currentRank: "Rank_1" | "Rank_2" | "Rank_3" |"Rank_4" | "Rank_5" | "Rank_6" | "Rank_7" |"Rank_8" | "Rank_9" | "Rank_10" | "Rank_11" | "Rank_12" | "Rank_13" ;
    import {Ranks} from "./ranks"
    import Lightning from '$lib/Svg/Lightning.svelte';
</script>


{#each Ranks as Rank}
    <div class="flex justify-center items-center gap-10 mb-14">
        <div class="w-40 flex justify-end items-center">
                <div class="bg-mainlightblue w-20 h-3 rounded-md flex items-center relative">
                    <div class="bg-mainlightblue w-5 h-5 rounded-full absolute right-0"></div>
                    <div class="bg-mainlightblue w-3 h-40 rounded absolute left-0 top-0"></div>
                </div>
        </div>
    
        <ProgressRankItems label={Rank.label} svgType={Rank.svgType} isCurrentRank={currentRank === Rank.svgType} />
    </div>
{/each}

