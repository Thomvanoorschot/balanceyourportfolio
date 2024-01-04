// Original file: proto/fund.proto

import type {
	FundInformation as _proto_FundInformation,
	FundInformation__Output as _proto_FundInformation__Output
} from '../proto/FundInformation';
import type {
	FundSectorWeighting as _proto_FundSectorWeighting,
	FundSectorWeighting__Output as _proto_FundSectorWeighting__Output
} from '../proto/FundSectorWeighting';
import type {
	FundHolding as _proto_FundHolding,
	FundHolding__Output as _proto_FundHolding__Output
} from '../proto/FundHolding';

export interface FundDetailsResponse {
	information?: _proto_FundInformation | null;
	sectorWeightings?: _proto_FundSectorWeighting[];
	sectors?: string[];
	fundHoldings?: _proto_FundHolding[];
}

export interface FundDetailsResponse__Output {
	information: _proto_FundInformation__Output | null;
	sectorWeightings: _proto_FundSectorWeighting__Output[];
	sectors: string[];
	fundHoldings: _proto_FundHolding__Output[];
}
