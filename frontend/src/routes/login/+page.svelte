<script lang="ts">
	import { getCurrentUser, logIn, logOut } from '$lib/auth';
	import { onMount } from 'svelte';

	// ==
	let determinedLoggedIn = false;
	let loggedIn = false;
	let username = '';
	let password = '';
	//==

	onMount(async () => {
		loggedIn = await isUserLoggedIn();
		determinedLoggedIn = true;
	});

	async function isUserLoggedIn() {
		const user = await getCurrentUser();
		if (user) {
			console.log('user is logged in');
			return true;
		} else {
			console.log('user is not logged in');
		}

		return false;
	}

	async function handleLogin({ username, password }: { username: string; password: string }) {
		const loginRes = await logIn({ username, password });
		if (loginRes) {
			console.log('login success');
			loggedIn = true;
		} else {
			console.log('login failed');
		}
	}

	async function handleSignOut() {
		await logOut();
		loggedIn = false;
	}
</script>

<div class="container">
	{#if determinedLoggedIn && loggedIn}
		<div>
			<p>Welcome, you are already logged in!</p>
			<button on:click={handleSignOut}>Logout</button>
		</div>
	{:else if determinedLoggedIn && !loggedIn}
		<div>
			<label for="uname"><b>Username</b></label>
			<input type="text" placeholder="Enter Username" name="uname" required bind:value={username} />

			<label for="psw"><b>Password</b></label>
			<input
				type="password"
				placeholder="Enter Password"
				name="psw"
				required
				bind:value={password}
			/>

			<button type="submit" on:click={() => handleLogin({ username, password })}>Login</button>
		</div>
	{:else}
		<!-- <p>Checking if you are logged in...</p> -->
	{/if}
</div>

<style>
	button {
		display: inline-block;
		padding: 10px 20px;
		font-size: 16px;
		font-weight: bold;
		text-align: center;
		text-decoration: none;
		cursor: pointer;
		border: 2px solid #333;
		border-radius: 4px;
		background-color: #f0f0f0;
		color: #333;
		transition:
			background-color 0.3s ease,
			color 0.3s ease,
			border-color 0.3s ease;
	}

	button:hover {
		background-color: #333;
		color: #fff;
		border-color: #fff;
	}

	/* Input Style */
	input {
		display: block;
		width: 100%;
		padding: 10px;
		margin-bottom: 15px;
		font-size: 16px;
		border: 1px solid #ccc;
		border-radius: 4px;
		box-sizing: border-box;
	}

	/* Optional: Add styles for focus state */
	input:focus {
		border-color: #007bff;
		outline: none;
		/* Add any additional styles for the focus state */
	}
</style>
