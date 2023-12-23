
export function isPostPublised(post: Post): boolean {
    return post.creationDate !== 0;
}