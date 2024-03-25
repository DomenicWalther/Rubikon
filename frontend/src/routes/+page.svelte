<script>
	import { Timer } from '$lib/components';
	import UserButton from 'clerk-sveltekit/client/UserButton.svelte';
	import SignedIn from 'clerk-sveltekit/client/SignedIn.svelte';
	import SignedOut from 'clerk-sveltekit/client/SignedOut.svelte';
	import * as Dialog from '$lib/components/shadcn/ui/dialog';
	let dialogOpen = false;
</script>

<SignedIn>
	<UserButton afterSignOutUrl="/" />
</SignedIn>
<SignedOut>
	<a href="/sign-in">Sign in</a><span>|</span><a href="/sign-up">Sign up</a>
</SignedOut>
<Timer />
