import { ReactElement, useState } from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import DialogContent from '@mui/material/DialogContent';
import DialogTitle from '@mui/material/DialogTitle';
import DialogActions from '@mui/material/DialogActions';
import Stack from '@mui/material/Stack';


export type ConfirmationProps = {
	message: string | ReactElement;
	title?: string;
	buttonText?: string | null;
	children: any;
}

const defaultCallbackEvent: Function = (): void => { };
let callbackFc: Function = defaultCallbackEvent;

export function ConfirmationDialog({ message = "", title = "", buttonText, children }: ConfirmationProps) {
	const [isOpen, setIsOpen] = useState(false);

	const show = (callback: any) => (event: any) => {
		event.preventDefault()

		event = {
			...event,
			target: { ...event.target, value: event.target.value }
		}

		callbackFc = (): void => callback(event);
		setIsOpen(true);
	}

	const onConfirm = () => {
		callbackFc()
		handleClose();
	}

	const handleClose = () => {
		setIsOpen(false);
		callbackFc = defaultCallbackEvent;
	};

	return (
		<>
			{children(show)}
			<Dialog
				open={isOpen}
				keepMounted
				fullWidth
				onClose={handleClose}
				aria-labelledby="alert-dialog-slide-title"
				aria-describedby="alert-dialog-slide-description"
			>
				<DialogTitle id="alert-dialog-title">{title || "Confirmation"}</DialogTitle>
				<DialogContent dividers>
					<Stack direction="row">
						{message}
					</Stack>
				</DialogContent>
				<DialogActions>
					<Button onClick={handleClose} autoFocus>
						cancel
					</Button>
					<Button onClick={onConfirm} color="secondary">
						{
							buttonText || "delete"
						}
					</Button>
				</DialogActions>
			</Dialog>
		</>
	)
}