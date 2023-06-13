import { useState } from "react";
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LoadingButton from '@mui/lab/LoadingButton';
import SendIcon from '@mui/icons-material/Send';
import RestartAltIcon from '@mui/icons-material/RestartAlt';
import TextField from '@mui/material/TextField';
import { Auth } from 'aws-amplify';
import { Alert } from "@mui/material";

export type ResetPasswordProps = {
    onSuccess?: (email: string, password: string) => void
}

export function ResetPassword({ onSuccess }: ResetPasswordProps) {
    const [email, setEmail] = useState<string>("");

    const [error, setError] = useState<boolean | string>(false);
    const [loading, setLoading] = useState<boolean>(false);
    const [codeLoading, setCodeLoading] = useState<boolean>(false);
    const [showNewPasswordStep, setShowNewPasswordStep] = useState<boolean>(false);

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);

        const email: string | undefined = data.get("forgotEmail")?.toString();
        if (!email) {
            setError("Please enter your e-mail address.");
            return;
        }

        const code: string | undefined = data.get("code")?.toString();
        if (!code) {
            setError("Please enter the code we sent you by mail.");
            return;
        }

        const password: string | undefined = data.get("password")?.toString();
        if (!password) {
            setError("Please enter your password.");
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
            const result = await Auth.forgotPasswordSubmit(email, code, password);
            console.log(result);
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

    const ResendCode = async () => {
        if (!email) {
            setError("Please enter your e-mail address.");
            return;
        }

        setCodeLoading(true);
        try {
            const result = await Auth.forgotPassword(email);
            console.log(result);
            setError(false);
            setCodeLoading(false);
            setShowNewPasswordStep(true);
        } catch (e) {
            console.log(e);
            console.log(typeof e);
            const { message, toString } = e as Error;
            if (message) {
                setError(message)
            } else {
                setError(toString())
            }
            setCodeLoading(false);
        }
    };

    return (
        <>
            <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
                <Grid item xs={12} container justifyContent="flex-end" spacing={1}>
                    <Grid item xs={12} container justifyContent="flex-end">
                        <TextField
                            size="small" type="email" fullWidth variant="outlined" label="Email Address"
                            id="forgotEmail" name="forgotEmail" style={{ marginBottom: 10 }}
                            value={email} onChange={(e) => { setEmail(e.target.value) }}
                        />
                    </Grid>
                    {
                        showNewPasswordStep && <>
                            <Grid item xs={12} container justifyContent="flex-end">
                                <TextField
                                    size="small" fullWidth variant="outlined" label="Verification Code"
                                    id="code" name="code" style={{ marginBottom: 10 }}
                                />
                            </Grid>
                            <Grid item xs={12} container spacing={1}>
                                <Grid item xs={6} container>
                                    <TextField type="password" size="small" fullWidth variant="outlined" label="Password" id="password" name="password" style={{ marginBottom: 10 }} />
                                </Grid>
                                <Grid item xs={6} container>
                                    <TextField type="password" size="small" fullWidth variant="outlined" label="Confirm Password" id="confirmPassword" name="confirmPassword" style={{ marginBottom: 10 }} />
                                </Grid>
                            </Grid>
                        </>
                    }
                    {
                        error && <Grid item xs={12} container justifyContent="flex-end">
                            <Alert severity="warning" sx={{ width: "100%" }}>
                                {error}
                            </Alert>
                        </Grid>
                    }
                    <Grid item xs={12} container justifyContent="flex-end">
                        <LoadingButton
                            loading={codeLoading}
                            loadingPosition="start"
                            startIcon={<SendIcon />}
                            variant="outlined"
                            onClick={ResendCode}
                            sx={{ marginRight: 1 }}
                        >
                            Send Recovery Code
                        </LoadingButton>
                        {
                            showNewPasswordStep && <LoadingButton
                                loading={loading}
                                loadingPosition="start"
                                startIcon={<RestartAltIcon />}
                                variant="outlined"
                                type="submit"
                            >
                                Reset Password
                            </LoadingButton>
                        }

                    </Grid>
                </Grid>
            </Box>
        </>
    )
}