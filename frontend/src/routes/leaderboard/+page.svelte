<script lang="ts">
	export let data;
	let users = data.body;
	console.log(users);
</script>

<a href="/">Home</a>
<section>
	<h1 class="font-bold text-3xl text-center">Ranking</h1>
	{#each users as user}
		<h3>{user.ID}</h3>
		<h4>{user.experience}</h4>
	{/each}
</section>
