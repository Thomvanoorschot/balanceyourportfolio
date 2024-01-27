// Original file: proto/fund.proto


export interface FundSectorWeighting {
  'sectorName'?: (string);
  'percentage'?: (number | string);
  'fundId'?: (string);
}

export interface FundSectorWeighting__Output {
  'sectorName': (string);
  'percentage': (number);
  'fundId': (string);
}
