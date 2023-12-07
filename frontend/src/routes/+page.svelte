<script lang="ts">
	export let data;
	import { envVariables } from '$lib/envVariables';
	import { blogMetaData } from '$lib/blogMetaData';
	import { MetaTags } from 'svelte-meta-tags';
	import PostList from '$lib/components/PostList.svelte';
	import BigContainer from '$lib/components/BigContainer.svelte';
	import RankedPost from '$lib/components/RankedPost.svelte';
	import WidePost from '$lib/components/WidePost.svelte';
	import SmallPost from '$lib/components/SmallPost.svelte';

	const meta = {
		title: `Home | ${blogMetaData.blogTitle}`,
		description: blogMetaData.description,
		url: envVariables.basePath,
		siteName: blogMetaData.blogTitle,
		image: {
			url: `${envVariables.basePath}/background.jpeg`,
			width: 1000,
			height: 523,
			alt: 'image'
		}
	};
</script>

<MetaTags
	title={meta.title}
	description={meta.description}
	canonical={meta.url}
	openGraph={{
		description: meta.description,
		images: [
			{
				...meta.image
			}
		],
		siteName: meta.siteName,
		title: meta.title,
		type: 'website',
		url: meta.url
	}}
	twitter={{
		cardType: 'summary_large_image',
		title: meta.title,
		description: meta.description,
		image: meta.image.url,
		imageAlt: meta.image.alt
	}}
/>

<!-- LATEST POSTS -->
<BigContainer>
	<PostList posts={data.posts} />
</BigContainer>

<!-- TRENDING POSTS -->
<div
	class="flex md:w-auto main-container max-width-main pb-10 pt-10 sub-top-border sm:mx-12 mx-6 xl:mx-auto"
>
	<div class="w-full flex mb-4 flex-row items-center lg:px-8">
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
					d="M2.25 18L9 11.25l4.306 4.307a11.95 11.95 0 015.814-5.519l2.74-1.22m0 0l-5.94-2.28m5.94 2.28l-2.28 5.941"
				/>
			</svg>
		</span>

		<div>
			<p class="heading-text text-sm leading-4 uppercase tracking-wide sm:text-xs">
				Trending posts
			</p>
		</div>
	</div>
	<div class="w-full flex mb-4 flex-row items-start flex-wrap lg:px-8">
		{#each data.posts.slice(2, 6) as post}
			<RankedPost {post} />
		{/each}
	</div>
</div>


<!-- ALL POSTS -->
<!-- <div class="w-full flex main-container max-width-main pb-4 pt-10 sub-top-border posts-grid lg:w-auto lg:mx-12 md:block sm:mx-6"> -->

<div class="flex md:w-auto main-container max-width-main pb-10 pt-10 sub-top-border sm:mx-12 mx-6 xl:mx-auto posts-grid">	
	<div class="w-full mb-4 grid-left">
		{#each data.posts.slice(1, 10) as post}

			<SmallPost {post} />
		{/each}
	  <div class="w-full text-center">
		<!-- {/* <WideCard {...({} as Post)} /> */} -->
		<button
		  class="my-4 mx-auto p-2 cursor-pointer w-48 max-w-full load-more main-black font-semibold rounded flex flex-row justify-between items-center"
		>
		  <div class="flex-grow text-center">Show More</div>
		  <img class="ml-3" src="/down_arrow.svg" />
		</button>
	  </div>
	</div>
	<div class="grid-right hidden md:block">
	  <div class="sticky top-8 p-8 about-bg flex flex-col">
		<div class="w-full flex mb-4 flex-row items-center">
		  <img class="mr-3" src="bookmarks.svg" />
		  <div>
			<p class="heading-text text-sm leading-4 uppercase tracking-wide sm:text-xs">
			  About Reddium
			</p>
		  </div>
		</div>
		<div class="w-full pb-6">
		  <p class="text-sm">
			Ever wanted to browse Reddit while studying at Starbucks? Or
			while sitting on the subway to work? Worried that people around
			you would judge the subreddits you browse and the posts you
			read?
			<br />
			<br />
			{`Reddium is a Medium-themed Reddit client. The Reddium interface
			converts Reddit posts, discussions, and memes into well-crafted
			articles. Medium's layout feels a little more readable than
			Reddit's, removing all distractions and clutter. It also
			bypasses Reddit's frustrating mobile browser.`}
			<br />
			<br />I hope you enjoy this project! Feel free to suggest any
			features or report bugs on GitHub.
		  </p>
		</div>
		<div class="w-full pb-6 hidden">
		  <img class="w-4/12 float-right" src="/signature.png" />
		</div>
		<a
		  href="https://github.com/eightants/reddium/"
		  target="_blank"
		  rel="noopener noreferrer"
		>
		  <button class="mt-2 mx-1 p-2 pl-0 pb-3 cursor-pointer w-full max-w-full btn-black text-white rounded">
			✨ Star on GitHub
		  </button>
		</a>
		<a
		  href="https://ko-fi.com/eightants"
		  target="_blank"
		  rel="noopener noreferrer"
		>
		  <button class="mt-2 mx-1 p-2 pl-0 pb-3 cursor-pointer w-full max-w-full btn-outline-black text-white rounded">
			☕ Buy me a coffee
		  </button>
		</a>
	  </div>
	</div>
  </div>