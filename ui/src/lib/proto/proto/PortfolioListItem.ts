// Original file: proto/portfolio.proto

export interface PortfolioListItem {
	id?: string;
	fundId?: string;
	name?: string;
	amount?: number | string;
}

export interface PortfolioListItem__Output {
	id: string;
	fundId: string;
	name: string;
	amount: number;
}
