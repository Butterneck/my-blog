<script lang="ts">
    import { format } from 'fecha';

    export let post: Post;
</script>

<div class="sub-text mt-2">
    <span>{format(new Date(post.creationDate * 1000), 'MMM D')}</span>
    <span class="px-2">·</span>
    <span> 9 min read </span>

    {#if post.draft && (post.draft?.title != "" || post.draft?.body != "")}
        <!-- TODO: on click go to edit page -->
        <button
            class=" ml-4 p-1 px-3 text-sm cursor-pointer max-w-full btn-black text-white outline-1px rounded"
        >
            Draft
        </button>
    {/if}

</div>