// Original file: proto/fund.proto

import type {
	FilterFundsResponseEntry as _proto_FilterFundsResponseEntry,
	FilterFundsResponseEntry__Output as _proto_FilterFundsResponseEntry__Output
} from '../proto/FilterFundsResponseEntry';

export interface SearchFundsResponse {
	entries?: _proto_FilterFundsResponseEntry[];
}

export interface SearchFundsResponse__Output {
	entries: _proto_FilterFundsResponseEntry__Output[];
}
