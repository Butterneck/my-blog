<script lang="ts">
	export let post: Post | undefined = undefined;

	let title: string = post?.draft?.title || post?.title || '';
	let body: string = post?.draft?.body || post?.body || '';
	let assets: string[] = post?.draft?.assets || post?.assets || [];

	let newAssets: PostAsset[] = [];
	let deletedAssets: string[] = [];
	let assetInput: FileList;

	import 'highlight.js/styles/default.css';
	import { createPost, updatePost } from '$lib/api';
	import PostAssetBadge from './PostAssetBadge.svelte';

	async function save() {
		if (!post) {
			await createPost({
				title,
				body,
				assets: newAssets.reduce<File[]>((acc, asset) => {
					if (asset.file) {
						acc.push(asset.file);
					}
					return acc;
				}, [])
			});
			window.location.href = '/';
		} else {

			const foo = newAssets.reduce<File[]>((acc, asset) => {
					console.log(asset)
					if (asset.file) {
						acc.push(asset.file);
					}
					return acc;
				}, [])

				console.log(foo)
			await updatePost({
				slug: post.slug,
				title,
				body,
				newAssets: foo,
				deletedAssets,
			});
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
		handleNewAssets(event.dataTransfer?.files || new DataTransfer().files);
	}

	// Handle clicks on file input
	$: if (assetInput) {
		handleNewAssets(assetInput);
		assetInput = new DataTransfer().files;
	}

	function handleNewAssets(_newAssets: FileList) {
		for (const asset of _newAssets) {
			newAssets.push({ name: asset.name, file: asset });
		}
		newAssets = newAssets;
	}

	function removeNewAsset(asset: PostAsset) {
		console.log("removeNewAsset")
		newAssets = newAssets.filter((i: PostAsset) => i.name !== asset.name);
	}

	function removeExistingAsset(asset: string) {
		console.log("removeExistingAsset")
		assets = assets.filter((i: string) => i !== asset);
		deletedAssets.push(asset);
		deletedAssets = deletedAssets;
	}

	function restoreDeletedAsset(asset: string) {
		console.log("restoreDeletedAsset")
		deletedAssets = deletedAssets.filter((i: string) => i !== asset);
		assets.push(asset);
		assets = assets;
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
		{#each newAssets as newAsset}
			<PostAssetBadge
				asset={newAsset.name}
				color="green"
				onRemove={() => removeNewAsset(newAsset)}
			/>
		{/each}

		{#each assets as asset}
			<PostAssetBadge
				asset={asset}
				onRemove={() => removeExistingAsset(asset)}
			/>
		{/each}

		{#each deletedAssets as deletedAsset}
			<PostAssetBadge
				asset={deletedAsset}
				color="red"
				restore={true}
				onRestore={() => restoreDeletedAsset(deletedAsset)}
			/>
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
			<input bind:files={assetInput} id="dropzone-file" type="file" class="hidden" multiple />
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
