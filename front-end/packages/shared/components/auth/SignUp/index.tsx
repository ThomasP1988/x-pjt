import { useState } from "react";
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LoadingButton from '@mui/lab/LoadingButton';
import LockOpenIcon from '@mui/icons-material/LockOpen';
import TextField from '@mui/material/TextField';
import { Auth } from 'aws-amplify';
import { Alert } from "@mui/material";

export type SignUpProps = {
    onSuccess?: (email: string, password: string) => void
}

export function SignUp({ onSuccess }: SignUpProps) {
    const [error, setError] = useState<boolean | string>(false);
    const [loading, setLoading] = useState<boolean>(false);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);

        const email: string | undefined = data.get("email")?.toString();
        if (!email) {
            setError("Please enter your e-mail address.");
            return;
        }

        const password: string | undefined = data.get("password")?.toString();
        if (!password) {
            setError("Please confirm your password.");
            return;
        }
        const confirmPassword: string | undefined = data.get("confirmPassword")?.toString();
        if (!confirmPassword) {
            setError("Please enter your password.");
            return;
        }
        if (confirmPassword !== password) {
            setError("The passwords you entered are different.");
            return;
        }

        setLoading(true);
        try {
            await Auth.signUp({
                password,
                username: email
            });
            setError(false);
            setLoading(false);
            onSuccess?.(email, password);
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
                    <TextField size="small" type="email" fullWidth variant="outlined" label="Email Address" id="email" name="email" style={{ marginBottom: 10 }} />
                </Grid>
                <Grid item xs={12} container spacing={1}>
                    <Grid item xs={6} container>
                        <TextField type="password" size="small" fullWidth variant="outlined" label="Password" id="password" name="password" style={{ marginBottom: 10 }} />
                    </Grid>
                    <Grid item xs={6} container>
                        <TextField type="password" size="small" fullWidth variant="outlined" label="Confirm Password" id="confirmPassword" name="confirmPassword" style={{ marginBottom: 10 }} />
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
                        Sign Up
                    </LoadingButton>
                </Grid>
            </Grid>
        </Box>
    )
}
