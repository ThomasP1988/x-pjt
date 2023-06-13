import { useState } from "react";
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LoadingButton from '@mui/lab/LoadingButton';
import LockOpenIcon from '@mui/icons-material/LockOpen';
import TextField from '@mui/material/TextField';
import { Auth } from 'aws-amplify';
import { Alert } from "@mui/material";

export type NewPasswordRequiredProps = {
    onSuccess?: (user: string) => void
    user: any
}

export function NewPasswordRequired({ onSuccess, user }: NewPasswordRequiredProps) {
    const [error, setError] = useState<boolean | string>(false);
    const [loading, setLoading] = useState<boolean>(false);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);

        const oldPassword: string | undefined = data.get("oldPassword")?.toString();
        if (!oldPassword) {
            setError("Please enter your old password.");
            return;
        }

        const password: string | undefined = data.get("password")?.toString();
        if (!password) {
            setError("Please enter your new password.");
            return;
        }
        const confirmPassword: string | undefined = data.get("confirmPassword")?.toString();
        if (!confirmPassword) {
            setError("Please confirm your new password.");
            return;
        }
        if (confirmPassword !== password) {
            setError("The passwords you entered are different.");
            return;
        }

        setLoading(true);
        try {
            const userResult = await Auth.completeNewPassword(
                user,
                password
            );
            console.log("userResult", userResult);
            setError(false);
            setLoading(false);
            onSuccess?.(userResult);
        } catch (e) {
            console.log(e);
            console.log(typeof e);
            const { message, toString } = e as Error;
            if (message) {
                setError(message)
            } else {
                setError(toString())
            }
            setLoading(false);
        }
    };

    return (
        <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
            <Grid item xs={12} container justifyContent="flex-end" spacing={1}>
                <Grid item xs={12} container>
                    <TextField size="small" fullWidth variant="outlined" label="Old Password" id="oldPassword" name="oldPassword" style={{ marginBottom: 10 }} />
                </Grid>
                <Grid item xs={12} container spacing={1}>
                    <Grid item xs={6} container>
                        <TextField type="password" size="small" fullWidth variant="outlined" label="New Password" id="password" name="password" style={{ marginBottom: 10 }} />
                    </Grid>
                    <Grid item xs={6} container>
                        <TextField type="password" size="small" fullWidth variant="outlined" label="Confirm New Password" id="confirmPassword" name="confirmPassword" style={{ marginBottom: 10 }} />
                    </Grid>
                </Grid>
                {
                    error && <Grid item xs={12} container justifyContent="flex-end">
                        <Alert severity="warning" sx={{ width: "100%" }}>
                            {error}
                        </Alert>
                    </Grid>
                }
                <Grid item xs={12} container justifyContent="flex-end">
                    <LoadingButton
                        loading={loading}
                        loadingPosition="start"
                        startIcon={<LockOpenIcon />}
                        variant="outlined"
                        type="submit"
                    >
                        Set New Password
                    </LoadingButton>
                </Grid>
            </Grid>
        </Box>
    )
}
