<script>
	export let label;
	export let svg;
	export let backgroundColor = 'bg-mainblue';
	export let link = '/';
</script>

<li class="flex items-center ml-4">
	<a href="/"
		><div class="rounded p-1.5 mr-3 {backgroundColor}"><img src={svg} alt="text" /></div></a
	>
	<a href={link}>{label}</a>
</li>
