<script lang="ts">
	export let posts: Post[];
	import LargePost from './LargePost.svelte';
	import RankedPost from './RankedPost.svelte';
	import SmallPost from './SmallPost.svelte';
</script>

<div class="mid-container lg:px-8 md:pr-4 px-0">
	<LargePost post={posts[0]} />
</div>

<div class="mid-container lg:px-8 md:pl-4 px-0">
	{#each posts.slice(1, 5) as post}
		<SmallPost {post} />
	{/each}
</div>
