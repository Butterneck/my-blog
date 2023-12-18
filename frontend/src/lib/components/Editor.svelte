<script lang="ts">
	export let post: Post | undefined = undefined;

	let title: string = post?.draft?.title || post?.title || '';
	let body: string = post?.draft?.body || post?.body || '';
	let attachments: PostAttachment[] = post?.draft?.attachments || post?.attachments || [];
	let attachmentInput: FileList;

	import 'highlight.js/styles/default.css';
	import { createPost, updatePost } from '$lib/api';

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
	
	// Prevent default drag behaviors
	function handleDragOver(event: DragEvent) {
		event.preventDefault();
	}
	
	// Handle dropped files
	function handleDrop(event: DragEvent) {
		event.preventDefault();
		handleNewAttachments(event.dataTransfer?.files || new DataTransfer().files);
	}

	// Handle clicks on file input
	$: if (attachmentInput) {
		handleNewAttachments(attachmentInput);
		attachmentInput = new DataTransfer().files
	}

	function handleNewAttachments(newAttachments: FileList) {
		for (const attachment of newAttachments) {
			attachments.push({ name: attachment.name });
		}
		attachments = attachments
	}

	// Remove files from attachments
	function removeAttachment(attachment: PostAttachment) {
		attachments = attachments.filter((i: PostAttachment) => i.name !== attachment.name);
	}
</script>

<div class="home-container-top my-4">
	<!-- TITLE -->
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

	<!-- BODY -->
	<div class="relative mt-4" style="height: 75vh">
		<!-- <div class="relative w-full h-full">
			  <textarea
				bind:value={body}
				class="resize-none block w-full h-full rounded-md border-0 py-1.5 pl-4 pr-4 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
				placeholder=" "></textarea>
			</div>
		  </div> -->

		<textarea
			bind:value={body}
			id="message"
			class="resize-none block w-full h-full rounded-md border-0 py-1.5 pl-4 pr-4 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
			placeholder="Write your thoughts here..."
		></textarea>
	</div>

	<!-- Uploaded images -->
	<div class="mt-2">
		{#each attachments as attachment}
			<span
				id="badge-dismiss-dark"
				class="inline-flex items-center px-2 py-1 me-2 mb-1 text-sm font-medium text-gray-800 bg-gray-100 rounded dark:bg-gray-700 dark:text-gray-300"
			>
				{attachment.name}
				<button
					on:click={() => removeAttachment(attachment)}
					type="button"
					class="inline-flex items-center p-1 ms-2 text-sm text-gray-400 bg-transparent rounded-sm hover:bg-gray-200 hover:text-gray-900 dark:hover:bg-gray-600 dark:hover:text-gray-300"
					data-dismiss-target="#badge-dismiss-dark"
					aria-label="Remove"
				>
					<svg
						class="w-2 h-2"
						aria-hidden="true"
						xmlns="http://www.w3.org/2000/svg"
						fill="none"
						viewBox="0 0 14 14"
					>
						<path
							stroke="currentColor"
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
						/>
					</svg>
					<span class="sr-only">Remove badge</span>
				</button>
			</span>
		{/each}
	</div>

	<!-- UPLOAD IMAGES -->
	<div class="flex items-center justify-center w-full mt-2">
		<label
			on:drop={handleDrop}
			on:dragover={handleDragOver}
			for="dropzone-file"
			class="flex flex-col items-center justify-center w-full border-2 border-gray-300 border-dashed rounded-lg cursor-pointer bg-gray-50 dark:hover:bg-bray-800 dark:bg-gray-700 hover:bg-gray-100 dark:border-gray-600 dark:hover:border-gray-500 dark:hover:bg-gray-600"
		>
			<div class="flex flex-col items-center justify-center pt-5 pb-4">
				<svg
					class="w-8 h-8 mb-4 text-gray-500 dark:text-gray-400"
					aria-hidden="true"
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 20 16"
				>
					<path
						stroke="currentColor"
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"
					/>
				</svg>
				<p class="mb-2 text-sm text-gray-500 dark:text-gray-400">
					<span class="font-semibold">Click to upload</span> or drag and drop
				</p>
				<!-- <p class="text-xs text-gray-500 dark:text-gray-400">SVG, PNG, JPG or GIF</p> -->
			</div>
			<input bind:files={attachmentInput} id="dropzone-file" type="file" class="hidden" multiple />
		</label>
	</div>

	<!-- SUBMIT -->
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
</div>
