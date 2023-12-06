<script lang="ts">
	import { format } from 'fecha';

	export let post: Post;
</script>

<div class="large-post">
	<div class="overflow-hidden sm-border-b pb-8 mb-6">
		<!-- POST IMAGE -->
		<a href="TODO">
			<div class="w-full shimmer-bg cursor-pointer">
				<img
					class="big-post-img"
					src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fi.pinimg.com%2Foriginals%2F48%2Fa5%2Fc0%2F48a5c07d54cd6e36a4bc87f4376fb696.jpg&f=1&nofb=1&ipt=31d64d5ce92427ad837b375e5ae6aa4e2001b1204ebfa456285bb4d44e771326&ipo=images"
					alt=""
				/>
			</div>
		</a>

		<!-- POST AUTHOR -->
		<div class="sub-text font-medium main-black mt-3 flex flex-row items-center">
			<div>
				<img
					class="rounded-full w-5 h-5"
					src="https://reddium.vercel.app/avatars/avatar_15.jpg"
					alt=""
				/>
			</div>
			<a href={`/user/foo`}>
				<span class="ml-2 font-semibold">Butterneck</span>
			</a>
		</div>

		<!-- POST TITLE -->
		<a href={`/blog/${post.slug}`}>
			<h2 class="text-2xl mt-2 leading-6 line-clamp-2">{post.title}</h2>
		</a>

		<!-- POST METADATA -->
		<div class="sub-text mt-2">
			<span>{format(new Date(post.createdAt * 1000), 'MMM D')}</span>
			<span class="px-2">·</span>
			<a class="link-black-hover" href="TODO"> 9 min read </a>
		</div>
	</div>
</div>

<style lang="postcss">
	.big-post-img {
		min-height: 240px;
		background-size: cover;
	}
</style>
