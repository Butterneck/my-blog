import { getIdentityToken } from "./auth";

export async function getPosts() {
  const response = await fetch('https://blog.butterneck.me/api/v1/posts');
  const posts = await response.json();
  return posts;
}

export async function createPost(post: NewPost) {
  const idToken = await getIdentityToken();
  const response = await fetch('https://blog.butterneck.me/api/v1/posts', {
    method: 'POST',
    body: JSON.stringify(post),
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${idToken}`
    }
  });
  const createdPost = await response.json();
  return createdPost;
}