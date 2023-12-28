<script lang="ts">
	import { getCurrentUser } from '$lib/auth';
	import { onMount } from 'svelte';

	let ok = false;

	onMount(async () => {
		if (!(await getCurrentUser())) {
			return;
		}

		ok = true;
	});

	export let data: {
		body: {
			post: Post;
		};
	};

	import Editor from '$lib/components/Editor.svelte';
	import NotFound from '$lib/components/NotFound.svelte';

	const post = data.body.post;
</script>

{#if ok}
	<Editor {post} />
{:else}
	<NotFound />
{/if}
