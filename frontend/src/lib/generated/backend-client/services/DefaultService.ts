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
     * @returns Post A list of published posts
     * @throws ApiError
     */
    public static getApiV1Posts(): CancelablePromise<Array<Post>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/posts',
            errors: {
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Retrieve a list of posts
     * @returns AdminPost A list of posts
     * @throws ApiError
     */
    public static getApiV1AdminPosts(): CancelablePromise<Array<AdminPost>> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/admin/posts',
            errors: {
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Create a new post
     * @param requestBody
     * @returns any Post successfully created
     * @throws ApiError
     */
    public static postApiV1AdminPosts(
        requestBody: NewPostRequest,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/admin/posts',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                500: `Internal Server`,
            },
        });
    }

    /**
     * Retrieve a published post
     * @param postSlug
     * @returns Post The requested post
     * @throws ApiError
     */
    public static getApiV1Posts1(
        postSlug: string,
    ): CancelablePromise<Post> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/posts/{postSlug}',
            path: {
                'postSlug': postSlug,
            },
            errors: {
                404: `Post not found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Retrieve a post
     * @param postSlug
     * @returns AdminPost The requested post
     * @throws ApiError
     */
    public static getApiV1AdminPosts1(
        postSlug: string,
    ): CancelablePromise<AdminPost> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/api/v1/admin/posts/{postSlug}',
            path: {
                'postSlug': postSlug,
            },
            errors: {
                404: `Post not found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Update a post
     * @param postSlug
     * @param requestBody
     * @returns any Post successfully updated
     * @throws ApiError
     */
    public static putApiV1AdminPosts(
        postSlug: string,
        requestBody: UpdatePostRequest,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/api/v1/admin/posts/{postSlug}',
            path: {
                'postSlug': postSlug,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                404: `Post not found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Delete a post
     * @param postSlug
     * @returns void
     * @throws ApiError
     */
    public static deleteApiV1AdminPosts(
        postSlug: string,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/api/v1/admin/posts/{postSlug}',
            path: {
                'postSlug': postSlug,
            },
            errors: {
                404: `Post not found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Publish the draft of a post
     * @param postSlug
     * @returns any Post successfully created
     * @throws ApiError
     */
    public static postApiV1AdminPostsPublish(
        postSlug: string,
    ): CancelablePromise<any> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/admin/posts/{postSlug}/publish',
            path: {
                'postSlug': postSlug,
            },
            errors: {
                400: `Bad Request`,
                500: `Internal Server`,
            },
        });
    }

}
