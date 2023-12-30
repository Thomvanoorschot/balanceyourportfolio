// Original file: proto/fund.proto


export interface FilterFundsResponseEntry {
  'id'?: (string);
  'name'?: (string);
  'tickers'?: (string)[];
  'marketCap'?: (number | string);
  'currency'?: (string);
}

export interface FilterFundsResponseEntry__Output {
  'id': (string);
  'name': (string);
  'tickers': (string)[];
  'marketCap': (number);
  'currency': (string);
}
