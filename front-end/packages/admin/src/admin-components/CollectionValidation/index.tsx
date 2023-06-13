
import { useState } from "react";
import { Grid, Chip, Alert } from "@mui/material"
import { Collection } from "shared/repositories/collection/__generated__/Collection"
import LoadingButton from '@mui/lab/LoadingButton';
import CheckIcon from '@mui/icons-material/Check';
import DoDisturbIcon from '@mui/icons-material/DoDisturb';
import { useMutation } from "@apollo/client";
import { PropertyToField } from "shared/components/PropertyToField";
import { VALIDATE_COLLECTION } from "shared/repositories/collection/mutations/validate";
import { validateCollection, validateCollectionVariables } from "shared/repositories/collection/mutations/__generated__/validateCollection";
import { CollectionStatus } from "shared/__generated__/globalTypes";
import { ImageFromBucket } from "shared/components/ImageFromBucket";

export type CollectionValidationProps = {
    collection: Collection,
    onSubmit?: (updatedCollection: Collection) => void
}

export const CollectionValidation = ({ collection, onSubmit }: CollectionValidationProps) => {
    const [modifiedFields, setModifiedFields] = useState<Partial<validateCollectionVariables>>({});
    const [validate, { loading, error }] = useMutation<validateCollection, validateCollectionVariables>(VALIDATE_COLLECTION, {
        fetchPolicy: "network-only"
    });

    const submit = async (status: CollectionStatus) => {
    
        const variables: validateCollectionVariables = {
            address: collection.id,
            ...modifiedFields,
            status
        }

        try {
            const updatedCollection = await validate({
                variables
            });
            updatedCollection.data?.validateCollection && onSubmit?.(updatedCollection.data?.validateCollection);
        } catch (e) {
            console.log("submit", e);
        }
    };

    const onChange = ({ keyProperty, newValue }: {
        keyProperty: string;
        newValue: string;
    }): void => {
        console.log("key", keyProperty);
        const newModifiedFields = Object.assign({}, modifiedFields);
        (newModifiedFields as any)[keyProperty] = newValue;
        setModifiedFields(newModifiedFields);
    }
    console.log("modifiedFields", modifiedFields);

    return <Grid container justifyContent="flex-end" spacing={1}>
        {
            error && <Grid item xs={12} container justifyContent="flex-end">
                <Alert severity="warning" sx={{ width: "100%" }}>
                    {error.message}
                </Alert>
            </Grid>
        }
        <Grid item xs={12} container>
            <ImageFromBucket path={collection.thumbnailPath} maxHeight={180} maxWidth={180} />
        </Grid>
        <Grid item xs={12} container>
            <PropertyToField title="Name" keyProperty={"name"} value={modifiedFields["name"] || collection.name} onChange={onChange} />
        </Grid>
        <Grid item xs={12} container>
            <PropertyToField title="Symbol" keyProperty={"symbol"} value={modifiedFields["symbol"] || collection.symbol} onChange={onChange} />
        </Grid>
        <Grid item xs={12} container>
            <PropertyToField title="Description" keyProperty={"description"} value={modifiedFields["description"] || collection.description || ""} onChange={onChange} multiline rows={3} />
        </Grid>
        <Grid item xs={12} container>
            <PropertyToField title="Opensea slug" keyProperty={"openseaSlug"} value={modifiedFields["openseaSlug"] || collection.openseaSlug || ""} onChange={onChange} />
            <Chip
                label="Try opensea link"
                component="a"
                href={`https://opensea.io/assets/${collection.id}/${collection.firstItemId}`}
                variant="outlined"
                clickable
                target="_blank"
            />
        </Grid>
        <Grid item xs={12} container justifyContent="flex-end">
            <LoadingButton
                loading={loading}
                loadingPosition="start"
                startIcon={<DoDisturbIcon />}
                variant="outlined"
                onClick={() => submit(CollectionStatus.DENIED)}
                color="error"
            >
                Deny
            </LoadingButton>
            <LoadingButton
                sx={{ marginLeft: 1 }}
                loading={loading}
                loadingPosition="start"
                startIcon={<CheckIcon />}
                variant="outlined"
                onClick={() => submit(CollectionStatus.ACCEPTED)}
                color="success"
            >
                Verify
            </LoadingButton>
        </Grid>
    </Grid>
}