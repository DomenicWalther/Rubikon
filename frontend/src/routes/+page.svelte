<script>
	import SignedIn from 'clerk-sveltekit/client/SignedIn.svelte';
	import SignedOut from 'clerk-sveltekit/client/SignedOut.svelte';
</script>

<nav class="bg-white py-6 px-40 flex gap-10 justify-between">
	<h3 class="font-bold text-xl text-mainblue">Rubikon</h3>
	<ol class="flex gap-10">
		<li>Startseite</li>
		<li>Features</li>
		<li>Über Uns</li>
		<li>Preise</li>
		<li>Kontakt</li>
		<li>
			<SignedIn>
				<a href="/app" class="bg-mainorange  text-white font-bold rounded-full px-5 py-2"
					>Dashboard</a
				>
			</SignedIn>
			<SignedOut>
				<a href="/sign-in">Einloggen</a><span>|</span><a href="/sign-up">Registieren</a>
			</SignedOut>
		</li>
	</ol>
</nav>
<main class="flex flex-col items-center">
	<div class="mt-20 gap-40 px-40 flex justify-center items-center">
		<h1 class=" text-7xl font-bold text-mainblue">Organisiere endlich Arbeit und Leben</h1>
		<img src="https://via.placeholder.com/300x450" class="rounded-2xl" alt="placeholder" />
	</div>
	<button
		class="mt-20 text-xl font-bold text-white bg-mainorange px-7 py-4 rounded-full flex items-center justify-center gap-2"
		><span>Jetzt loslegen</span>
		<svg
			class=""
			style="--icon-color:#9c9da1"
			width="16"
			height="16"
			viewBox="0 0 16 16"
			fill="white"
			role="img"
			focusable="false"
			aria-hidden="true"
			><path
				d="M5.46967 11.4697C5.17678 11.7626 5.17678 12.2374 5.46967 12.5303C5.76256 12.8232 6.23744 12.8232 6.53033 12.5303L10.5303 8.53033C10.8207 8.23999 10.8236 7.77014 10.5368 7.47624L6.63419 3.47624C6.34492 3.17976 5.87009 3.17391 5.57361 3.46318C5.27713 3.75244 5.27128 4.22728 5.56054 4.52376L8.94583 7.99351L5.46967 11.4697Z"
			></path></svg
		></button
	>
	<section>
		<h2>Warum Rubikon?</h2>
	</section>
</main>
