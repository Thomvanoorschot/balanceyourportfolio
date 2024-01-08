import { kindeAuthClient, type SessionManager } from '@kinde-oss/kinde-auth-sveltekit';
import type { RequestEvent } from '@sveltejs/kit';
import UAParser from 'ua-parser-js';

export async function load({ request }: RequestEvent) {
	let profilePictureURL: string | null = '';
	const isAuthenticated = await kindeAuthClient.isAuthenticated(
		request as unknown as SessionManager
	);
	if (isAuthenticated) {
		const user = await kindeAuthClient.getUser(request as unknown as SessionManager);
		profilePictureURL = user.picture;
	}
	const userAgent = new UAParser(request.headers.get("user-agent") || "")
	const isMobile = ["mobile", "wearable"].includes(userAgent.getDevice().type || "");
	return {
		profilePictureURL,
		isAuthenticated,
		isMobile
	};
}
