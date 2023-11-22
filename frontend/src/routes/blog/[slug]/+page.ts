export async function load({ params }) {
  const slug = params.slug.toLowerCase();

  const post = await fetch(`https://blog.butterneck.me/api/v1/posts/${slug}`)
    .then((response) => response.json())
    .catch((error) => console.log("error", error));

  if (post) {
    return {
      body: { post },
    };
  }

  return {
    body: { message: "Not found", post: null },
  };
}
