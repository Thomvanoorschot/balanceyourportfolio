import {fundClient} from "$lib/server/grpc.ts";
import {safe} from "$lib/server/safe.ts";
import {fail} from "@sveltejs/kit";
import type {Actions, PageServerLoad} from './$types';
import type {Holding} from "$lib/holding.ts";
import type {FundDetailsResponse__Output} from "$lib/proto/proto/FundDetailsResponse.ts";
import type {FundDetailsRequest__Output} from "$lib/proto/proto/FundDetailsRequest.ts";
import type {FilterFundHoldingsRequest__Output} from "$lib/proto/proto/FilterFundHoldingsRequest.ts";
import type {FilterFundHoldingsResponse__Output} from "$lib/proto/proto/FilterFundHoldingsResponse.ts";


export const load = (async ({fetch, params, url, route}) => {
    const fundId = params.id
    const detailsReq: FundDetailsRequest__Output = {fundId: String(fundId)}
    const detailsResp = await safe(
        new Promise<FundDetailsResponse__Output>((resolve, reject) => {
            fundClient.getDetails(detailsReq, (err, response) => {
                return err || !response ? reject(err) : resolve(response);
            })
        }),
    );
    if (!detailsResp.success) {
        return fail(500, {error: "could not get fund details"});
    }

    // sectorWeightings = sectorWeightings.sort((x, y) => y.weighting.totalPercentage - x.weighting.totalPercentage)
    const holdings: Holding[] = detailsResp.data.fundHoldings.map(x => (
        {
            id: x.holdingId,
            name: x.holdingName,
            ticker: x.ticker,
            percentage: x.cumulativePercentage,
            funds:[]
        }
    ))
    return {
        sectors: detailsResp.data.sectors,
        fundInformation: detailsResp.data.information,
        fundSectorWeightings: detailsResp.data.sectorWeightings,
        holdings: holdings
    }

}) satisfies PageServerLoad;


export const actions = {
    filterHoldings: async ({request, url, params,}) => {
        const formData = await request.formData()
        const holdingsLength = Number(formData.get("holdingsLength") || 0)
        const fundId = String(params.id || "")
        const searchTerm = String(formData.get("searchTerm") || "")
        const selectedSectors = JSON.parse(String(formData.get("selectedSectors")))
        const holdingsReq: FilterFundHoldingsRequest__Output = {
            fundId: fundId,
            limit: 20,
            offset: holdingsLength,
            searchTerm: searchTerm,
            selectedSectors: selectedSectors
        }
        const holdingsResp = await safe(
            new Promise<FilterFundHoldingsResponse__Output>((resolve, reject) => {
                fundClient.filterHoldings(holdingsReq, (err, response) => {
                    return err || !response ? reject(err) : resolve(response);
                })
            }),
        );
        if (!holdingsResp.success) {
            return fail(500, {error: "could not filter holdings"});
        }
        const holdings: Holding[] = holdingsResp.data.entries.map(x => (
            {
                id: x.holdingId,
                name: x.holdingName,
                ticker: x.ticker,
                percentage: x.cumulativePercentage,
                funds: []
            }
        ))
        return {
            holdings: holdings
        };
    }
} satisfies Actions;