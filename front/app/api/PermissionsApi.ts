import {type ApiResponse, BaseApi} from "./BaseApi";
import {type Permission} from "../types/Permission";
import type {User} from "../types/User";

export class PermissionsApi extends BaseApi {

    public static async getAll(): Promise<Permission[]> {
        const response: ApiResponse<Permission[]> | null = await this.get<Permission[]>('/qwe');

        if (response === null) {
            return [];
        }

        return response.data
    }

    public static async create(permission: Permission): Promise<void> {
        const response: ApiResponse<void> | null = await this.post<void>('/');
    }

    public static async remove(id: string): Promise<void> {
        const response: ApiResponse<void> | null = await this.del('/');
    }
}