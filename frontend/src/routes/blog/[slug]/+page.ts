import { getPost } from "$lib/api";

export async function load({ params }) {
  const slug = params.slug.toLowerCase();

  const post = await getPost(slug);

  if (post) {
    return {
      body: { post },
    };
  }

  return {
    body: { message: "Not found", post: null },
  };
}
