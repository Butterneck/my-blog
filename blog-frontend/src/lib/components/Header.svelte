<script>
	import NavMenu from './NavMenu.svelte';
</script>

<header>
	<nav class="flex items-center justify-center max-width-main z-50 h-16 my-6 sm:mx-12 mx-6 xl:mx-auto">
		<div class="flex-grow flex items-center">
			<a href="/">
				<div class="pr-4 nav-img h-8 flex flex-row items-center cursor-pointer border-0">
					<img class="h-6 sm:h-full" src="https://reddium.vercel.app/reddium_symbol.svg" alt=""/>
					<h1 class="ml-4 site-name text-3xl tracking-tighter block text-black">Butterneck</h1>
				</div>
			</a>
		</div>
		<NavMenu />
	</nav>
</header>
