<script lang="ts">
	import LargePost from './LargePost.svelte';
	import SmallPost from './SmallPost.svelte';
	export let posts: Post[];
</script>

<div class="flex home-container-top">
	<div class="mid-container xl:px-8 md:pr-4 px-0">
		<LargePost post={posts[0]} />
	</div>

	<div class="mid-container xl:px-8 md:pl-4 px-0">
		{#each posts.slice(1) as post}
			<SmallPost {post} />
		{/each}
	</div>
</div>
