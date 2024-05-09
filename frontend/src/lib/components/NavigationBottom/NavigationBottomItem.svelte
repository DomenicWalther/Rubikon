<script lang="ts">
    export let label
    export let link
    export let svgType 
    export let isActive: boolean
    import SvgIcon from '../SVGIcon.svelte';
</script>

<li class="px-8 rounded-full">
    <a href={link} class="flex items-center flex-col"><SvgIcon type={svgType} color="white" isActiveItem={isActive}/>{label}</a>
</li>