<script lang="ts">
	import PostMetadata from './PostMetadata.svelte';
	export let post: Post;
</script>

<div class="mid-container">
	<div class="mr-4 pb-8 flex overflow-hidden">
		<!-- RANK  -->
		<div class="h-full -mt-2 mr-4">
			<span class="main-grey text-3xl font-bold"> 1 </span>
		</div>

		<!-- POST AUTHOR -->
		<div class="w-full block flex-grow">
			<div class="sub-text font-medium main-black mt-1 flex flex-row items-center">
				<div>
					<img
						class="rounded-full w-5 h-5"
						src="https://reddium.vercel.app/avatars/avatar_15.jpg"
						alt=""
					/>
				</div>
				<a href={`/user/foo`}><span class="ml-2 font-semibold">Butterneck</span></a>
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
	</div>
</div>
