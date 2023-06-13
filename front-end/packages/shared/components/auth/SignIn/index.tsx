import { useState, useEffect } from "react";
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LoadingButton from '@mui/lab/LoadingButton';
import LoginIcon from '@mui/icons-material/Login';
import TextField from '@mui/material/TextField';
import { Auth } from 'aws-amplify';
import { Alert, Link } from "@mui/material";
import { ResetPasswordDialog } from "../ResetPasswordDialog";
import { NewPasswordRequiredDialog } from "../NewPasswordRequiredDialog";

export type SigninProps = {
    onSuccess?: () => void
}

export function SignIn({ onSuccess }: SigninProps) {
    const [error, setError] = useState<boolean | string>(false);
    const [loading, setLoading] = useState<boolean>(false);
    const [user, setUser] = useState<any>();
    const [openResetPassword, setOpenResetPassword] = useState<boolean>(false);
    const [openSetNewPassword, setOpenSetNewPassword] = useState<boolean>(false);

    useEffect(() => {
        console.log("useEffect", user);
        if (user) {
            if (user.challengeName === "NEW_PASSWORD_REQUIRED") {
                setOpenSetNewPassword(true);
            } else {
                onSuccess?.();
            }
        }

    }, [user, onSuccess]);

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
            setError("Please enter your password.");
            return;
        }
        setLoading(true);
        try {
            const user = await Auth.signIn(email, password);
            console.log("user", user);
            setUser(user);
            setError(false);
            setLoading(false);
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

    const resetPasswordSuccess = async (email: string, password: string) => {
        setOpenResetPassword(false);
        setError(false);
        setLoading(true);
        try {
            const user = await Auth.signIn(email, password);
            setUser(user);
            setLoading(false);
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
    }

    const setNewPasswordSuccess = async (user: any) => {
        setOpenSetNewPassword(false);
        setError(false);
        console.log("user", user);
        if (user) {
            const newUser = Object.assign({}, user);
            setUser(newUser);
        }
    }

    return (
        <>
            <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
                <Grid item xs={12} container justifyContent="flex-end" spacing={1}>
                    <Grid item xs={12} container justifyContent="flex-end">
                        <TextField size="small" type="email" fullWidth variant="outlined" label="Email Address" id="email" name="email" style={{ marginBottom: 10 }} />
                    </Grid>
                    <Grid item xs={12} container justifyContent="flex-end">
                        <TextField type="password" size="small" fullWidth variant="outlined" label="Password" id="password" name="password" style={{ marginBottom: 10 }} />
                    </Grid>
                    {
                        error && <Grid item xs={12} container justifyContent="flex-end">
                            <Alert severity="warning" sx={{ width: "100%" }}>
                                {error}
                            </Alert>
                        </Grid>
                    }
                    <Grid item xs={12} container justifyContent="flex-end" alignContent="center" alignItems="center">
                        <Link onClick={() => setOpenResetPassword(true)} sx={{ marginRight: 1, cursor: "pointer" }}>Forgot password?</Link>
                        <LoadingButton
                            loading={loading}
                            loadingPosition="start"
                            startIcon={<LoginIcon />}
                            variant="outlined"
                            type="submit"
                        >
                            Sign in
                        </LoadingButton>
                    </Grid>
                </Grid>
            </Box>
            <ResetPasswordDialog open={openResetPassword} close={() => setOpenResetPassword(false)} onSuccess={resetPasswordSuccess} />
            <NewPasswordRequiredDialog open={openSetNewPassword} close={() => setOpenSetNewPassword(false)} onSuccess={setNewPasswordSuccess} user={user} />
        </>
    )
}