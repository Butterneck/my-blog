export const ssr = false;
export const prerender = false;
import { DefaultService } from "$lib/generated/backend-client";
import type { Post } from "$lib/generated/backend-client";


export async function load(): Promise<{ posts: Post[], trendingPosts: Post[] }> {
  let posts = await DefaultService.getApiV1Posts();

  return { posts: posts, trendingPosts: posts.slice(0, 4) };
}
