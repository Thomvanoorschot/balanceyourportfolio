import type { RequestEvent, RequestHandler } from './$types';
import { error, json } from '@sveltejs/kit';
import { safe } from '$lib/server/safe.ts';
import { portfolioClient } from '$lib/server/grpc.ts';
import type { PortfoliosResponse__Output } from '$lib/proto/proto/PortfoliosResponse.ts';
import type { PortfoliosRequest__Output } from '$lib/proto/proto/PortfoliosRequest.ts';
import type { UpdatePortfolioFundAmountRequest__Output } from '$lib/proto/proto/UpdatePortfolioFundAmountRequest.ts';
import type { Empty__Output } from '$lib/proto/proto/Empty.ts';
import { kindeAuthClient, type SessionManager } from '@kinde-oss/kinde-auth-sveltekit';

export const GET: RequestHandler = async ({ request }: RequestEvent) => {
	const isAuthenticated = await kindeAuthClient.isAuthenticated(
		request as unknown as SessionManager
	);
	if (!isAuthenticated) {
		return json('');
	}
	const user = await kindeAuthClient.getUser(request as unknown as SessionManager);
	const req: PortfoliosRequest__Output = {
		userId: user.id
	};
	const resp = await safe(
		new Promise<PortfoliosResponse__Output>((resolve, reject) => {
			portfolioClient.getPortfolios(req, (err, response) => {
				return err || !response ? reject(err) : resolve(response);
			});
		})
	);
	if (!resp.success) {
		error(400, 'could not fetch portfolios');
		return json('');
	}

	return json(resp.data.entries);
};
export interface PatchValueRequest {
	portfolioId: string;
	fundId: string;
	amount: number;
}
export const PATCH: RequestHandler = async ({ request }: RequestEvent) => {
	const body: PatchValueRequest = await request.json();
	const req: UpdatePortfolioFundAmountRequest__Output = {
		portfolioId: body.portfolioId,
		fundId: body.fundId,
		amount: body.amount
	};
	const resp = await safe(
		new Promise<Empty__Output>((resolve, reject) => {
			portfolioClient.updatePortfolioFundAmount(req, (err, response) => {
				return err || !response ? reject(err) : resolve(response);
			});
		})
	);
	if (!resp.success) {
		error(400, 'could not fetch portfolios');
		return json('');
	}
	return json('');
};
