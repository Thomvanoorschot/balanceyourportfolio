// Original file: proto/portfolio.proto

import type { Long } from '@grpc/proto-loader';

export interface UpdatePortfolioFundAmountRequest {
  'portfolioId'?: (string);
  'fundId'?: (string);
  'amount'?: (number | string | Long);
}

export interface UpdatePortfolioFundAmountRequest__Output {
  'portfolioId': (string);
  'fundId': (string);
  'amount': (number);
}
