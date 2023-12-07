<script lang="ts">
	export let data;
	import { envVariables } from '$lib/envVariables';
	import { blogMetaData } from '$lib/blogMetaData';
	import { MetaTags } from 'svelte-meta-tags';
	import PostList from '$lib/components/PostList.svelte';
	import BigContainer from '$lib/components/BigContainer.svelte';
	import RankedPost from '$lib/components/RankedPost.svelte';

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

<BigContainer>
	<PostList posts={data.posts} />
</BigContainer>

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
