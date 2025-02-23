// Original file: src/proto/payment_service.proto

import type { _payment_service_PaymentStatus_Values, _payment_service_PaymentStatus_Values__Output } from '../payment_service/PaymentStatus';
import type { Long } from '@grpc/proto-loader';

export interface Payment {
  'id'?: (string);
  'orderId'?: (string);
  'onlPaymentIntentId'?: (string);
  'amount'?: (number | string | Long);
  'status'?: (_payment_service_PaymentStatus_Values);
  'paymentMethod'?: (string);
}

export interface Payment__Output {
  'id'?: (string);
  'orderId'?: (string);
  'onlPaymentIntentId'?: (string);
  'amount'?: (number);
  'status'?: (_payment_service_PaymentStatus_Values__Output);
  'paymentMethod'?: (string);
}
