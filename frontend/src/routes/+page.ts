export const ssr = false;
export const prerender = false;


export async function load() {
  let posts = await fetch("https://blog.butterneck.me/api/v1/posts",)
    .then(response => response.json())
    // .then(response => response.text())
    // .then(result => console.log(result))
    .catch(error => console.log('error', error));

  return { posts: posts };
}
