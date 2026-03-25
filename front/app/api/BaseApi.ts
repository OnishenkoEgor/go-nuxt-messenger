import {useFetch} from "nuxt/app";
import type {FetchError} from "ofetch";
import type {AsyncData} from 'nuxt/app'

const ERROR_STATUS = 'error';

export interface ApiResponse<T = null> {
    errors: boolean,
    data: T
}

export abstract class BaseApi {
    private static baseUrl: string = 'http://localhost:8080/api';
    private static corsHeaders: object = {
        "Access-Control-Allow-Methods": 'GET,HEAD,PATCH,POST,DELETE',
        "Access-Control-Allow-Origin": '*',
        "Access-Control-Allow-Credentials": "true",
    };

    private static contentTypeHeaders: object = {
        "Content-Type": "application/json"
    };

    protected static async get<T>(route: string, params: object = {}): Promise<ApiResponse<T> | null> {
        const response = await $fetch<ApiResponse<T>>(this.prepareUrl(route, params), {
            method: 'GET',
            headers: {
                ...this.corsHeaders,
                ...this.contentTypeHeaders
            }
        }).catch(err => err);

        return this.prepareResponse(response);
    }

    protected static async post<T>(route: string, params: object = {}, body: object = {}): Promise<ApiResponse<T> | null> {
        const response = await $fetch<ApiResponse<T>>(this.prepareUrl(route, params), {
            method: 'POST',
            headers: {
                ...this.corsHeaders,
                ...this.contentTypeHeaders
            },
            body: JSON.stringify(body)
        }).catch(err => err);

        return this.prepareResponse(response);
    }

    protected static async patch(): Promise<ApiResponse> {
        return new Promise(() => null);
    }

    protected static async put<T>(route: string, params: object = {}, body: object = {}): Promise<ApiResponse<T> | null> {
        const response = await $fetch<ApiResponse<T>>(this.prepareUrl(route, params), {
            method: 'PUT',
            headers: {
                ...this.corsHeaders,
                ...this.contentTypeHeaders
            },
            body: JSON.stringify(body)
        }).catch(err => err);

        return this.prepareResponse(response);
    }

    protected static async del<T>(route: string): Promise<ApiResponse<T> | null> {
        const response = await $fetch<ApiResponse<T>>(this.prepareUrl(route), {
            method: 'DELETE',
            headers: {
                ...this.corsHeaders,
                ...this.contentTypeHeaders
            },
        }).catch(err => err);

        return this.prepareResponse(response);
    }

    private static prepareUrl(route: string, params: object = {}): string {
        //TODO add query params
        return this.baseUrl + route;
    }

    private static prepareResponse<T>(response: ApiResponse<T> | undefined): ApiResponse<T> | null {
        if (response === undefined) {
            return null;
        }

        return response;
    }
}