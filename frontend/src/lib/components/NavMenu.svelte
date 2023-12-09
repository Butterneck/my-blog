<script lang="ts">
	import { signInWithRedirect, signOut } from 'aws-amplify/auth';
	import { Hub } from '@aws-amplify/core';

	const listener = (data: any) => {
		console.log(data);
	};

	Hub.listen('auth', listener);

	async function handleLogin() {
		await signInWithRedirect();
	}

	async function handleLogout() {
		await signOut();
	}
</script>

<div class="items-center flex flex-row h-full justify-end">
	<button
		class="my-4 ml-4 p-1 px-3 text-sm cursor-pointer max-w-full btn-black text-white outline-1px rounded"
		on:click={handleLogin}
	>
		Login
	</button>
</div>
