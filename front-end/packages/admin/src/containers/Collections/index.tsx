import { Box, Button, Dialog, DialogContent, DialogTitle, IconButton, Stack, Tab, Tabs, Typography } from "@mui/material";
import { useState } from "react";
import { ListCollection } from "shared/components/ListCollections";
import { Collection } from "shared/repositories/collection/__generated__/Collection";
import { CollectionStatus } from "shared/__generated__/globalTypes";
import CloseIcon from '@mui/icons-material/Close';
import { CollectionValidation } from "../../admin-components/CollectionValidation";

export const Collections = () => {

    const [selectedCollection, setSelectedCollection] = useState<Collection | null>(null);
    const [selectedCollectionStatus, setSelectedCollectionStatus] = useState<CollectionStatus | null>(null)

    return <>
        <Typography variant="h4" gutterBottom>Collections</Typography>
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs value={selectedCollectionStatus} onChange={(_, newValue: CollectionStatus | null) => setSelectedCollectionStatus(newValue)} aria-label="basic tabs example">
                <Tab label="All" value={null} />
                <Tab label="Accepted" value={CollectionStatus.ACCEPTED} />
                <Tab label="Pending validation" value={CollectionStatus.PENDING_VALIDATION} />
                <Tab label="denied" value={CollectionStatus.DENIED} />
            </Tabs>
        </Box>
        <ListCollection status={selectedCollectionStatus} limit={1} actions={(item: Collection) => {
            return <>
                {
                    item.status === CollectionStatus.PENDING_VALIDATION && <Button onClick={() => setSelectedCollection(item)}>Process validation</Button>
                }
            </>
        }} />
        <Dialog
            open={Boolean(selectedCollection)}
            onClose={() => setSelectedCollection(null)}
            aria-labelledby="alert-dialog-title"
            aria-describedby="alert-dialog-description"
            fullWidth
            // fullScreen
        >
            <DialogTitle id="alert-dialog-title" >
                Process collection validation
                <IconButton
                    aria-label="close"
                    onClick={() => setSelectedCollection(null)}
                    sx={{
                        position: 'absolute',
                        right: 8,
                        top: 8,
                        color: (theme) => theme.palette.grey[500],
                    }}
                >
                    <CloseIcon />
                </IconButton>
            </DialogTitle>
            <DialogContent>
                    <Stack sx={{ paddingTop: 1 }}>
                        {selectedCollection && <CollectionValidation collection={selectedCollection} onSubmit={() => setSelectedCollection(null)} />}
                    </Stack>
            </DialogContent>
        </Dialog>
    </>

}