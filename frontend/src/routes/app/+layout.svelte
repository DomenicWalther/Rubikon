<script>
	import Navigation from '$lib/components/Navigation/Navigation.svelte';
	export let data;
</script>

<Navigation userCurrency={data.user.currency} />
<slot />
