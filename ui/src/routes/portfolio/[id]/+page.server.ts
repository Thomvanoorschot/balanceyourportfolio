import {fundClient, portfolioClient} from "$lib/server/grpc";
import {safe} from "$lib/server/safe";
import {fail} from "@sveltejs/kit";
import type {Actions, PageServerLoad} from './$types';
import type {PortfolioDetailsRequest__Output} from "$lib/proto/proto/PortfolioDetailsRequest";
import type {PortfolioDetailsResponse__Output} from "$lib/proto/proto/PortfolioDetailsResponse";
import {colors, stringToRandomInteger} from "$lib/utils";
import type {PortfolioFundSectorWeighting__Output} from "$lib/proto/proto/PortfolioFundSectorWeighting";
import type {PortfolioSectorWeighting} from "$lib/portfolio";
import type {Holding} from "$lib/holding";
import type {FilterPortfolioFundHoldingsRequest__Output} from "$lib/proto/proto/FilterPortfolioFundHoldingsRequest.ts";
import type {
    FilterPortfolioFundHoldingsResponse__Output
} from "$lib/proto/proto/FilterPortfolioFundHoldingsResponse.ts";


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
    const colorMap = new Map<string, {fundName: string, color: string}>
    let sectorWeightings: PortfolioSectorWeighting[] = []
    Object.keys(detailsResp.data?.portfolioFundSectorWeightings).forEach((key) => {
        const value: PortfolioFundSectorWeighting__Output = detailsResp.data?.portfolioFundSectorWeightings[key];
        sectorWeightings.push({sectorName: key, weighting: value})
        value.fundSectorWeighting.forEach(x => {
            const color = colorMap.get(x.fundId)
            if (!color) {
                const randomColor = colorsCopy[stringToRandomInteger(x.fundId, colorsCopy.length)]
                colorMap.set(x.fundId, {fundName: x.fundName, color: randomColor})
                colorsCopy = colorsCopy.filter(x => x != randomColor)
            }
        })
    });
    sectorWeightings = sectorWeightings.sort((x, y) => y.weighting.totalPercentage - x.weighting.totalPercentage)
    const holdings: Holding[] = detailsResp.data.portfolioFundHoldings.map(x => (
        {
            id: x.holdingId,
            name: x.holdingName,
            ticker: x.ticker,
            percentage: x.cumulativePercentage,
            funds: x.funds.map(f => ({fundId: f.fundId, ratiodPercentage: f.ratiodPercentage}))
                .sort((a,b,) => a.fundId.localeCompare(b.fundId))
        }
    ))
    return {
        sectors: detailsResp.data.sectors,
        fundInformation: detailsResp.data.fundInformation,
        portfolioFundSectorWeightings: sectorWeightings,
        colorMap: colorMap,
        holdings: holdings
    }

}) satisfies PageServerLoad;


export const actions = {
    filterHoldings: async ({request, url, params,}) => {
        const formData = await request.formData()
        const holdingsLength = Number(formData.get("holdingsLength") || 0)
        const portfolioId = String(params.id || "")
        const searchTerm = String(formData.get("searchTerm") || "")
        const selectedSectors = JSON.parse(String(formData.get("selectedSectors")))
        const holdingsReq: FilterPortfolioFundHoldingsRequest__Output = {
            portfolioId: portfolioId,
            limit: 20,
            offset: holdingsLength,
            searchTerm: searchTerm,
            selectedSectors: selectedSectors
        }
        const holdingsResp = await safe(
            new Promise<FilterPortfolioFundHoldingsResponse__Output>((resolve, reject) => {
                portfolioClient.filterPortfolioHoldings(holdingsReq, (err, response) => {
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
                funds: x.funds.map(f => ({fundId: f.fundId, ratiodPercentage: f.ratiodPercentage}))
                    .sort((a,b,) => a.fundId.localeCompare(b.fundId))
            }
        ))
        return {
            holdings: holdings
        };
    }
} satisfies Actions;