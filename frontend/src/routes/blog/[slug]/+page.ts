import { getPost } from "$lib/api";
import rehypeStringify from 'rehype-stringify'
import remarkParse from 'remark-parse'
import remarkRehype from 'remark-rehype'
import rehypePrettyCode from 'rehype-pretty-code';
import { unified } from 'unified'

export async function load({ params }) {
  const slug = params.slug.toLowerCase();

  const post = await getPost(slug);

  const body = post.draft?.body ? post.draft.body : post.body;

  const file = await unified()
    .use(remarkParse)
    .use(remarkRehype)
    .use(rehypePrettyCode, {
      theme: 'nord',
    })
    .use(rehypeStringify)
    .process(body)

  if (post) {
    return {
      post: post, renderedBody: String(file),
    };
  }

  return {
    message: "Not found", post: null,
  };
}
