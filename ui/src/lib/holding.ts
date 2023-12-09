export interface Holding {
    id: string;
    ticker: string;
    name: string;
    percentage: number;
    funds: FundHolding[]
}

export interface FundHolding {
    fundId: string;
    ratiodPercentage: number;
}