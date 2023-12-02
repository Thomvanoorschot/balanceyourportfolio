import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';


type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  proto: {
    FilterHoldingsRequest: MessageTypeDefinition
    FundDetailsResponse: MessageTypeDefinition
    FundInformation: MessageTypeDefinition
    FundSectorWeighting: MessageTypeDefinition
    GetFundDetailsRequest: MessageTypeDefinition
    HoldingsListResponse: MessageTypeDefinition
    HoldingsResponse: MessageTypeDefinition
    SearchFundsEntry: MessageTypeDefinition
    SearchFundsRequest: MessageTypeDefinition
    SearchFundsResponse: MessageTypeDefinition
  }
}

