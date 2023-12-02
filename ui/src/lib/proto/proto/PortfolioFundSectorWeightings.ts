// Original file: proto/portfolio.proto

import type { FundSectorWeighting as _proto_FundSectorWeighting, FundSectorWeighting__Output as _proto_FundSectorWeighting__Output } from '../proto/FundSectorWeighting';

export interface PortfolioFundSectorWeightings {
  'fundName'?: (string);
  'percentageOfTotal'?: (number | string);
  'fundSectorWeightings'?: (_proto_FundSectorWeighting)[];
}

export interface PortfolioFundSectorWeightings__Output {
  'fundName': (string);
  'percentageOfTotal': (number);
  'fundSectorWeightings': (_proto_FundSectorWeighting__Output)[];
}
