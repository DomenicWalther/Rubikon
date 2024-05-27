<script lang="ts">
    export let label
    export let svgType 
    export let isCurrentRank
    import SvgIcon from '$lib/components/SVGIcon.svelte';
</script>


<SvgIcon type={svgType} color="white" isActiveItem={false} />


<div class="flex flex-col">
        {#if isCurrentRank}
        <h3 class="text-2xl font-medium text-mainlightblue">Aktueller</h3>
        <h3 class="text-2xl font-medium text-mainlightblue">Zustand</h3>
        {/if}
        <div class="w-40">
            <h3 class="text-2xl">{label}</h3>
        </div>
</div>