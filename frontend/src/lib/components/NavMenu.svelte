<script lang="ts">
	import { signInWithRedirect, signOut } from 'aws-amplify/auth';
	import { Hub } from '@aws-amplify/core';
	import { onMount } from 'svelte';
	import { getCurrentUser } from '$lib/auth';

	let isUserLoggedIn = false;
	async function handleLogin() {
		await signInWithRedirect();
	}

	async function handleLogout() {
		await signOut();
	}

	onMount(async () => {
		isUserLoggedIn = !!(await getCurrentUser());
	});
</script>

<div class="items-center flex flex-row h-full justify-end">
	{#if isUserLoggedIn}
		<button
			class="hidden sm:block my-4 ml-4 p-1 px-3 sub-opacity-68 link-black-hover text-sm cursor-pointer max-w-full btn-outline-black rounded"
		>
			New post
		</button>
		<button
			class="my-4 ml-4 p-1 px-3 text-sm cursor-pointer max-w-full btn-black text-white outline-1px rounded"
			on:click={handleLogout}
		>
			Logout
		</button>
	{:else}
		<button
			class="my-4 ml-4 p-1 px-3 text-sm cursor-pointer max-w-full btn-black text-white outline-1px rounded"
			on:click={handleLogin}
		>
			Login
		</button>
	{/if}
</div>
