<script>
	import { Timer } from '$lib/components';
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
	import SignedIn from 'clerk-sveltekit/client/SignedIn.svelte';
	import SignedOut from 'clerk-sveltekit/client/SignedOut.svelte';
	import { onMount } from 'svelte';

	onMount(async () => {
		fetch('http://localhost:3000/foo', {
			method: 'POST',
			headers: {
				Authorization: `Bearer ${await Clerk.session.getToken()}`
			}
		}).then((res) => {
			if (res.ok) {
				console.log(res.text());
			} else {
				console.log('Failure');
			}
		});
	});
</script>

<SignedIn>
	<UserButton afterSignOutUrl="/" />
</SignedIn>
<SignedOut>
	<a href="/sign-in">Sign in</a><span>|</span><a href="/sign-up">Sign up</a>
</SignedOut>
<Timer />
