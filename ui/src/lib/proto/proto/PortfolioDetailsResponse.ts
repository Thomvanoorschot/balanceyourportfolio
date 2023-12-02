// Original file: proto/portfolio.proto

import type { FundInformation as _proto_FundInformation, FundInformation__Output as _proto_FundInformation__Output } from '../proto/FundInformation';
import type { PortfolioFundSectorWeightings as _proto_PortfolioFundSectorWeightings, PortfolioFundSectorWeightings__Output as _proto_PortfolioFundSectorWeightings__Output } from '../proto/PortfolioFundSectorWeightings';

export interface PortfolioDetailsResponse {
  'sectors'?: (string)[];
  'fundInformation'?: (_proto_FundInformation)[];
  'portfolioFundSectorWeightings'?: (_proto_PortfolioFundSectorWeightings)[];
}

export interface PortfolioDetailsResponse__Output {
  'sectors': (string)[];
  'fundInformation': (_proto_FundInformation__Output)[];
  'portfolioFundSectorWeightings': (_proto_PortfolioFundSectorWeightings__Output)[];
}
