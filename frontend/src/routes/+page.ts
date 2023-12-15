export const ssr = false;
export const prerender = false;
import { getPosts } from '$lib/api';

export async function load(): Promise<{ posts: Post[], trendingPosts: Post[] }> {

  const posts = await getPosts();

  return { posts: posts, trendingPosts: posts.slice(0, 4) };
}
