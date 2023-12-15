<script lang="ts">
	import SmallPost from './SmallPost.svelte';
	import LoadMorePostsBtn from './LoadMorePostsBtn.svelte';
	import type { Post } from '$lib/generated/backend-client';
	export let posts: Post[];
</script>

<div class="home-container posts-grid">
	<div class="w-full mb-4 grid-left lg:pl-8">
		<div class="w-full mb-4 flex flex-row items-center">
			<span class="mr-2">
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
						d="M20.25 7.5l-.625 10.632a2.25 2.25 0 01-2.247 2.118H6.622a2.25 2.25 0 01-2.247-2.118L3.75 7.5M10 11.25h4M3.375 7.5h17.25c.621 0 1.125-.504 1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125z"
					/>
				</svg>
			</span>

			<div>
				<p class="heading-text text-sm leading-4 uppercase tracking-wide sm:text-xs">All posts</p>
			</div>
		</div>
		{#each posts as post}
			<SmallPost {post} />
		{/each}
		<LoadMorePostsBtn />
	</div>
	<div class="grid-right hidden md:block">
		<div class="sticky top-8 p-8 about-bg flex flex-col mr-8">
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
