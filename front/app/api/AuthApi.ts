import {BaseApi} from "./BaseApi";

export type AuthResponse = {
    token: string
}

export type AuthRequest = {
    login: string,
    password: string
}
export class AuthApi extends BaseApi{
    public login(){

    }

    public register(){

    }
}