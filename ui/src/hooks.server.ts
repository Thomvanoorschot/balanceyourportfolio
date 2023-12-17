import type {Handle} from '@sveltejs/kit'

export const handle: Handle = async ({event, resolve}) => {
    const response = await resolve(event, {
        transformPageChunk: ({html}) => {
            const cookie = event.request.headers.get("cookie")
            const theme = cookie?.includes("theme=")
            const rgx = /theme=([a-zA-Z]*)/
            if (cookie) {
                const arr = rgx.exec(cookie);
                if (arr && arr[1]) {

                    return  html.replace(`<html lang="en"`, `<html lang="en" class="${arr[1]}"`)
                }
            }
            return html
        },
        filterSerializedResponseHeaders: (name) => name.startsWith('x-'),
        preload: ({type, path}) => type === 'js' || path.includes('/important/')
    });

    return response;
}