import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type {
	FundServiceClient as _proto_FundServiceClient,
	FundServiceDefinition as _proto_FundServiceDefinition
} from './proto/FundService';
import type {
	PortfolioServiceClient as _proto_PortfolioServiceClient,
	PortfolioServiceDefinition as _proto_PortfolioServiceDefinition
} from './proto/PortfolioService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
	new (...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
	proto: {
		Empty: MessageTypeDefinition;
		FilterFundHoldingsRequest: MessageTypeDefinition;
		FilterFundHoldingsResponse: MessageTypeDefinition;
		FilterFundsRequest: MessageTypeDefinition;
		FilterFundsResponse: MessageTypeDefinition;
		FilterFundsResponseEntry: MessageTypeDefinition;
		FilterPortfolioFundHoldingsRequest: MessageTypeDefinition;
		FilterPortfolioFundHoldingsResponse: MessageTypeDefinition;
		FundDetailsRequest: MessageTypeDefinition;
		FundDetailsResponse: MessageTypeDefinition;
		FundHolding: MessageTypeDefinition;
		FundInformation: MessageTypeDefinition;
		FundSectorWeighting: MessageTypeDefinition;
		FundService: SubtypeConstructor<typeof grpc.Client, _proto_FundServiceClient> & {
			service: _proto_FundServiceDefinition;
		};
		Portfolio: MessageTypeDefinition;
		PortfolioDetailsRequest: MessageTypeDefinition;
		PortfolioDetailsResponse: MessageTypeDefinition;
		PortfolioFundHolding: MessageTypeDefinition;
		PortfolioFundHoldingEntry: MessageTypeDefinition;
		PortfolioFundSectorWeighting: MessageTypeDefinition;
		PortfolioFundSectorWeightingEntry: MessageTypeDefinition;
		PortfolioListItem: MessageTypeDefinition;
		PortfolioService: SubtypeConstructor<typeof grpc.Client, _proto_PortfolioServiceClient> & {
			service: _proto_PortfolioServiceDefinition;
		};
		PortfoliosRequest: MessageTypeDefinition;
		PortfoliosResponse: MessageTypeDefinition;
		SearchFundsRequest: MessageTypeDefinition;
		SearchFundsResponse: MessageTypeDefinition;
		UpdatePortfolioFundAmountRequest: MessageTypeDefinition;
		UpsertPortfolioRequest: MessageTypeDefinition;
		UpsertPortfolioResponse: MessageTypeDefinition;
	};
}
