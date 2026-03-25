import {type User} from "../types/User";
import {type ApiResponse, BaseApi} from "./BaseApi";

export class UsersApi extends BaseApi {

    public static async getAll(): Promise<User[]> {
        const response: ApiResponse<User[]> | null = await this.get<User[]>('/users');

        if (response === null) {
            throw new Error('Failed get users.');
        }

        if (response.errors) {
            throw new Error(`Error on get users: ${response.data}`);
        }

        return response.data;
    }

    public static async getById(userId: number): Promise<User> {
        const response: ApiResponse<User> | null = await this.get(`/users/${userId}`, {});

        if (response === null) {
            throw new Error('Failed get user by id.');
        }

        if (response.errors) {
            throw new Error(`Error on get user by id: ${response.data}`);
        }

        return response.data;
    }

    public static async create(user: User): Promise<void> {
        const response: ApiResponse | null = await this.post('/users/create', {}, user);

        if (response === null) {
            throw new Error('Failed create user request.');
        }

        if (response.errors) {
            throw new Error(`Error on create: ${response.data}`);
        }
    }

    public static async update(id: number, user: User): Promise<void> {
        const response: ApiResponse | null = await this.put(`/users/${id}`, {}, user);

        if (response === null) {
            throw new Error('Failed update user request');
        }

        if (response.errors) {
            throw new Error(`Error on update: ${response.data}`);
        }
    }

    public static async delById(id: number): Promise<void> {
        const response: ApiResponse | null = await this.del(`/users/${id}`);

        if (response === null) {
            throw new Error('Failed delete user request.');
        }

        if (response.errors) {
            throw new Error(`Error on delete: ${response.data}`);
        }
    }
}