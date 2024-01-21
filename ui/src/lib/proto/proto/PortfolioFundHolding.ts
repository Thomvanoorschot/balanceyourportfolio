// Original file: proto/portfolio.proto

import type { PortfolioFundHoldingEntry as _proto_PortfolioFundHoldingEntry, PortfolioFundHoldingEntry__Output as _proto_PortfolioFundHoldingEntry__Output } from '../proto/PortfolioFundHoldingEntry';

export interface PortfolioFundHolding {
  'ticker'?: (string);
  'holdingId'?: (string);
  'holdingName'?: (string);
  'cumulativePercentage'?: (number | string);
  'funds'?: (_proto_PortfolioFundHoldingEntry)[];
}

export interface PortfolioFundHolding__Output {
  'ticker': (string);
  'holdingId': (string);
  'holdingName': (string);
  'cumulativePercentage': (number);
  'funds': (_proto_PortfolioFundHoldingEntry__Output)[];
}
