<script>
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
	import SignedIn from 'clerk-sveltekit/client/SignedIn.svelte';

	export let data;
	let userLevel = data.user.experience;
	console.log(data.user);
</script>

<SignedIn>
	<UserButton afterSignOutUrl="/" />
</SignedIn>

<main class="flex items-center flex-col mt-10">
	<h2>Was machen wir Tobias?</h2>
	<div class="flex gap-5">
		<div class="h-40 w-20 bg-mainblue"></div>
		<p>{userLevel} XP</p>
	</div>
	<a href="/app/timer" class="mt-15 bg-mainblue text-white font-bold rounded-lg px-5 py-3 mt-10"
		>Reise starten</a
	>
</main>
