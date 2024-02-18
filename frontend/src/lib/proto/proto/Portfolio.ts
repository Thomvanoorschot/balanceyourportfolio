// Original file: proto/portfolio.proto

import type { PortfolioListItem as _proto_PortfolioListItem, PortfolioListItem__Output as _proto_PortfolioListItem__Output } from '../proto/PortfolioListItem';

export interface Portfolio {
  'id'?: (string);
  'name'?: (string);
  'entries'?: (_proto_PortfolioListItem)[];
}

export interface Portfolio__Output {
  'id': (string);
  'name': (string);
  'entries': (_proto_PortfolioListItem__Output)[];
}
