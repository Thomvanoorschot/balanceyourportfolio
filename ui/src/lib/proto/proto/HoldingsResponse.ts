// Original file: proto/fund.proto


export interface HoldingsResponse {
  'ticker'?: (string);
  'name'?: (string);
  'type'?: (string);
  'sector'?: (string);
  'amount'?: (number | string);
  'percentageOfTotal'?: (number | string);
  'marketValue'?: (number | string);
}

export interface HoldingsResponse__Output {
  'ticker': (string);
  'name': (string);
  'type': (string);
  'sector': (string);
  'amount': (number);
  'percentageOfTotal': (number);
  'marketValue': (number);
}
