// Original file: proto/fund.proto


export interface FundHolding {
  'ticker'?: (string);
  'holdingId'?: (string);
  'holdingName'?: (string);
  'cumulativePercentage'?: (number | string);
}

export interface FundHolding__Output {
  'ticker': (string);
  'holdingId': (string);
  'holdingName': (string);
  'cumulativePercentage': (number);
}
