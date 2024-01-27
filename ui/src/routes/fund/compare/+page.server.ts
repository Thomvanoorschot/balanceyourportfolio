import type { Actions, PageServerLoad } from './$types';
import { safe } from '$lib/server/safe.ts';
import { fundClient } from '$lib/server/grpc.ts';
import { fail } from '@sveltejs/kit';
import type { CompareFundRequest__Output } from '$lib/proto/proto/CompareFundRequest.ts';
import type { CompareFundResponse__Output } from '$lib/proto/proto/CompareFundResponse.ts';


export const load = (async ({ fetch, params, url, route }) => {
	const fundOne = url.searchParams.get('fundOne');
	const fundTwo = url.searchParams.get('fundTwo');
	if (!fundOne || !fundTwo) {
		return {};
	}

	const compareReq: CompareFundRequest__Output = {
		fundOne: fundOne,
		fundTwo: fundTwo
	};
	const compareResp = await safe(
		new Promise<CompareFundResponse__Output>((resolve, reject) => {
			fundClient.compareFunds(compareReq, (err, response) => {
				return err || !response ? reject(err) : resolve(response);
			});
		})
	);
	if (!compareResp.success) {
		return fail(500, { error: 'could not compare funds' });
	}
	const colorMap = new Map<string, { fundName: string; color: string }>();
	colorMap.set(fundOne, { fundName: compareResp.data.fundOneName, color: '#f582ae' });
	colorMap.set(fundTwo, { fundName: compareResp.data.fundTwoName, color:  '#008080' });
	return {
		comparison: compareResp.data,
		colorMap: colorMap
	};
}) satisfies PageServerLoad;

export const actions = {
	compareFunds: async ({ request, url, params }) => {
		const formData = await request.formData();
		const fundOne = String(formData.get('fundOne') || '');
		const fundTwo = String(formData.get('fundTwo') || '');

		if (!fundOne || !fundTwo) {
			return fail(500, { error: 'need both funds to compare' });
		}

		const compareReq: CompareFundRequest__Output = {
			fundOne: fundOne,
			fundTwo: fundTwo
		};
		const compareResp = await safe(
			new Promise<CompareFundResponse__Output>((resolve, reject) => {
				fundClient.compareFunds(compareReq, (err, response) => {
					return err || !response ? reject(err) : resolve(response);
				});
			})
		);
		if (!compareResp.success) {
			return fail(500, { error: 'could not compare funds' });
		}
		const colorMap = new Map<string, { fundName: string; color: string }>();
		colorMap.set(fundOne, { fundName: compareResp.data.fundOneName, color: '#f582ae' });
		colorMap.set(fundTwo, { fundName: compareResp.data.fundTwoName, color:  '#008080' });
		return {
			comparison: compareResp.data,
			colorMap: colorMap
		};
	}
} satisfies Actions;
