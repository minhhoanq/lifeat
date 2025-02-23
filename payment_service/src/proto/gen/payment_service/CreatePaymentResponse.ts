// Original file: src/proto/payment_service.proto

import type { Payment as _payment_service_Payment, Payment__Output as _payment_service_Payment__Output } from '../payment_service/Payment';

export interface CreatePaymentResponse {
  'payment'?: (_payment_service_Payment | null);
}

export interface CreatePaymentResponse__Output {
  'payment'?: (_payment_service_Payment__Output);
}
