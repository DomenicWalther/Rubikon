<script>
    	import Premium from '$lib/svg/premium.svg';
	import Einstellungen from '$lib/svg/einstellungen.svg';
	import Shop from '$lib/svg/shop.svg';
	import Gruppen from '$lib/svg/gruppen.svg';
	import Statistiken from '$lib/svg/statistiken.svg';
	import Bewerten from '$lib/svg/bewerten.svg';
	import Teilen from '$lib/svg/teilen.svg';
	import Hilfe from '$lib/svg/hilfe.svg';
	import Über_uns from '$lib/svg/über_uns.svg';
	import Abmelden from '$lib/svg/abmelden.svg';
    import { isSidebarOpen } from '$lib/store';
</script>

<nav class="ham py-1 fixed z-10 bg-white w-full {$isSidebarOpen
    ? 'left-0'
    : '-left-full'} transition-[left] duration-300 ease-in-out" id="navigation">
    <ul class="items-center font-medium">
        <li class="py-5 px-5 bg-gray-300">
            <a href="/" class="font-medium">Rubikon</a>
            <a href="/" class="font-medium"><br />7 Tage Probezeit</a>
        </li>
        <li class="flex items-center mt-7 ml-4 mb-8">
            <a href="/"><div class="rounded bg-red-500 p-1.5 mr-3"><img src={Premium} alt="text" /></div></a>
            <a href="/">Auf Premium upgraden</a>
        </li>
        <li class="flex items-center ml-4 mb-6">
            <a href="/"><div class="rounded bg-mainblue p-1.5 mr-3"><img src={Einstellungen} alt="text" /></div></a>
            <a href="/">Einstellungen</a>
        </li>
        <li class="flex items-center ml-4 mb-6">
            <div class="rounded bg-mainblue p-1.5 mr-3"><a href="/"><img src={Shop} alt="text" /></a></div>
            <a href="/">Shop</a>
        </li>
        <li class="flex items-center ml-4 mb-6">
            <div class="rounded bg-mainblue p-1 mr-3">
                <a href="/"><img src={Gruppen} alt="text" class="w-6 h-6" /></a>
            </div>
            <a href="/">Gruppen</a>
        </li>
        <li class="flex items-center ml-4 mb-6">
            <a href="/"><div class="rounded bg-mainblue p-1.5 mr-3"><img src={Statistiken} alt="text" /></div></a>
            <a href="/">Statistiken</a>
        </li>
        <li class="flex items-center ml-4 mb-6">
            <a href="/"><div class="rounded bg-mainblue p-1.5 mr-3"><img src={Bewerten} alt="text" /></div></a>
            <a href="/">Bewerten Sie die App</a>
        </li>
        <li class="flex items-center ml-4 mb-6">
            <div class="rounded bg-mainblue p-1.5 mr-3"><img src={Teilen} alt="text" /></div>
            <a href="/">Teilen</a>
        </li>
        <li class="flex items-center ml-4 mb-6">
            <div class="rounded bg-mainblue p-1.5 mr-3"><img src={Hilfe} alt="text" /></div>
            <a href="/">Hilfe</a>
        </li>
        <li class="flex items-center ml-4 mb-6">
            <div class="rounded bg-mainblue p-1.5 mr-3"><img src={Über_uns} alt="text" /></div>
            <a href="/">Über uns</a>
        </li>
        <li class="flex items-center px-4 py-20">
            <div class="p-1.5 mr-3"><img src={Abmelden} alt="text" /></div>
            <a href="/">Abmelden</a>
        </li>
    </ul>
</nav>