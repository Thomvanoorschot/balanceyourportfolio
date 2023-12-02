import {portfolioClient} from "$lib/server/grpc";
import {safe} from "$lib/server/safe";
import {fail} from "@sveltejs/kit";
import type {PageServerLoad} from './$types';
import type {PortfolioDetailsRequest__Output} from "$lib/proto/proto/PortfolioDetailsRequest";
import type {PortfolioDetailsResponse__Output} from "$lib/proto/proto/PortfolioDetailsResponse";
import {colors, stringToRandomInteger} from "$lib/utils";

export const load = (async ({fetch, params, url, route}) => {
    const portfolioId =params.id
    const detailsReq: PortfolioDetailsRequest__Output = {portfolioId: String(portfolioId)}
    const detailsResp = await safe(
        new Promise<PortfolioDetailsResponse__Output>((resolve, reject) => {
            portfolioClient.getPortfolioDetails(detailsReq, (err, response) => {
                return err || !response ? reject(err) : resolve(response);
            })
        }),
    );
    if (!detailsResp.success) {
        return fail(500, {error: "could not search funds"});
    }
    let colorsCopy = [...colors]
    const colorMap = new Map<string, string>()
    detailsResp.data?.portfolioFundSectorWeightings.forEach((x) => {
        const color = colorMap.get(x.fundName)
        if (!color) {
            const randomColor = colorsCopy[stringToRandomInteger(x.fundName, colorsCopy.length)]
            colorMap.set(x.fundName, randomColor)
            colorsCopy = colorsCopy.filter(x => x != randomColor)
        }
    });
    return {
        details: detailsResp.data,
        colorMap: colorMap
    };

}) satisfies PageServerLoad;