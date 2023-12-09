export interface Fund {
    id: string;
    name: string;
    tickers: string[];
}

export interface FundSectorWeighting {
    sectorName: string;
    percentage: number;
}

export interface FundHoldingsFilter {
    fundId: string;
    searchTerm: string;
    sectorName: string;
    limit: number;
    offset: number;
}