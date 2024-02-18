import type { RequestEvent, RequestHandler } from './$types';
import { fundClient } from '$lib/server/grpc.ts';
import type { SearchFundsRequest } from '$lib/proto/proto/SearchFundsRequest.ts';
import type { SearchFundsResponse } from '$lib/proto/proto/SearchFundsResponse.ts';
import { error, json } from '@sveltejs/kit';
import { safe } from '$lib/server/safe.ts';

export const POST: RequestHandler = async ({ request }: RequestEvent) => {
	const body = await request.json();
	const searchTerm = String(body.value || '');

	if (searchTerm === '') {
		return json([]);
	}
	const req: SearchFundsRequest = { searchTerm: searchTerm.toString() };
	const resp = await safe(
		new Promise<SearchFundsResponse>((resolve, reject) => {
			fundClient.searchFunds(req, (err, response) => {
				return err || !response ? reject(err) : resolve(response);
			});
		})
	);
	if (!resp.success) {
		error(400, 'could not search funds');
		return json('');
	}

	return json(resp.data.entries);
};
