import { AuthApi } from "@/generated/swagger"

const endpoint = process.env.VUE_APP_GATEWAY_ENDPOINT;

const authApi = new AuthApi(undefined, endpoint)

export const AuthAPI = authApi;


