<script lang="ts">
	export let posts: Post[];
	import LargePost from './LargePost.svelte';
	import SmallPost from './SmallPost.svelte';

</script>

<div class="post-list">
	{#if posts.length > 0}
		<LargePost post={posts[0]} />

		{#each posts.slice(1) as post}
			<SmallPost {post} />
		{/each}
	{/if}
</div>
