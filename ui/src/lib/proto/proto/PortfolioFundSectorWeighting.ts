// Original file: proto/portfolio.proto

import type {
	PortfolioFundSectorWeightingEntry as _proto_PortfolioFundSectorWeightingEntry,
	PortfolioFundSectorWeightingEntry__Output as _proto_PortfolioFundSectorWeightingEntry__Output
} from '../proto/PortfolioFundSectorWeightingEntry';

export interface PortfolioFundSectorWeighting {
	totalPercentage?: number | string;
	fundSectorWeighting?: _proto_PortfolioFundSectorWeightingEntry[];
}

export interface PortfolioFundSectorWeighting__Output {
	totalPercentage: number;
	fundSectorWeighting: _proto_PortfolioFundSectorWeightingEntry__Output[];
}
