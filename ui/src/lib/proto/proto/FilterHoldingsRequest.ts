// Original file: proto/fund.proto

import type { Long } from '@grpc/proto-loader';

export interface FilterHoldingsRequest {
  'fundId'?: (string);
  'searchTerm'?: (string);
  'sectorName'?: (string);
  'limit'?: (number | string | Long);
  'offset'?: (number | string | Long);
}

export interface FilterHoldingsRequest__Output {
  'fundId': (string);
  'searchTerm': (string);
  'sectorName': (string);
  'limit': (number);
  'offset': (number);
}
