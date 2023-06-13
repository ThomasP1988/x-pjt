import { useState, useEffect } from "react";
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LoadingButton from '@mui/lab/LoadingButton';
import CheckIcon from '@mui/icons-material/Check';
import SendIcon from '@mui/icons-material/Send';

import TextField from '@mui/material/TextField';
import { Auth } from 'aws-amplify';
import { Alert } from "@mui/material";

export type ConfirmationCodeProps = {
    onSuccess?: () => void,
    email?: string,
    password?: string
}

export function ConfirmationCode({ onSuccess, email: emailProps, password }: ConfirmationCodeProps) {
    console.log("emailProps", emailProps);
    const [email, setEmail] = useState<string>(emailProps || "");
    const [error, setError] = useState<boolean | string>(false);
    const [info, setInfo] = useState<boolean | string>(false);
    const [loading, setLoading] = useState<boolean>(false);
    const [loadingResend, setLoadingResend] = useState<boolean>(false);

    useEffect(() => {
        setEmail(emailProps || "")
    }, [emailProps, setEmail])

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);
        console.log("email", email);
        console.log("email 2", data.get("email")?.toString());

        const emailForm: string | undefined = email !== "" ? email : data.get("email")?.toString();
        if (!emailForm) {
            setError("Please enter your e-mail address.");
            return;
        }

        const code: string | undefined = data.get("code")?.toString();
        if (!code) {
            setError("Please enter your confirmation code.");
            return;
        }
        setLoading(true);
        setInfo(false);
        try {
            const result = await Auth.confirmSignUp(emailForm, code);
            console.log("result", result);
            setError(false);

            if (password) {
                await Auth.signIn(emailForm, password);
            }
            setLoading(false);
            onSuccess?.();
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

    const resendCode = async () => {
        setLoadingResend(true);
        setError(false);
        setInfo(false);
        try {
            await Auth.resendSignUp(email);
            setError(false);
            setLoadingResend(false);
            setInfo("Code has been succesfully sent.")
        } catch (e) {
            console.log(e);
            const { message, toString } = e as Error;
            if (message) {
                setError(message)
            } else {
                setError(toString())
            }
            setLoadingResend(false);
        }
    }

    return (
        <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
            <Grid item xs={12} container justifyContent="flex-end" spacing={1}>
                <Grid item xs={12} container sx={{ paddingBottom: 1 }}>
                    <Alert severity="info" sx={{ width: "100%" }}>
                        We've sent you a verification code on your e-mail address.
                    </Alert>
                </Grid>

                {
                    !emailProps && <Grid item xs={12} container justifyContent="flex-end">
                        <TextField
                            size="small" type="email" fullWidth variant="outlined"
                            label="Email Address" id="email" name="email" style={{ marginBottom: 10 }}
                            value={email}
                            onChange={(e) => { setEmail(e.target.value) }}
                        />
                    </Grid>
                }
                <Grid item xs={12} container>
                    <TextField type="text" size="small" fullWidth variant="outlined" label="Verification code" id="code" name="code" style={{ marginBottom: 10 }} />
                </Grid>
                {
                    error && <Grid item xs={12} container justifyContent="flex-end">
                        <Alert severity="warning" sx={{ width: "100%" }}>
                            {error}
                        </Alert>
                    </Grid>
                }
                {
                    info && <Grid item xs={12} container justifyContent="flex-end">
                        <Alert severity="success" sx={{ width: "100%" }}>
                            {info}
                        </Alert>
                    </Grid>
                }
                <Grid item xs={12} container justifyContent="flex-end">
                    <LoadingButton
                        loading={loadingResend}
                        loadingPosition="start"
                        startIcon={<SendIcon />}
                        variant="outlined"
                        onClick={resendCode}
                        sx={{ marginRight: 1 }}
                    >
                        Resend Code
                    </LoadingButton>
                    <LoadingButton
                        loading={loading}
                        loadingPosition="start"
                        startIcon={<CheckIcon />}
                        variant="outlined"
                        type="submit"
                    >
                        Verify
                    </LoadingButton>
                </Grid>
            </Grid>
        </Box>
    )
}