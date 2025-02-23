export class Payment {
    constructor(
        public id: string,
        public order_id: string,
        public onl_payment_intent_id: string,
        public amount: number,
        public status: string,
        public payment_method: string,
    ) { }
}