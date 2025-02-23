// Original file: src/proto/payment_service.proto

import type { Long } from '@grpc/proto-loader';

export interface CreatePaymentRequest {
  'orderId'?: (string);
  'onlPaymentIntentId'?: (string);
  'amount'?: (number | string | Long);
  'paymentMethod'?: (string);
}

export interface CreatePaymentRequest__Output {
  'orderId'?: (string);
  'onlPaymentIntentId'?: (string);
  'amount'?: (number);
  'paymentMethod'?: (string);
}
