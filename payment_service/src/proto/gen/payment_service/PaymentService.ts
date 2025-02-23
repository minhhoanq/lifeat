// Original file: src/proto/payment_service.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { CreatePaymentRequest as _payment_service_CreatePaymentRequest, CreatePaymentRequest__Output as _payment_service_CreatePaymentRequest__Output } from '../payment_service/CreatePaymentRequest';
import type { CreatePaymentResponse as _payment_service_CreatePaymentResponse, CreatePaymentResponse__Output as _payment_service_CreatePaymentResponse__Output } from '../payment_service/CreatePaymentResponse';

export interface PaymentServiceClient extends grpc.Client {
  CreatePayment(argument: _payment_service_CreatePaymentRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_payment_service_CreatePaymentResponse__Output>): grpc.ClientUnaryCall;
  CreatePayment(argument: _payment_service_CreatePaymentRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_payment_service_CreatePaymentResponse__Output>): grpc.ClientUnaryCall;
  CreatePayment(argument: _payment_service_CreatePaymentRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_payment_service_CreatePaymentResponse__Output>): grpc.ClientUnaryCall;
  CreatePayment(argument: _payment_service_CreatePaymentRequest, callback: grpc.requestCallback<_payment_service_CreatePaymentResponse__Output>): grpc.ClientUnaryCall;
  createPayment(argument: _payment_service_CreatePaymentRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_payment_service_CreatePaymentResponse__Output>): grpc.ClientUnaryCall;
  createPayment(argument: _payment_service_CreatePaymentRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_payment_service_CreatePaymentResponse__Output>): grpc.ClientUnaryCall;
  createPayment(argument: _payment_service_CreatePaymentRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_payment_service_CreatePaymentResponse__Output>): grpc.ClientUnaryCall;
  createPayment(argument: _payment_service_CreatePaymentRequest, callback: grpc.requestCallback<_payment_service_CreatePaymentResponse__Output>): grpc.ClientUnaryCall;
  
}

export interface PaymentServiceHandlers extends grpc.UntypedServiceImplementation {
  CreatePayment: grpc.handleUnaryCall<_payment_service_CreatePaymentRequest__Output, _payment_service_CreatePaymentResponse>;
  
}

export interface PaymentServiceDefinition extends grpc.ServiceDefinition {
  CreatePayment: MethodDefinition<_payment_service_CreatePaymentRequest, _payment_service_CreatePaymentResponse, _payment_service_CreatePaymentRequest__Output, _payment_service_CreatePaymentResponse__Output>
}
