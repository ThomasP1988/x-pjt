import React from 'react';
import Stack from '@mui/material/Stack';
import CircularProgress from '@mui/material/CircularProgress';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
// import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import { Order } from '../../../grpc/orderbook';
import LightbulbIcon from '@mui/icons-material/Lightbulb';
import Typography from '@mui/material/Typography';
import { useColors } from "../../../lib/hooks/colors";
type Props = {
    open: boolean,
    loading: boolean,
    error: string | null,
    close?: () => void,
    order: Order | null
}

export function ProcessingOrderDialog({ open, loading, close, error, order }: Props) {

    const colors = useColors();

    const handleClose = () => {
        close?.();
    };

    return (
        <>
            <Dialog
                open={open}
                onClose={handleClose}
                aria-labelledby="alert-dialog-title"
                aria-describedby="alert-dialog-description"
                fullWidth
            >
                <DialogTitle id="alert-dialog-title">
                    Processing Order
                </DialogTitle>
                <DialogContent>
                    <Stack justifyContent="center" alignContent="center" alignItems="center" sx={{ paddingTop: 1 }}>
                        {
                            loading ?
                                <CircularProgress />
                                :
                                <>
                                    {
                                        error ? <>
                                            <LightbulbIcon sx={{ color: colors.blue, fontSize: 42 }} />
                                            <Typography variant="h4">Something happened.</Typography>
                                            <Typography variant="body1">{error === "" ? "An error happened while trying to process your order, our teams have been warned, please try again later" : error}</Typography>
                                        </> : <>
                                            {
                                                order && <>
                                                    success
                                                </>
                                            }
                                        </>
                                    }
                                </>
                        }
                    </Stack>
                    {/* <DialogContentText id="alert-dialog-description">
                        Let Google help apps determine location. This means sending anonymous
                        location data to Google, even when no apps are running.
                    </DialogContentText> */}
                </DialogContent>
                <DialogActions>
                    <Button onClick={handleClose}>Close</Button>
                </DialogActions>
            </Dialog>
        </>
    );
}
