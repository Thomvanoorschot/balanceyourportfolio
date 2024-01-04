// Original file: proto/fund.proto

export interface FundInformation {
	id?: string;
	name?: string;
	outstandingShares?: number | string;
	effectiveDate?: string;
}

export interface FundInformation__Output {
	id: string;
	name: string;
	outstandingShares: number;
	effectiveDate: string;
}
