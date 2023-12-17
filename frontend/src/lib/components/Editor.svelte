<script lang="ts">
	export let post: Post|undefined = undefined;

	let title: string = post?.draft?.title || post?.title || '';
	let body: string = post?.draft?.body || post?.body || '';

	import { Editor } from 'bytemd';
	import highlight from '@bytemd/plugin-highlight';
	import gfm from '@bytemd/plugin-gfm';
	import 'highlight.js/styles/default.css';
	import { createPost, updatePost } from '$lib/api';

	const editorPlugins = [
		gfm(),
		highlight()
		// Add more plugins here
	];

	function handleChange(e: any) {
		body = e.detail.value;
	}

	function uploadImages(e: File) {
		// TODO: Implement logic to upload images to S3
		console.log(e);
		console.log('upload images');
	}

	async function save() {
		if (!post) {
			await createPost({ title, body });
			window.location.href = '/';
		} else {
			post.title = title;
			post.body = body;
			await updatePost(post);
			window.location.href = '/blog/' + post.slug;
		}
	}
</script>

<div class="home-container-top my-4">
	<!-- <label for="price" class="block text-sm font-medium leading-6 text-gray-900">Title</label> -->
	<div class="relative mt-2 rounded-md shadow-sm">
		<input
			bind:value={title}
			type="text"
			name="price"
			id="price"
			class="block w-full rounded-md border-0 py-1.5 pl-4 pr-4 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
			placeholder="Title"
		/>
	</div>
</div>

<Editor value={body} plugins={editorPlugins} {uploadImages} on:change={handleChange} />

<div class="flex justify-center">
	<button
		class="my-4 py-1 px-28 text-sm cursor-pointer max-w-full btn-black text-white outline-1px rounded"
		on:click={save}
	>
		{#if !post}
			Create draft
		{:else}
			Update draft
		{/if}
	</button>
</div>
