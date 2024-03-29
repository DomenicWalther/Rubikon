<script>
	import { Timer } from '$lib/components';
	import SignedIn from 'clerk-sveltekit/client/SignedIn.svelte';
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
</script>

<SignedIn>
	<UserButton afterSignOutUrl="/" />
</SignedIn>
<Timer />
