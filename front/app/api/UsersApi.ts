import {type User} from "../types/User";
import {type ApiResponse, BaseApi} from "./BaseApi";

export class UsersApi extends BaseApi {

    public static async getAll(): Promise<User[]> {
        const response: ApiResponse<User[]> | null = await this.get<User[]>('/users');

        if (response === null) {
            return [];
        }

        return response.data;
    }

    public static async getById(userId: number): Promise<User | null> {
        return null;
    }

    public static async create(user: User): Promise<void> {
        return this.post('/users/create', {}, user).then(() => {
        });
    }

    public static async update(user: User): Promise<void> {

    }

    public static async remove(id: string): Promise<void> {
        return this.del(`/users/${id}`).then(res => {
            console.log(res);
        });
    }
}