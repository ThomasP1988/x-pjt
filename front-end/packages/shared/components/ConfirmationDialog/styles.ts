import { createStyles, makeStyles, Theme } from '@mui/material/styles';

export default makeStyles((theme: Theme) =>
	createStyles({
		content: {
			marginBottom: theme.spacing(3)
		},
		submitButton: {
			marginBottom: 1,
			[theme.breakpoints.up('sm')]: {
				marginBottom: 0
			}
		}
	})
);
