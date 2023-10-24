export interface Fund {
    id: string;
    name: string;
    tickers: string[];
}
export interface FundDetails {
    information: FundInformation;
    sectors: string[];
    sectorWeightings: FundSectorWeighting[];
}

export interface FundSectorWeighting {
    sectorName: string;
    percentage: number;
}

export interface FundHolding {
    ticker: string;
    name: string;
    type: string;
    sector: string;
    amount: number;
    percentageOfTotal: number;
    marketValue: number;
}

export interface FundInformation {
    id: string;
    name: string;
    outstandingShares: string;
    effectiveDate: string;
}

export interface FundHoldingsFilter {
    fundId: string;
    searchTerm: string;
    sectorName: string;
    limit: number;
    offset: number;
}