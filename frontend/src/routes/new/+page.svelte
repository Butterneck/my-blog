<script>
	import { createPost } from '$lib/postsApi';

	let title = '';
	let body = '';

	async function createNewPost() {
		const newPost = await createPost({
			title,
			body
		});
	}
</script>

<h1>New Post</h1>
<input bind:value={title} placeholder="Post title" />
<textarea bind:value={body} placeholder="Post body" />
<button on:click={createNewPost}>Create New Post</button>

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

	/* Textarea Style */
	textarea {
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
	textarea:focus {
		border-color: #007bff;
		outline: none;
		/* Add any additional styles for the focus state */
	}
</style>
