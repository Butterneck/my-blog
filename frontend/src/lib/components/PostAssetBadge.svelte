<script lang="ts">
	export let asset: string;
	export let onRemove = () => {};
	export let color = 'gray';
	export let onRestore = () => {};
	export let restore = false;

	async function internalOnRemove() {
		await onRemove();
	}

	async function internalOnRestore() {
		await onRestore();
	}
</script>

<span
	id="badge-dismiss-dark"
	class="inline-flex items-center px-2 py-1 me-2 mb-1 text-sm font-medium text-{color}-800 bg-{color}-100 rounded dark:bg-{color}-700 dark:text-{color}-300"
>
	{asset}
	{#if !restore}
		<button
			on:click={internalOnRemove}
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
	{:else}
		<button
			on:click={internalOnRestore}
			type="button"
			class="inline-flex items-center p-1 ms-2 text-sm text-gray-400 bg-transparent rounded-sm hover:bg-gray-200 hover:text-gray-900 dark:hover:bg-gray-600 dark:hover:text-gray-300"
			data-dismiss-target="#badge-dismiss-dark"
			aria-label="Restore"
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 24 24"
				stroke-width="1.5"
				stroke="currentColor"
				class="w-3 h-3"
			>
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					d="M9 15 3 9m0 0 6-6M3 9h12a6 6 0 0 1 0 12h-3"
				/>
			</svg>

			<span class="sr-only">Remove badge</span>
		</button>
	{/if}
</span>
