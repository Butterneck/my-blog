import { getCurrentUser } from "./auth";
import { DefaultService } from "./generated/backend-client";

export async function getPosts(): Promise<Post[]> {
    if (await getCurrentUser()) {
        return await DefaultService.getAllPosts();
    }

    return await DefaultService.getPublishedPosts();
}

export async function getPost(slug: string): Promise<Post> {
    if (await getCurrentUser()) {
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