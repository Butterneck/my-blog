<script lang="ts">
    import { format } from 'fecha';
    import PostAuthor from './PostAuthor.svelte';

    export let post: Post;
</script>

<div class="sub-text mt-2">
    <span>{format(new Date(post.createdAt * 1000), 'MMM D')}</span>
    <span class="px-2">·</span>
    <a class="link-black-hover" href="TODO"> 9 min read </a>
</div>