<script lang="ts">
	import PostMetadata from './PostMetadata.svelte';
	import PostAuthor from './PostAuthor.svelte';

	export let post: Post;
</script>

<article class="flex flex-col items-start justify-between">
	<div class="pb-8 flex items-start w-full">
		<div class="mr-5 w-full min-w-0 block">
			<!-- POST AUTHOR -->
			<div class="w-full block pr-5 flex-grow">
				<PostAuthor />
			</div>

			<!-- POST TITLE -->
			<div class="mt-2">
				<a href={`/blog/${post.slug}`}>
					<h2 class="text-base font-bold tracking-normal text-ellipsis text-zinc-900 line-clamp-2">
						{post.draft?.title || post.title}
					</h2>
				</a>
			</div>

			<!-- POST METADATA -->
			<PostMetadata post={post} />
		</div>

		<!-- POST IMAGE -->
		<a href={`/blog/${post.slug}`}>
			<img
				class="align-middle w-28 md:w-32 lg:w-36 xl:w-40 h-28 overflow-clip aspect-square max-w-none"
				src="https://media.istockphoto.com/id/583809524/it/foto/alberta-selvaggia-vicino-a-banff.jpg?s=612x612&w=0&k=20&c=wNnbWawMkk9idvuKsw4gxdJ8OQX0PCxgGDcsTdrwM4E="
				alt=""
			/>
		</a>
	</div>
</article>