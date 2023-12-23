<script lang="ts">
	export let data: {
		post: Post;
		renderedBody?: string;
	};
	import { MetaTags } from 'svelte-meta-tags';
	import { blogMetaData } from '$lib/blogMetaData';
	import { envVariables } from '$lib/envVariables';
	import { getCurrentUser } from '$lib/auth';
	import { publishPost } from '$lib/api';
	import { isPostPublised } from '$lib/utils';

	const post = data.post;

	const meta = {
		title: `${post.title} | ${blogMetaData.blogTitle}`,
		// description: post.description ?? post.title,
		description: post.title,
		url: `/blog/${post.slug}`,
		siteName: blogMetaData.blogTitle,
		author: blogMetaData.blogTitle,
		image: {
			url: `${envVariables.basePath}/background.jpeg`,
			width: 1000,
			height: 523,
			alt: 'image'
		}
	};

	async function publish() {
		await publishPost(post.slug);
		window.location.href = '/blog/' + post.slug;
	}

	async function unpublish() {
		// await unpublishPost(post.slug, true);
		// window.location.href = '/blog/' + post.slug;
	}
</script>

<MetaTags
	title={meta.title}
	description={meta.description}
	canonical={meta.url}
	openGraph={{
		article: {
			authors: [meta.author]
		},
		images: [
			{
				...meta.image
			}
		],
		description: meta.description,
		siteName: meta.siteName,
		title: meta.title,
		type: 'article',
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

<!-- <div class="md:w-auto sm:mx-auto max-w-680 pb-2 mt-6 mx-6 w-auto"> -->
<div class="md:w-auto max-w-680 mx-6 md:mx-auto xl:mx-auto">
	<!-- POST TITLE -->
	<span>
		<div>
			<span
				class="align-middle heading-font sm:text-5xl font-bold sm:leading-tight text-3xl leading-9 pt-8"
				>{post.draft?.title || post.title}</span
			>

			{#await getCurrentUser() then currentUser}
				{#if currentUser}
					<span class="align-middle">
						<a href="/blog/{post.slug}/edit">
							<button
								class="ml-10 mt-2 p-1 px-3 text-sm cursor-pointer max-w-full btn-black text-white outline-1px rounded"
							>
								Edit
							</button>
						</a>
					</span>

					{#if post.draft?.body || post.draft?.title}
						<span class="align-middle">
							<button
								class="ml-2 mt-2 p-1 px-3 text-sm cursor-pointer max-w-full bg-green-500 text-white outline-1px rounded"
								on:click={publish}
							>
								Publish
							</button>
						</span>
					{/if}

					{#if isPostPublised(post)}
						<span class="align-middle">
							<button
								class="ml-2 mt-2 p-1 px-3 text-sm cursor-pointer max-w-full bg-red-500 text-white outline-1px rounded"
								on:click={unpublish}
							>
								Unpublish
							</button>
						</span>
					{/if}
				{/if}
			{/await}
		</div>

		<!-- POST IMAGE -->
		<figure class="mt-12">
			<img
				class="w-full shimmer-bg"
				src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi.pinimg.com%2Foriginals%2F48%2Fa5%2Fc0%2F48a5c07d54cd6e36a4bc87f4376fb696.jpg&f=1&nofb=1&ipt=31d64d5ce92427ad837b375e5ae6aa4e2001b1204ebfa456285bb4d44e771326&ipo=images"
				width="100%"
				alt=""
			/>
			<figcaption class="mt-2 mx-auto sub-opacity-54 text-center text-sm">Caption</figcaption>
		</figure>

		<!-- POST CONTENT -->
		<div
			class="mb-20 mt-6 heading-font text-xl whitespace-pre-line main-black post-content sm:text-lg"
		>
			<!-- {post.draft?.body || post.body} -->
			{@html data.renderedBody}
		</div>

		<!-- POST FOOTER -->
		<!-- <div class="w-full mt-4 pt-4 flex flex-row justify-between items-center"> -->
		<!-- CLAP BUTTON -->
		<!-- <div class="flex flex-row items-center">
			<div class="flex flex-row items-center sub-opacity-54 tracking-tight">
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
						d="M9.813 15.904L9 18.75l-.813-2.846a4.5 4.5 0 00-3.09-3.09L2.25 12l2.846-.813a4.5 4.5 0 003.09-3.09L9 5.25l.813 2.846a4.5 4.5 0 003.09 3.09L15.75 12l-2.846.813a4.5 4.5 0 00-3.09 3.09zM18.259 8.715L18 9.75l-.259-1.035a3.375 3.375 0 00-2.455-2.456L14.25 6l1.036-.259a3.375 3.375 0 002.455-2.456L18 2.25l.259 1.035a3.375 3.375 0 002.456 2.456L21.75 6l-1.035.259a3.375 3.375 0 00-2.456 2.456zM16.894 20.567L16.5 21.75l-.394-1.183a2.25 2.25 0 00-1.423-1.423L13.5 18.75l1.183-.394a2.25 2.25 0 001.423-1.423l.394-1.183.394 1.183a2.25 2.25 0 001.423 1.423l1.183.394-1.183.394a2.25 2.25 0 00-1.423 1.423z"
					/>
				</svg>

				<div>
					<p class="ml-2">{`1000 claps`}</p>
				</div>
			</div> -->

		<!-- COMMENTS -->
		<!-- <div class="ml-4 flex flex-row items-center sub-opacity-54 tracking-tight">
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
						d="M12 20.25c4.97 0 9-3.694 9-8.25s-4.03-8.25-9-8.25S3 7.444 3 12c0 2.104.859 4.023 2.273 5.48.432.447.74 1.04.586 1.641a4.483 4.483 0 01-.923 1.785A5.969 5.969 0 006 21c1.282 0 2.47-.402 3.445-1.087.81.22 1.668.337 2.555.337z"
					/>
				</svg>

				<div>
					<p class="ml-1">23 comments</p>
				</div>
			</div> -->
		<!-- </div> -->

		<!-- SHARE BUTTON -->
		<!-- <div class="flex flex-row items-center sub-opacity-54 tracking-tight">
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
					d="M9 8.25H7.5a2.25 2.25 0 00-2.25 2.25v9a2.25 2.25 0 002.25 2.25h9a2.25 2.25 0 002.25-2.25v-9a2.25 2.25 0 00-2.25-2.25H15m0-3l-3-3m0 0l-3 3m3-3V15"
				/>
			</svg>
		</div> -->
		<!-- </div> -->
	</span>
</div>
