import React from 'react';
import Stack from '@mui/material/Stack';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { ConfirmationCode, ConfirmationCodeProps } from '../ConfirmationCode';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';

type Props = {
    open: boolean,
    close?: () => void,
} & ConfirmationCodeProps;

export function ConfirmationCodeDialog({ open, close, ...confirmationProps }: Props) {

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
                    Verify
                    <IconButton
                        aria-label="close"
                        onClick={close}
                        sx={{
                            position: 'absolute',
                            right: 8,
                            top: 8,
                            color: (theme) => theme.palette.grey[500],
                        }}
                    >
                        <CloseIcon />
                    </IconButton>
                </DialogTitle>
                <DialogContent>
                    <Stack justifyContent="center" alignContent="center" alignItems="center" sx={{ paddingTop: 1 }}>
                        <ConfirmationCode {...confirmationProps} />
                    </Stack>
                </DialogContent>
            </Dialog>
        </>
    );
}
