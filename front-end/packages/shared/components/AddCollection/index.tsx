import { FormEvent, useState } from "react";
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import LoadingButton from '@mui/lab/LoadingButton';
import SendIcon from '@mui/icons-material/Send';
import TextField from '@mui/material/TextField';
import Alert from '@mui/material/Alert';
import { useMutation } from "@apollo/client";
import { submitCollection, submitCollectionVariables, submitCollection_submitCollection } from "../../repositories/collection/mutations/__generated__/submitCollection";
import { SUBMIT_COLLECTION } from "../../repositories/collection/mutations/submit";

export type AddCollectionProps = {
    onSuccess?: (collection: submitCollection_submitCollection) => void
}

export const AddCollection = ({ onSuccess }: AddCollectionProps) => {
    const [error, setError] = useState<boolean | string>(false);
    const [collection, setCollection] = useState<submitCollection_submitCollection | null>(null);

    const [submit, { loading, error: errorSubmit }] = useMutation<submitCollection, submitCollectionVariables>(SUBMIT_COLLECTION, {
        fetchPolicy: "network-only"
    });

    async function handleSubmit(event: FormEvent<HTMLFormElement>) {
        event.preventDefault();
        const data = new FormData(event.currentTarget);
        setError(false);

        const address: string | undefined = data.get("address")?.toString();
        if (!address) {
            setError("Please enter the contract address of the collection you are trying to submit.");
            return;
        }
        const description: string | undefined = data.get("description")?.toString();

        try {
            const addedCollection: submitCollection_submitCollection | null | undefined = (await submit({
                variables: {
                    address,
                    description
                }
            })).data?.submitCollection;

            if (addedCollection) {
                setCollection(addedCollection);
                onSuccess?.(addedCollection);
            }
        } catch (e) {
            console.log(e);
        }
    }

    return (
        <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
            <Grid item xs={12} container justifyContent="flex-end" spacing={1}>
                <Grid item xs={12} container>
                    <TextField
                        size="small"
                        fullWidth
                        variant="outlined"
                        label="Contract address"
                        id="address"
                        name="address"
                        style={{ marginBottom: 10 }}
                        helperText="Each collection have an address to identify it. Example:0xaadc2d4261199ce24a4b0a57370c4fcf43bb60aa"
                    />
                </Grid>
                <Grid item xs={12} container>
                    <TextField
                        size="small"
                        multiline
                        fullWidth
                        variant="outlined"
                        label="Description"
                        id="description"
                        name="description"
                        style={{ marginBottom: 10 }}
                        rows={4}
                    />
                </Grid>
                {
                    (error || errorSubmit?.message) && <Grid item xs={12} container justifyContent="flex-end">
                        <Alert severity="warning" sx={{ width: "100%" }}>
                            {error || errorSubmit?.message}
                        </Alert>
                    </Grid>
                }
                {
                    collection && <Grid item xs={12} container justifyContent="flex-end">
                        <Alert severity="info" sx={{ width: "100%" }}>
                            The collection has been successfully submited, it is now under verification.
                        </Alert>
                    </Grid>
                }
                <Grid item xs={12} container justifyContent="flex-end">
                    <LoadingButton
                        loading={loading}
                        loadingPosition="start"
                        startIcon={<SendIcon />}
                        variant="outlined"
                        type="submit"
                    >
                        Submit new collection
                    </LoadingButton>
                </Grid>
            </Grid>
        </Box>
    )
}