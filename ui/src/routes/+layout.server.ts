import {kindeAuthClient, type SessionManager} from '@kinde-oss/kinde-auth-sveltekit';
import type {RequestEvent} from '@sveltejs/kit';

export async function load({request}: RequestEvent) {
    let profilePictureURL: string | null = ""
    const isAuthenticated = await kindeAuthClient.isAuthenticated(
        request as unknown as SessionManager
    );
    if (isAuthenticated) {
        const user = await kindeAuthClient.getUser(request as unknown as SessionManager)
        profilePictureURL = user.picture
    }

    return {
        profilePictureURL,
        isAuthenticated
    };
}