<script lang="ts">
	import WidePost from './WidePost.svelte';
	export let posts: Post[];
</script>

<div
	class="md:w-auto main-container max-width-main pb-10 pt-10 sub-top-border sm:mx-12 mx-6 xl:mx-auto posts-grid"
>
	<div class="w-full mb-4 grid-left lg:pl-8">
		{#each posts as post}
			<WidePost {post} />
		{/each}
		<div class="w-full text-center">
			<!-- {/* <WideCard {...({} as Post)} /> */} -->
			<button
				class="my-4 mx-auto p-2 cursor-pointer w-48 max-w-full load-more main-black font-semibold rounded flex flex-row justify-between items-center"
			>
				<div class="flex-grow text-center">Show More</div>
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					class="w-6 h-6"
				>
					<path stroke-linecap="round" stroke-linejoin="round" d="M19.5 8.25l-7.5 7.5-7.5-7.5" />
				</svg>
			</button>
		</div>
	</div>
	<div class="grid-right hidden md:block">
		<div class="sticky top-8 p-8 about-bg flex flex-col">
			<div class="w-full flex mb-4 flex-row items-center">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					stroke-width="1.5"
					stroke="currentColor"
					class="w-6 h-6"
				>
					<path
						stroke-linecap="round"
						stroke-linejoin="round"
						d="M17.593 3.322c1.1.128 1.907 1.077 1.907 2.185V21L12 17.25 4.5 21V5.507c0-1.108.806-2.057 1.907-2.185a48.507 48.507 0 0111.186 0z"
					/>
				</svg>

				<div>
					<p class="heading-text text-sm leading-4 uppercase tracking-wide sm:text-xs">
						About Reddium
					</p>
				</div>
			</div>
			<div class="w-full pb-6">
				<p class="text-sm">
					Ever wanted to browse Reddit while studying at Starbucks? Or while sitting on the subway
					to work? Worried that people around you would judge the subreddits you browse and the
					posts you read?
					<br />
					<br />
					{`Reddium is a Medium-themed Reddit client. The Reddium interface
			converts Reddit posts, discussions, and memes into well-crafted
			articles. Medium's layout feels a little more readable than
			Reddit's, removing all distractions and clutter. It also
			bypasses Reddit's frustrating mobile browser.`}
					<br />
					<br />I hope you enjoy this project! Feel free to suggest any features or report bugs on
					GitHub.
				</p>
			</div>
			<div class="w-full pb-6 hidden">
				<img class="w-4/12 float-right" src="/signature.png" />
			</div>
			<a href="https://github.com/eightants/reddium/" target="_blank" rel="noopener noreferrer">
				<button
					class="mt-2 mx-1 p-2 pl-0 pb-3 cursor-pointer w-full max-w-full btn-black text-white rounded"
				>
					✨ Star on GitHub
				</button>
			</a>
			<a href="https://ko-fi.com/eightants" target="_blank" rel="noopener noreferrer">
				<button
					class="mt-2 mx-1 p-2 pl-0 pb-3 cursor-pointer w-full max-w-full btn-outline-black text-white rounded"
				>
					☕ Buy me a coffee
				</button>
			</a>
		</div>
	</div>
</div>
