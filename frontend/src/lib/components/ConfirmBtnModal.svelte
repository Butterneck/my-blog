<script lang="ts">
	export let btnText: string;
	export let color: string = 'red';
	export let header: string;
	export let message: string;
	export let confirmMsg: string;
	export let cancelMsg: string;

	let isOpen = false;

	export let onConfirm: () => void;
	export let onCancel: () => void = () => {};

	function openModal() {
		isOpen = true;
	}

	async function onCancelInternal() {
		await onCancel();
		isOpen = false;
	}

	async function onConfirmInternal() {
		await onConfirm();
	}

	let bg100: string;
	let bg500: string;
	let bg600: string;
	let hoverBg500: string;
	let text600: string;

	switch (color) {
		case 'purple':
			bg100 = 'bg-purple-100';
			bg500 = 'bg-purple-500';
			bg600 = 'bg-purple-600';
			hoverBg500 = 'hover:bg-purple-500';
			text600 = 'text-purple-600'
			break;
		case 'green':
			bg100 = 'bg-green-100';
			bg500 = 'bg-green-500';
			bg600 = 'bg-green-600';
			hoverBg500 = 'hover:bg-green-500';
			text600 = 'text-green-600'
			break;
		default:
			bg100 = 'bg-red-100';
			bg500 = 'bg-red-500';
			bg600 = 'bg-red-600';
			hoverBg500 = 'hover:bg-red-500';
			text600 = 'text-red-600'
			break;
		
	}
</script>

<span class="align-middle">
	<button
		class="ml-2 mt-2 p-1 px-3 text-sm cursor-pointer max-w-full {bg500} text-white outline-1px rounded"
		type="button"
		on:click={openModal}
	>
		{btnText}
	</button>
</span>

<div
	class="{isOpen ? 'block' : 'hidden'} relative z-10"
	aria-labelledby="modal-title"
	role="dialog"
	aria-modal="true"
>
	<div class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"></div>

	<div class="fixed inset-0 z-10 w-screen overflow-y-auto">
		<div class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
			<div
				class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg"
			>
				<div class="bg-white px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
					<div class="sm:flex sm:items-start">
						<div
							class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full {bg100} sm:mx-0 sm:h-10 sm:w-10"
						>
							<svg
								class="h-6 w-6 {text600}"
								fill="none"
								viewBox="0 0 24 24"
								stroke-width="1.5"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									d="M12 9v3.75m-9.303 3.376c-.866 1.5.217 3.374 1.948 3.374h14.71c1.73 0 2.813-1.874 1.948-3.374L13.949 3.378c-.866-1.5-3.032-1.5-3.898 0L2.697 16.126zM12 15.75h.007v.008H12v-.008z"
								/>
							</svg>
						</div>
						<div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
							<h3 class="text-base font-semibold leading-6 text-gray-900" id="modal-title">
								{header}
							</h3>
							<div class="mt-2">
								<p class="text-sm text-gray-500">
									{message}
								</p>
							</div>
						</div>
					</div>
				</div>
				<div class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
					<button
						type="button"
						class="inline-flex w-full justify-center rounded-md {bg600} px-3 py-2 text-sm font-semibold text-white shadow-sm {hoverBg500} sm:ml-3 sm:w-auto"
						on:click={onConfirmInternal}>{confirmMsg}</button
					>
					<button
						type="button"
						class="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 hover:bg-gray-50 sm:mt-0 sm:w-auto"
						on:click={onCancelInternal}>{cancelMsg}</button
					>
				</div>
			</div>
		</div>
	</div>
</div>
