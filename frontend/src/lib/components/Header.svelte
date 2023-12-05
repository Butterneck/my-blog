<script>
	import { blogMetaData } from '$lib/blogMetaData';
	import { signInWithRedirect } from 'aws-amplify/auth';
	import { signOut } from 'aws-amplify/auth';
	import { getCurrentUser } from '$lib/auth';
	import { Hub } from 'aws-amplify/utils';

	// @ts-ignore
	const listener = (data) => {
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

<nav class="header">
	<a href="/" class="header__title">{blogMetaData.blogTitle}</a>
	<ul class="header__links">
		<li class="header__links_item">
			<a href="/" class="header__links__item__link">Home</a>
		</li>
		<li class="header__links_item">
			<a href="/about" class="header__links__item__link">About</a>
		</li>
		<!-- <li class="header__links_item">
			<a href="/new" class="header__links__item__link">New Post</a>
		</li>
		<li class="header__links_item">
			<span class="header__links__item__link" on:click={handleLogout}>Logout</span>
		</li>
		<li class="header__links_item">
			<span class="header__links__item__link" on:click={handleLogin}>Login</span>
		</li> -->
	</ul>
</nav>

<style>
	.header {
		height: 6rem;
		display: flex;
		justify-content: space-between;
		flex-direction: row;
	}

	.header__title {
		margin-left: 1rem;
		color: var(--main-color);
		text-decoration: none;
		font-weight: bold;
		font-size: 1.5rem;
		line-height: 6.2rem;
	}

	.header__links {
		display: flex;
		flex-direction: row;
		align-items: center;
		list-style: none;
	}

	.header__links_item {
		margin-left: 1rem;
	}

	.header__links_item:last-child {
		margin-right: 1rem;
	}

	.header__links__item__link {
		padding-bottom: 0.2rem;
		color: var(--secondary-color);
		text-decoration: none;
	}

	.header__links__item__link:hover {
		border-bottom: 1.5px solid var(--secondary-color);
	}
</style>
