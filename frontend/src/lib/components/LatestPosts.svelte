<script lang="ts">
    import LargePost from "./LargePost.svelte";
    import SmallPost from "./SmallPost.svelte";
    export let posts: Post[];
</script>

<div class="flex md:w-auto main-container max-width-main mx-6 sm:mx-12 pb-5 xl:mx-auto">
	<div class="mid-container lg:px-8 md:pr-4 px-0">
        <LargePost post={posts[0]} />
    </div>
    
    <div class="mid-container lg:px-8 md:pl-4 px-0">
        {#each posts.slice(1) as post}
            <SmallPost {post} />
        {/each}
    </div>
    
</div>