// Original file: proto/fund.proto


export interface OverlappingHolding {
  'holdingId'?: (string);
  'holdingName'?: (string);
  'overlappingPercentage'?: (number | string);
  'fundOnePercentage'?: (number | string);
  'fundTwoPercentage'?: (number | string);
}

export interface OverlappingHolding__Output {
  'holdingId': (string);
  'holdingName': (string);
  'overlappingPercentage': (number);
  'fundOnePercentage': (number);
  'fundTwoPercentage': (number);
}
