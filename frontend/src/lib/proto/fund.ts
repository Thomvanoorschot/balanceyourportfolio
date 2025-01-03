import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';


type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  proto: {
    CompareFundRequest: MessageTypeDefinition
    CompareFundResponse: MessageTypeDefinition
    FilterFundHoldingsRequest: MessageTypeDefinition
    FilterFundHoldingsResponse: MessageTypeDefinition
    FilterFundsRequest: MessageTypeDefinition
    FilterFundsResponse: MessageTypeDefinition
    FilterFundsResponseEntry: MessageTypeDefinition
    FundDetailsRequest: MessageTypeDefinition
    FundDetailsResponse: MessageTypeDefinition
    FundHolding: MessageTypeDefinition
    FundInformation: MessageTypeDefinition
    FundSectorWeighting: MessageTypeDefinition
    NonOverlappingHolding: MessageTypeDefinition
    OverlappingHolding: MessageTypeDefinition
    SearchFundsRequest: MessageTypeDefinition
    SearchFundsResponse: MessageTypeDefinition
  }
}

