// Original file: proto/fund.proto

import type { Long } from '@grpc/proto-loader';

export interface FilterFundHoldingsRequest {
	fundId?: string;
	searchTerm?: string;
	selectedSectors?: string[];
	limit?: number | string | Long;
	offset?: number | string | Long;
}

export interface FilterFundHoldingsRequest__Output {
	fundId: string;
	searchTerm: string;
	selectedSectors: string[];
	limit: number;
	offset: number;
}
