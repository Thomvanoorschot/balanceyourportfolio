// Original file: proto/fund.proto

import type { Long } from '@grpc/proto-loader';

export interface FilterFundsRequest {
  'searchTerm'?: (string);
  'providers'?: (string)[];
  'limit'?: (number | string | Long);
  'offset'?: (number | string | Long);
}

export interface FilterFundsRequest__Output {
  'searchTerm': (string);
  'providers': (string)[];
  'limit': (number);
  'offset': (number);
}
