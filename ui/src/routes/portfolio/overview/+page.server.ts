import type {Actions, PageServerLoad} from './$types';
import {fundClient, portfolioClient} from "$lib/server/grpc";
import {safe} from "$lib/server/safe";
import {fail} from "@sveltejs/kit";
import type {PortfoliosResponse__Output} from "$lib/proto/proto/PortfoliosResponse";
import type {UpsertPortfolioResponse__Output} from "$lib/proto/proto/UpsertPortfolioResponse";
import type {UpsertPortfolioRequest__Output} from "$lib/proto/proto/UpsertPortfolioRequest";
import type {SearchFundsRequest__Output} from "$lib/proto/proto/SearchFundsRequest";
import type {SearchFundsResponse__Output} from "$lib/proto/proto/SearchFundsResponse";
import type {Portfolio__Output} from "$lib/proto/proto/Portfolio";

export const load = (async ({}) => {
    const portfoliosResp = await safe(
        new Promise<PortfoliosResponse__Output>((resolve, reject) => {
            portfolioClient.GetPortfolios({}, (err, response) => {
                return err || !response ? reject(err) : resolve(response);
            })
        }),
    );
    if (!portfoliosResp.success) {
        return fail(500, {error: "could not fetch portfolios"});
    }
    return {
        portfolios: portfoliosResp.data,
    };
}) satisfies PageServerLoad;


export const actions = {
    upsertPortfolio: async ({request, url,}) => {
        const formData = await request.formData()
        const portfolio: Portfolio__Output = JSON.parse(String(formData.get("portfolio")))
        portfolio.entries = portfolio.entries.filter(x => x.amount > 0)
        const req: UpsertPortfolioRequest__Output = {
            portfolio: portfolio,
        }
        const portfolioResp = await safe(
            new Promise<UpsertPortfolioResponse__Output>((resolve, reject) => {
                portfolioClient.upsertPortfolio(req, (err, response) => {
                    return err || !response ? reject(err) : resolve(response);
                })
            }),
        );
        if (!portfolioResp.success) {
            return fail(500, {error: "could not upsert portfolio"});
        }
        return {
            portfolio: portfolioResp.data.portfolio
        };
    },
    searchFunds: async ({request}) => {
        const formData = await request.formData()
        const searchTerm = String(formData.get("searchTerm"))
        if (searchTerm === "") {
            return {
                funds: []
            }
        }
        const req: SearchFundsRequest__Output = {searchTerm: searchTerm, }
        const resp = await safe(
            new Promise<SearchFundsResponse__Output>((resolve, reject) => {
                fundClient.searchFunds(req, (err, response) => {
                    return err || !response ? reject(err) : resolve(response);
                })
            }),
        );
        if (!resp.success) {
            return fail(500, {error: "could not search funds"});
        }
        return {
            funds: resp.data.entries
        }
    }
} satisfies Actions;