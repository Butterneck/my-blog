<script lang="ts">
	export let data: {
		posts: Post[];
		// trendingPosts: Post[];
		nextPageToken: string;
	};
	import { envVariables } from '$lib/envVariables';
	import { blogMetaData } from '$lib/blogMetaData';
	import { MetaTags } from 'svelte-meta-tags';
	import LatestPosts from '$lib/components/LatestPosts.svelte';
	// import TrendingPosts from '$lib/components/TrendingPosts.svelte';
	import PostsArchive from '$lib/components/PostsArchive.svelte';

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
<LatestPosts posts={data.posts.slice(0, 5)} />

<!-- TRENDING POSTS -->
<!-- <TrendingPosts posts={data.trendingPosts} /> -->

<!-- ALL POSTS -->
<PostsArchive posts={data.posts} nextPageToken={data.nextPageToken} />

