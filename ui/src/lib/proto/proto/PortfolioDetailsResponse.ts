// Original file: proto/portfolio.proto

import type {
	FundInformation as _proto_FundInformation,
	FundInformation__Output as _proto_FundInformation__Output
} from '../proto/FundInformation';
import type {
	PortfolioFundSectorWeighting as _proto_PortfolioFundSectorWeighting,
	PortfolioFundSectorWeighting__Output as _proto_PortfolioFundSectorWeighting__Output
} from '../proto/PortfolioFundSectorWeighting';
import type {
	PortfolioFundHolding as _proto_PortfolioFundHolding,
	PortfolioFundHolding__Output as _proto_PortfolioFundHolding__Output
} from '../proto/PortfolioFundHolding';

export interface PortfolioDetailsResponse {
	sectors?: string[];
	fundInformation?: _proto_FundInformation[];
	portfolioFundSectorWeightings?: { [key: string]: _proto_PortfolioFundSectorWeighting };
	portfolioFundHoldings?: _proto_PortfolioFundHolding[];
}

export interface PortfolioDetailsResponse__Output {
	sectors: string[];
	fundInformation: _proto_FundInformation__Output[];
	portfolioFundSectorWeightings: { [key: string]: _proto_PortfolioFundSectorWeighting__Output };
	portfolioFundHoldings: _proto_PortfolioFundHolding__Output[];
}
