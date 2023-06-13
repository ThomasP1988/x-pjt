import React, { useState } from 'react';
import { ConfirmationDialog } from '../ConfirmationDialog';
import { CancelOrder } from '../../repositories/orderbook-grpc';
import LoadingButton from '@mui/lab/LoadingButton';
import { CancelOrderArgs, Order } from '../../grpc/orderbook';
import { ErrorDialog } from "../ErrorDialog";

export type CancelProps = {
    orderId: string,
    updatedOrder?: (order: Order) => void
}

export const CancelOrderButton = ({ orderId, updatedOrder }: CancelProps) => {
    const [loadingCancel, setLoadingCancel] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    const cancel = async () => {
        setLoadingCancel(true);
        setError(null);

        const args: CancelOrderArgs = {
            orderId
        }

        let orderUpdated: Order | null = null;

        try {
            orderUpdated = await CancelOrder(args);
        } catch (e) {
            setLoadingCancel(false);
            if ((e as Error).message) {
                setError((e as Error).message);
            }
            return;
        }
        setLoadingCancel(false);
        if (orderUpdated && updatedOrder) {
            updatedOrder(orderUpdated)
        }
    }

    return <ConfirmationDialog
        title="Confirmation"
        message="Are you sure you want to delete this order?"
        buttonText="Delete"
    >
        {
            (confirm: any) => {
                return <><LoadingButton
                    onClick={confirm(() => cancel())}
                    loading={loadingCancel}
                >
                    Cancel Order
                </LoadingButton>
                    <ErrorDialog error={error || ""} handleClose={() => setError(null)} open={Boolean(error)} />
                </>
            }
        }
    </ConfirmationDialog>
} 