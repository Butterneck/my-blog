/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { AdminPost } from '../models/AdminPost';
import type { NewPostRequest } from '../models/NewPostRequest';
import type { Post } from '../models/Post';
import type { UpdatePostRequest } from '../models/UpdatePostRequest';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class DefaultService {

    /**
     * Retrieve a list of published posts
     * @returns any A list of published posts
     * @throws ApiError
     */
    public static getPublishedPosts({
        pageSize,
        nextPageToken,
    }: {
        /**
         * The numbers of items to return
         */
        pageSize?: number,
        /**
         * The page token to use for pagination
         */
        nextPageToken?: string,
    }): CancelablePromise<{
        /**
         * The next page token to use for pagination
         */
        nextPageToken?: string;
        /**
         * The list of posts
         */
        posts: Array<Post>;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/posts',
            query: {
                'pageSize': pageSize,
                'nextPageToken': nextPageToken,
            },
            errors: {
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Retrieve a list of posts
     * @returns any A list of posts
     * @throws ApiError
     */
    public static getAllPosts({
        pageSize,
        nextPageToken,
    }: {
        /**
         * The numbers of items to return
         */
        pageSize?: number,
        /**
         * The page token to use for pagination
         */
        nextPageToken?: string,
    }): CancelablePromise<{
        /**
         * The next page token to use for pagination
         */
        nextPageToken?: string;
        /**
         * The list of posts
         */
        posts: Array<AdminPost>;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/admin/posts',
            query: {
                'pageSize': pageSize,
                'nextPageToken': nextPageToken,
            },
            errors: {
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Create a new post
     * @returns any Post successfully created
     * @throws ApiError
     */
    public static createPost({
        formData,
    }: {
        formData: NewPostRequest,
    }): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/admin/posts',
            formData: formData,
            mediaType: 'multipart/form-data',
            errors: {
                400: `Bad Request`,
                500: `Internal Server`,
            },
        });
    }

    /**
     * Retrieve a published post
     * @returns Post The requested post
     * @throws ApiError
     */
    public static getPublishedPost({
        slug,
    }: {
        slug: string,
    }): CancelablePromise<Post> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/posts/{slug}',
            path: {
                'slug': slug,
            },
            errors: {
                404: `Post not found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Retrieve a post
     * @returns AdminPost The requested post
     * @throws ApiError
     */
    public static getAnyPost({
        slug,
    }: {
        slug: string,
    }): CancelablePromise<AdminPost> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/admin/posts/{slug}',
            path: {
                'slug': slug,
            },
            errors: {
                404: `Post not found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update a post
     * @returns any Post successfully updated
     * @throws ApiError
     */
    public static updatePost({
        slug,
        formData,
    }: {
        slug: string,
        formData: UpdatePostRequest,
    }): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/v1/admin/posts/{slug}',
            path: {
                'slug': slug,
            },
            formData: formData,
            mediaType: 'multipart/form-data',
            errors: {
                400: `Bad Request`,
                404: `Post not found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Delete a post
     * @returns void
     * @throws ApiError
     */
    public static deletePost({
        slug,
    }: {
        slug: string,
    }): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/v1/admin/posts/{slug}',
            path: {
                'slug': slug,
            },
            errors: {
                404: `Post not found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Publish the draft of a post
     * @returns any Post successfully created
     * @throws ApiError
     */
    public static publishPost({
        slug,
    }: {
        slug: string,
    }): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/admin/posts/{slug}/publish',
            path: {
                'slug': slug,
            },
            errors: {
                400: `Bad Request`,
                500: `Internal Server`,
            },
        });
    }

}
