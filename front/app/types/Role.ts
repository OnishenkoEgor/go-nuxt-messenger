import {type Permission} from "./Permission";

export interface Role {
    id?: string,
    name: string,
    permissions: Permission[]
}