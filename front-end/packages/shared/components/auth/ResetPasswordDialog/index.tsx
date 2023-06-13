import React from 'react';
import Stack from '@mui/material/Stack';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { ResetPassword, ResetPasswordProps } from '../ResetPassword';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';

type Props = {
    open: boolean,
    close?: () => void,
} & ResetPasswordProps;

export function ResetPasswordDialog({ open, close, ...resetPasswordProps }: Props) {

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
                    Reset Password
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
                        <ResetPassword {...resetPasswordProps} />
                    </Stack>
                </DialogContent>
            </Dialog>
        </>
    );
}
