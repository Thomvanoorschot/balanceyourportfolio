import { sessionHooks, type Handler } from '@kinde-oss/kinde-auth-sveltekit';

export const handle: Handler = async ({ event, resolve }) => {
	console.log(event.request.headers)
	await sessionHooks({ event });
	return resolve(event);
};