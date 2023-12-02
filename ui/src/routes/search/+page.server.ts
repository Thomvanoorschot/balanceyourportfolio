import type {PageServerLoad} from './$types';
import type {SearchFundsRequest} from "$lib/proto/proto/SearchFundsRequest";
import {fundClient} from "$lib/server/grpc";
import {safe} from "$lib/server/safe";
import type {SearchFundsResponse} from "$lib/proto/proto/SearchFundsResponse";
import {fail} from "@sveltejs/kit";

export const load = (async ({fetch, params, url}) => {
    const searchTerm =url.searchParams.get("searchTerm") || "";
    if (searchTerm === ""){
        return {
            funds: []
        }
    }
    const req: SearchFundsRequest = {searchTerm:  searchTerm.toString()}
    const resp = await safe(
        new Promise<SearchFundsResponse>((resolve, reject) => {
            fundClient.searchFunds(req, (err, response) => {
                return err || !response ? reject(err) : resolve(response);
            })
        }),
    );
    if (!resp.success) {
        return fail(500, { error: "could not search funds" });
    }
    return {
        funds: resp.data.entries
    }
}) satisfies PageServerLoad;