// Original file: proto/portfolio.proto

import type { Long } from '@grpc/proto-loader';

export interface FilterPortfolioHoldingsRequest {
  'portfolioId'?: (string);
  'searchTerm'?: (string);
  'sectorName'?: (string);
  'limit'?: (number | string | Long);
  'offset'?: (number | string | Long);
}

export interface FilterPortfolioHoldingsRequest__Output {
  'portfolioId': (string);
  'searchTerm': (string);
  'sectorName': (string);
  'limit': (number);
  'offset': (number);
}
