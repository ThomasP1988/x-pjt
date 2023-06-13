import { OrderStatus } from '../../constants';
import { Colors } from "../../lib/hooks/colors";

export const PrintStatus = (input: OrderStatus): string => {
    const labelledStatus: Record<OrderStatus, string> = {
        [OrderStatus.Open]: "Open",
        [OrderStatus.Cancelled]: "Cancelled",
        [OrderStatus.Empty]: "Empty",
        [OrderStatus.PartiallyFilled]: "Partially Filled",
        [OrderStatus.Filled]: "Filled",
    }
    return labelledStatus[input];
}

export const ColorStatus = (input: OrderStatus, colors: Colors): string => {
    const labelledStatus: Record<OrderStatus, string> = {
        [OrderStatus.Open]: colors.green,
        [OrderStatus.Cancelled]: colors.red,
        [OrderStatus.Empty]: colors.orange,
        [OrderStatus.PartiallyFilled]: colors.green,
        [OrderStatus.Filled]: colors.blue,
    }
    return labelledStatus[input];
}