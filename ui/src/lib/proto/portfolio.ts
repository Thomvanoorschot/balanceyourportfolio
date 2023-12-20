import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';


type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  proto: {
    FilterFundHoldingsRequest: MessageTypeDefinition
    FilterFundHoldingsResponse: MessageTypeDefinition
    FilterPortfolioFundHoldingsRequest: MessageTypeDefinition
    FilterPortfolioFundHoldingsResponse: MessageTypeDefinition
    FundDetailsRequest: MessageTypeDefinition
    FundDetailsResponse: MessageTypeDefinition
    FundHolding: MessageTypeDefinition
    FundInformation: MessageTypeDefinition
    FundSectorWeighting: MessageTypeDefinition
    Portfolio: MessageTypeDefinition
    PortfolioDetailsRequest: MessageTypeDefinition
    PortfolioDetailsResponse: MessageTypeDefinition
    PortfolioFundHolding: MessageTypeDefinition
    PortfolioFundHoldingEntry: MessageTypeDefinition
    PortfolioFundSectorWeighting: MessageTypeDefinition
    PortfolioFundSectorWeightingEntry: MessageTypeDefinition
    PortfolioListItem: MessageTypeDefinition
    PortfoliosRequest: MessageTypeDefinition
    PortfoliosResponse: MessageTypeDefinition
    SearchFundsEntry: MessageTypeDefinition
    SearchFundsRequest: MessageTypeDefinition
    SearchFundsResponse: MessageTypeDefinition
    UpsertPortfolioRequest: MessageTypeDefinition
    UpsertPortfolioResponse: MessageTypeDefinition
  }
}

