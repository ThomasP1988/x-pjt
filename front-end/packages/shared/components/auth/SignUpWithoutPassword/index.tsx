import { useState, useEffect } from "react";
import Grid from '@mui/material/Grid';
import Stack from '@mui/material/Stack';
import LoadingButton from '@mui/lab/LoadingButton';
import LockOpenIcon from '@mui/icons-material/LockOpen';
import TextField from '@mui/material/TextField';
import { Alert } from "@mui/material";
import { useMutation } from "@apollo/client";
import { INVITE_USER } from "../../../repositories/user/mutations/invite";
import { inviteUser_inviteUser, inviteUserVariables } from "../../../repositories/user/mutations/__generated__/inviteUser";

export type SignUpProps = {
    onSuccess?: (email: string) => void
}

export function SignUpWithoutPassword({ onSuccess }: SignUpProps) {
    const [error, setError] = useState<boolean | string>(false);
    const [invite, {  loading: loadingInvite, error: errorInvite }] = useMutation<inviteUser_inviteUser, inviteUserVariables>(INVITE_USER, {
        fetchPolicy: "network-only"
    });

    useEffect(() => {
        if (errorInvite) {
            setError(errorInvite.message)
        }
    }, [setError, errorInvite])

    const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);

        const email: string | undefined = data.get("emailOnly")?.toString();
        if (!email) {
            setError("Please enter your e-mail address.");
            return;
        }

        try {
            await invite({
                variables: {
                    email
                }
            })
            setError(false);
            onSuccess?.(email);
        } catch (e) {
            console.log(e);
            console.log(typeof e);
            const { message, toString } = e as Error;
            if (message) {
                setError(message)
            } else {
                setError(toString())
            }
        }
    };

    return (<>
        <Stack component="form" direction="row" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
            <TextField size="small" type="email" variant="outlined" label="Email Address" id="emailOnly" name="emailOnly" sx={{ width: 300 }} />
            <LoadingButton
                loading={loadingInvite}
                loadingPosition="start"
                startIcon={<LockOpenIcon />}
                variant="contained"
                type="submit"
                sx={{ marginLeft: -1, borderTopLeftRadius: 0, borderBottomLeftRadius: 0 }}
                color="success"
                disableElevation={true}

            >
                Get Started
            </LoadingButton>
        </Stack>
        {
            error && <Grid item xs={12} container justifyContent="flex-end">
                <Alert severity="warning" sx={{ width: "100%" }}>
                    {error}
                </Alert>
            </Grid>
        }
    </>)
}
