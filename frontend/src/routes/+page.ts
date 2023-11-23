export const ssr = false;
export const prerender = false;
import { getPosts } from "$lib/postsApi";


export async function load() {
  let posts = await getPosts();

  return { posts: posts };
}
