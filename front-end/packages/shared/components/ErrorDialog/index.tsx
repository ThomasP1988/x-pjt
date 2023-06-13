import * as React from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogActions from '@mui/material/DialogActions';
import DialogContent from '@mui/material/DialogContent';
import DialogContentText from '@mui/material/DialogContentText';
import DialogTitle from '@mui/material/DialogTitle';
import Stack from '@mui/material/Stack';
import FmdBadIcon from '@mui/icons-material/FmdBad';
import { useColors } from "../../lib/hooks/colors";

export type ErrorDialogArgs = {
    open?: boolean,
    handleClose: () => void,
    error: string
}

export function ErrorDialog({ open, handleClose, error }: ErrorDialogArgs) {
    const colors = useColors();

    return (
        <Dialog
            open={Boolean(open)}
            onClose={handleClose}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
            fullWidth
        >
            <DialogTitle id="alert-dialog-title">
                <Stack alignContent="center" alignItems="center" direction="row">
                    <FmdBadIcon sx={{ color: colors.orange, marginRight: 1 }} />
                    Oops! An error happened.
                </Stack>
            </DialogTitle>
            <DialogContent>
                <DialogContentText id="alert-dialog-description">
                    {error}
                </DialogContentText>
            </DialogContent>
            <DialogActions>
                <Button onClick={handleClose}>Close</Button>
            </DialogActions>
        </Dialog>
    );
}