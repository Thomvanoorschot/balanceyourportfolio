import { redirect } from '@sveltejs/kit';

export const load = (async () => {
	throw redirect(302, 'fund/overview'); // needs `throw` in v1
});