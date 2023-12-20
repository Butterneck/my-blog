import { getCurrentUser } from "./auth";
import { DefaultService } from "./generated/backend-client";

export async function getPosts({pageSize, nextPageToken}: {pageSize?: number, nextPageToken?: string}): Promise<{nextPageToken?: string, posts: Post[]}> {
    if (await getCurrentUser()) {
        return await DefaultService.getAllPosts({pageSize: pageSize, nextPageToken: nextPageToken});
    }

    return await DefaultService.getPublishedPosts({});
}

export async function getPost(slug: string): Promise<Post> {
    const currentUser = await getCurrentUser();
    if (currentUser) {
        return await DefaultService.getAnyPost({
            slug: slug,
        });
    }

    return await DefaultService.getPublishedPost({
        slug: slug,
    });
}

export async function createPost(post: NewPost): Promise<any> {
    return await DefaultService.createPost({
        requestBody: post,
    });
}

export async function updatePost(post: Post): Promise<any> {
    return await DefaultService.updatePost({
        slug: post.slug,
        requestBody: {
            title: post.title,
            body: post.body,
        },
    });
}

export async function deletePost(slug: string): Promise<any> {
    return await DefaultService.deletePost({
        slug: slug,
    });
}

export async function publishPost(slug: string): Promise<any> {
    return await DefaultService.publishPost({
        slug: slug,
    });
}