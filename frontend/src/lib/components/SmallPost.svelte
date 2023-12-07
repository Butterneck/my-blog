<script lang="ts">
	import { format } from 'fecha';

	export let post: Post;
</script>

<article class="flex flex-col items-start justify-between">
	<div class="pb-8 flex items-start w-full">
		<div class="mr-5 w-full min-w-0 block">
			<!-- POST AUTHOR -->
			<div class="w-full block pr-5 flex-grow">
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
			</div>

			<!-- POST TITLE -->
			<div class="mt-2">
				<a href={`/blog/${post.slug}`}>
					<h2 class="text-base font-bold tracking-normal text-ellipsis text-zinc-900 line-clamp-2">
						{post.title}
					</h2>
				</a>
			</div>

			<!-- POST METADATA -->
			<div class="sub-text mt-2">
				<span>{format(new Date(post.createdAt * 1000), 'MMM D')}</span>
				<span class="px-0.5">·</span>
				<span>9 min read</span>
			</div>
		</div>

		<!-- POST IMAGE -->
		<a class="cursor-pointer m-0 p-0" href={`/blog/${post.slug}`}>
			<img
				class="align-middle w-28 h-28 overflow-clip aspect-square max-w-none"
				src="https://media.istockphoto.com/id/583809524/it/foto/alberta-selvaggia-vicino-a-banff.jpg?s=612x612&w=0&k=20&c=wNnbWawMkk9idvuKsw4gxdJ8OQX0PCxgGDcsTdrwM4E="
				alt=""
			/>
		</a>
	</div>
</article>