import {useFetch} from "nuxt/app";

export interface ApiResponse<T = null> {
    errors: string,
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

    public static async get<T>(route: string, params: object = {}): Promise<ApiResponse<T> | null> {
        //TODO add query params
        const {data, status, error} = await useFetch<ApiResponse<T>>(this.prepareUrl(route, params), {
            method: 'GET',
            headers: {
                ...this.corsHeaders,
                ...this.contentTypeHeaders
            }
        });

        if (!data.value) {
            return null;
        }

        return data.value;
    }

    public static async post<T>(route: string, params: object = {}, body: object = {}): Promise<ApiResponse<T> | null> {
        const {data} = await useFetch<ApiResponse<T>>(this.prepareUrl(route, params),{
            method: 'POST',
            headers: {
                ...this.corsHeaders,
                ...this.contentTypeHeaders
            },
            body: JSON.stringify(body)
        });

        if (!data.value) {
            return null;
        }

        return data.value;
    }

    public static async patch(): Promise<ApiResponse> {
        return new Promise(() => null);
    }

    public static async put(): Promise<ApiResponse> {
        return new Promise(() => null);
    }

    public static async del<T>(route:string): Promise<ApiResponse<T> | null> {
        const {data} = await useFetch<ApiResponse<T>>(this.prepareUrl(route), {
            method: 'DELETE',
            headers: {
                ...this.corsHeaders,
                ...this.contentTypeHeaders
            },
        });

        if (!data.value) {
            return null;
        }

        return data.value;
    }

    private static prepareUrl(route: string, params: object = {}): string {
        //TODO add query params
        return this.baseUrl + route;
    }
}