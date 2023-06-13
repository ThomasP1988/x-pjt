export enum BidAsk {
    Bid,
    Ask
}

export enum SellBuy {
    Buy, // should stay 0
    Sell // should stay 1
};

export enum LimitMarket {
    Limit, // should stay 0
    Market // should stay 1
};

export enum OrderStatus {
    Empty = "empty",
    Filled = "filled",
    Cancelled = "cancelled",
    Open = "open",
    PartiallyFilled = "partially_filled"
}

