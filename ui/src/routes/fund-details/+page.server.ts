import type {Actions, PageServerLoad} from './$types';
import {fundClient} from "$lib/server/grpc";
import {safe} from "$lib/server/safe";
import {fail} from "@sveltejs/kit";
import type {FundDetailsResponse__Output} from "$lib/proto/proto/FundDetailsResponse";
import type {FilterHoldingsRequest, FilterHoldingsRequest__Output} from "$lib/proto/proto/FilterHoldingsRequest";
import type {HoldingsListResponse__Output} from "$lib/proto/proto/HoldingsListResponse";
import type {GetFundDetailsRequest__Output} from "$lib/proto/proto/GetFundDetailsRequest";

export const load = (async ({fetch, params, url}) => {
    const fundId = url.searchParams.get("fundId")
    const detailsReq: GetFundDetailsRequest__Output = {fundId: fundId || ""}
    const detailsResp = await safe(
        new Promise<FundDetailsResponse__Output>((resolve, reject) => {
            fundClient.getDetails(detailsReq, (err, response) => {
                return err || !response ? reject(err) : resolve(response);
            })
        }),
    );
    if (!detailsResp.success) {
        return fail(500, {error: "could not search funds"});
    }

    const holdingsReq: FilterHoldingsRequest = {fundId: fundId || "", limit: 20}
    const holdingsResp = await safe(
        new Promise<HoldingsListResponse__Output>((resolve, reject) => {
            fundClient.filterHoldings(holdingsReq, (err, response) => {
                return err || !response ? reject(err) : resolve(response);
            })
        }),
    );
    if (!holdingsResp.success) {
        return fail(500, {error: "could not filter holdings"});
    }
    return {
        details: detailsResp.data,
        holdings: holdingsResp.data
    };

}) satisfies PageServerLoad;

export const actions = {
    filterHoldings: async ({request, url,}) => {
        const formData = await request.formData()
        const holdingsLength = Number(formData.get("holdingsLength") || 0)
        const fundId = String(formData.get("fundId") || "")
        const searchTerm = String(formData.get("searchTerm") || "")
        const sectorName = String(formData.get("sectorName") || "")
        const holdingsReq: FilterHoldingsRequest__Output = {
            fundId: fundId,
            limit: 20,
            offset: holdingsLength,
            searchTerm: searchTerm,
            sectorName: sectorName
        }

        const holdingsResp = await safe(
            new Promise<HoldingsListResponse__Output>((resolve, reject) => {
                fundClient.filterHoldings(holdingsReq, (err, response) => {
                    return err || !response ? reject(err) : resolve(response);
                })
            }),
        );
        if (!holdingsResp.success) {
            return fail(500, {error: "could not filter holdings"});
        }
        return {
            holdings: holdingsResp.data
        };
    }
} satisfies Actions;