import type { Actions, PageServerLoad } from './$types';
import { error, fail, json } from '@sveltejs/kit';
import { safe } from '$lib/server/safe.ts';
import { fundClient } from '$lib/server/grpc.ts';
import type { FilterFundsResponse__Output } from '$lib/proto/proto/FilterFundsResponse.ts';
import type { FilterFundsRequest__Output } from '$lib/proto/proto/FilterFundsRequest.ts';

export const load = (async ({}) => {
	// const body = await request.json()
	// const searchTerm = String(body.value || "")

	// if (searchTerm === ""){
	//     return json([])
	// }
	const req: FilterFundsRequest__Output = { searchTerm: '', providers: [], limit: 20, offset: 0 };
	const resp = await safe(
		new Promise<FilterFundsResponse__Output>((resolve, reject) => {
			fundClient.filterFunds(req, (err, response) => {
				return err || !response ? reject(err) : resolve(response);
			});
		})
	);
	if (!resp.success) {
		error(400, 'could not fetch funds');
		return json('');
	}

	return {
		funds: resp.data.entries
	};
}) satisfies PageServerLoad;

export const actions = {
	filterFunds: async ({ request, url, params }) => {
		const formData = await request.formData();
		const fundsLength = Number(formData.get('fundsLength') || 0);
		const searchTerm = String(formData.get('searchTerm') || '');
		const selectedProviders = JSON.parse(String(formData.get('selectedProviders')));

		const filterReq: FilterFundsRequest__Output = {
			limit: 20,
			offset: fundsLength,
			searchTerm: searchTerm,
			providers: selectedProviders
		};
		const fundsResp = await safe(
			new Promise<FilterFundsResponse__Output>((resolve, reject) => {
				fundClient.filterFunds(filterReq, (err, response) => {
					return err || !response ? reject(err) : resolve(response);
				});
			})
		);
		if (!fundsResp.success) {
			return fail(500, { error: 'could not filter holdings' });
		}
		return {
			funds: fundsResp.data.entries
		};
	}
} satisfies Actions;
