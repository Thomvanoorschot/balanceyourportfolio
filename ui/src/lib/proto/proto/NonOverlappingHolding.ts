// Original file: proto/fund.proto


export interface NonOverlappingHolding {
  'holdingId'?: (string);
  'holdingName'?: (string);
  'holdingTicker'?: (string);
  'nonOverlappingPercentage'?: (number | string);
}

export interface NonOverlappingHolding__Output {
  'holdingId': (string);
  'holdingName': (string);
  'holdingTicker': (string);
  'nonOverlappingPercentage': (number);
}
