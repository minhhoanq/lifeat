// Original file: src/proto/payment_service.proto


// Original file: src/proto/payment_service.proto

export const _payment_service_PaymentStatus_Values = {
  PENDING: 0,
  SUCCESS: 1,
  FAILED: 2,
} as const;

export type _payment_service_PaymentStatus_Values =
  | 'PENDING'
  | 0
  | 'SUCCESS'
  | 1
  | 'FAILED'
  | 2

export type _payment_service_PaymentStatus_Values__Output = typeof _payment_service_PaymentStatus_Values[keyof typeof _payment_service_PaymentStatus_Values]

export interface PaymentStatus {
}

export interface PaymentStatus__Output {
}
