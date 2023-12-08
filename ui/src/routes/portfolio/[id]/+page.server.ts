import {portfolioClient} from "$lib/server/grpc";
import {safe} from "$lib/server/safe";
import {fail} from "@sveltejs/kit";
import type {PageServerLoad} from './$types';
import type {PortfolioDetailsRequest__Output} from "$lib/proto/proto/PortfolioDetailsRequest";
import type {PortfolioDetailsResponse__Output} from "$lib/proto/proto/PortfolioDetailsResponse";
import {colors, stringToRandomInteger} from "$lib/utils";
import type {PortfolioFundSectorWeighting__Output} from "$lib/proto/proto/PortfolioFundSectorWeighting";
import type {PortfolioSectorWeighting} from "$lib/portfolio";


export const load = (async ({fetch, params, url, route}) => {
    const portfolioId = params.id
    const detailsReq: PortfolioDetailsRequest__Output = {portfolioId: String(portfolioId)}
    const detailsResp = await safe(
        new Promise<PortfolioDetailsResponse__Output>((resolve, reject) => {
            portfolioClient.getPortfolioDetails(detailsReq, (err, response) => {
                return err || !response ? reject(err) : resolve(response);
            })
        }),
    );
    if (!detailsResp.success) {
        return fail(500, {error: "could not get portfolio details"});
    }
    let colorsCopy = [...colors]
    const colorMap = new Map<string, string>()
    let sectorWeightings: PortfolioSectorWeighting[] = []
    Object.keys(detailsResp.data?.portfolioFundSectorWeightings).forEach((key) => {
        const value: PortfolioFundSectorWeighting__Output = detailsResp.data?.portfolioFundSectorWeightings[key];
        sectorWeightings.push({sectorName: key, weighting: value})
        value.fundSectorWeighting.forEach(x => {
            const color = colorMap.get(x.fundName)
            if (!color) {
                const randomColor = colorsCopy[stringToRandomInteger(x.fundName, colorsCopy.length)]
                colorMap.set(x.fundName, randomColor)
                colorsCopy = colorsCopy.filter(x => x != randomColor)
            }
        })
    });
    sectorWeightings = sectorWeightings.sort((x, y) => y.weighting.totalPercentage - x.weighting.totalPercentage)
    return {
        sectors: detailsResp.data.sectors,
        fundInformation: detailsResp.data.fundInformation,
        portfolioFundSectorWeightings: sectorWeightings,
        colorMap: colorMap
    };

}) satisfies PageServerLoad;