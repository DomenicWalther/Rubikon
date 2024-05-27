<script lang="ts">
    import ProgressRankItems from './ProgressRankItems.svelte';
	let currentRank: "Rank_1" | "Rank_2" | "Rank_3" |"Rank_4" | "Rank_5" | "Rank_6" | "Rank_7" |"Rank_8" | "Rank_9" | "Rank_10" | "Rank_11" | "Rank_12" | "Rank_13" ;
    import {Ranks} from "./ranks"
    export let userExperience
    import Lightning from '$lib/Svg/Lightning.svelte';
    
	let newWidth = 35;
	let newHeight = 35;
</script>


{#each Ranks as Rank, index}
    <div class="flex justify-center items-center gap-8 mb-14">
        <div class="w-40 flex justify-center items-center">
            
            <div class="w-20 flex justify-center items-center relative">
                {#if Rank.svgType !== Ranks[Ranks.length -1].svgType}
                <div class="flex flex-col items-center absolute top-6"><Lightning width={newWidth} height={newHeight}/>
                    <p class="mt-1 font-medium text-base">{Rank.required_xp}-{Ranks[index+1].required_xp}</p>
                </div>
                {/if}
            </div>

                <div class="w-20">
                    <div class="bg-mainlightblue w-20 h-2 rounded-md flex items-center relative">
                        <div class="bg-mainlightblue w-5 h-5 rounded-full absolute right-0"></div>
                        {#if Rank.svgType !== Ranks[Ranks.length -1].svgType}
                        <div class="bg-mainlightblue w-2 h-[150px] rounded absolute left-0 top-0"></div>
                        {/if}
                    </div>
                </div>
        </div>
        <ProgressRankItems label={Rank.label} svgType={Rank.svgType} isCurrentRank={userExperience >= Rank.required_xp && userExperience < Ranks[index+1].required_xp} />
    </div>
{/each}

