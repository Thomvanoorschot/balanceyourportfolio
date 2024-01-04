// Original file: proto/portfolio.proto

import type { Long } from '@grpc/proto-loader';

export interface FilterPortfolioFundHoldingsRequest {
	portfolioId?: string;
	searchTerm?: string;
	selectedSectors?: string[];
	limit?: number | string | Long;
	offset?: number | string | Long;
}

export interface FilterPortfolioFundHoldingsRequest__Output {
	portfolioId: string;
	searchTerm: string;
	selectedSectors: string[];
	limit: number;
	offset: number;
}
