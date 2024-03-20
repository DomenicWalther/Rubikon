<script lang="ts">
	import SignedIn from 'clerk-sveltekit/client/SignedIn.svelte';
	import UserProfile from 'clerk-sveltekit/client/UserProfile.svelte';
</script>

<SignedIn let:user>
	<p>Welcome, {user?.fullName ?? 'buddy'}!</p>
	<p>You are in the admin panel.</p>
	<UserProfile />
</SignedIn>
