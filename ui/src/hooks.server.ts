import { sessionHooks, type Handler } from '@kinde-oss/kinde-auth-sveltekit';

export const handle: Handler = async ({ event, resolve }) => {
    await sessionHooks({ event });
    const response = await resolve(event);
    return response;
}
export function getSession(request:any) {
    console.log("AAAAAAAa")
    return {
        mobile: request.headers['sec-ch-ua-mobile'] === '?1'
    }
}