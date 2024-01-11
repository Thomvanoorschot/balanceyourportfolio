import { sessionHooks, type Handler } from '@kinde-oss/kinde-auth-sveltekit';

export const handle: Handler = async ({ event, resolve }) => {
	await sessionHooks({ event });
	return resolve(event);
};