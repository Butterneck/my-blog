<script lang="ts">
	import { Editor } from 'bytemd';
	import highlight from '@bytemd/plugin-highlight';
	import gfm from '@bytemd/plugin-gfm';
	import 'highlight.js/styles/default.css';

	/**
	 * @type {any}
	 */
	let value = 'Ciao';
	const plugins = [
		gfm(),
		highlight()
		// Add more plugins here
	];

	/**
	 * @param {{ detail: { value: any; }; }} e
	 */
	function handleChange(e) {
		value = e.detail.value;
	}

	function uploadImages(e: File) {
		// TODO: Implement logic to upload images to S3
		console.log(e);
		console.log('upload images');
	}

	function publishPost() {
		console.log('publish post');
	}
</script>

<Editor {value} {plugins} {uploadImages} on:change={handleChange} />

<div class="flex justify-center">
	<button
		class="my-4 py-1 px-28 text-sm cursor-pointer max-w-full btn-black text-white outline-1px rounded"
		on:click={publishPost}
	>
		Publish
	</button>
</div>
