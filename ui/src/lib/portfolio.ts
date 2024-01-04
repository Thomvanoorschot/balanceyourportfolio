import type { PortfolioFundSectorWeighting__Output } from '$lib/proto/proto/PortfolioFundSectorWeighting';

export interface Portfolio {
	id: string;
	name: string;
	items: PortfolioListItem[];
}

export interface PortfolioListItem {
	id: string;
	fundId: string;
	name: string;
	amount: number | undefined;
}

export interface PortfolioHoldingsFilter {
	portfolioId: string;
	searchTerm: string;
	sectorName: string;
	limit: number;
	offset: number;
}

export interface PortfolioSectorWeighting {
	sectorName: string;
	weighting: PortfolioFundSectorWeighting__Output;
}
