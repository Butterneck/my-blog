export const ssr = false;
export const prerender = false;
import { getPosts } from '$lib/api';

export async function load(): Promise<{ posts: Post[], nextPageToken?: string }> {

  const { nextPageToken, posts } = await getPosts({});

  return { posts: posts, nextPageToken: nextPageToken };
}
