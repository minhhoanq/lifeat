import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { PaymentServiceClient as _payment_service_PaymentServiceClient, PaymentServiceDefinition as _payment_service_PaymentServiceDefinition } from './payment_service/PaymentService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  payment_service: {
    CreatePaymentRequest: MessageTypeDefinition
    CreatePaymentResponse: MessageTypeDefinition
    Payment: MessageTypeDefinition
    PaymentService: SubtypeConstructor<typeof grpc.Client, _payment_service_PaymentServiceClient> & { service: _payment_service_PaymentServiceDefinition }
    PaymentStatus: MessageTypeDefinition
  }
}

