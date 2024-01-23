// Original file: proto/fund.proto

import type { OverlappingHolding as _proto_OverlappingHolding, OverlappingHolding__Output as _proto_OverlappingHolding__Output } from '../proto/OverlappingHolding';
import type { Long } from '@grpc/proto-loader';

export interface CompareFundResponse {
  'totalOverlappingPercentage'?: (number | string);
  'overlappingHoldings'?: (_proto_OverlappingHolding)[];
  'overlappingHoldingsCount'?: (number | string | Long);
  'fundOneHoldingCount'?: (number | string | Long);
  'fundOneOverlappingCountPercentage'?: (number | string);
  'fundTwoHoldingCount'?: (number | string | Long);
  'fundTwoOverlappingCountPercentage'?: (number | string);
  'fundOneName'?: (string);
  'fundTwoName'?: (string);
}

export interface CompareFundResponse__Output {
  'totalOverlappingPercentage': (number);
  'overlappingHoldings': (_proto_OverlappingHolding__Output)[];
  'overlappingHoldingsCount': (number);
  'fundOneHoldingCount': (number);
  'fundOneOverlappingCountPercentage': (number);
  'fundTwoHoldingCount': (number);
  'fundTwoOverlappingCountPercentage': (number);
  'fundOneName': (string);
  'fundTwoName': (string);
}
