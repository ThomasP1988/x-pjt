import Stack from '@mui/material/Stack';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import { AddCollection, AddCollectionProps } from '../AddCollection';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';

type Props = {
    open: boolean,
    close?: () => void,
} & AddCollectionProps;

export function AddCollectionDialog({ open, close, ...addCollectionProps }: Props) {

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
                    Submit new collection
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
                        <AddCollection {...addCollectionProps} />
                    </Stack>
                </DialogContent>
            </Dialog>
        </>
    );
}
