import { useState, useEffect, FC } from 'react';
import { Popover, Typography, Stack, Alert, List, ListItemButton, Skeleton, ListItem, ListItemText } from '@mui/material';
import { Collection } from "./collection";
import { searchCollections_searchCollections_results } from 'shared/repositories/collection/queries/__generated__/searchCollections';


type Data = {
    results: searchCollections_searchCollections_results[],
    total: number
}

export type Props = {
    search: string,
    anchor?: HTMLInputElement | null,
    error?: Error,
    loading?: boolean,
    data?: Data,
    onClick?: (item: searchCollections_searchCollections_results) => void
}

export const Suggestions: FC<Props> = ({ search, anchor, error, loading, data, onClick }: Props) => {

    const [resultPopoverOpen, setResultPopoverOpen] = useState<boolean>(false);

    useEffect(() => {
        if (search && search !== "") {
            setResultPopoverOpen(true)
        }
    }, [search, setResultPopoverOpen])


    function handleClose() {
        setResultPopoverOpen(false);
    }

    const id = resultPopoverOpen ? 'search-results' : undefined;

    const onSelect = (item: searchCollections_searchCollections_results) => {
        onClick?.(item)
        setResultPopoverOpen(false);
    }

    return (<>
        <Popover
            id={id}
            open={resultPopoverOpen}
            anchorEl={anchor}
            onClose={handleClose}
            anchorOrigin={{
                vertical: 'bottom',
                horizontal: 'right',
            }}
            transformOrigin={{
                vertical: 'top',
                horizontal: 'right',
            }}
            disableAutoFocus={true}
            disableEnforceFocus={true}
            sx={{ marginTop: 1 }}
        >
            <Stack sx={{ width: 600, padding: 2, height: 300 }}>
                <Typography>
                    Search for <Typography display="inline" sx={{ fontWeight: "bold" }}>{search}</Typography>
                </Typography>
                {
                    error && <Alert severity="warning">
                        <Typography> {error?.message ? error?.message : "Something went wrong."}</Typography>
                    </Alert>
                }
                <List component="nav">
                    {
                        (loading) && <>
                            <ListItem>
                                <ListItemText primary={<Skeleton />} secondary={<Skeleton />} />
                            </ListItem>
                            <ListItem>
                                <ListItemText primary={<Skeleton />} secondary={<Skeleton />} />
                            </ListItem>
                            <ListItem>
                                <ListItemText primary={<Skeleton />} secondary={<Skeleton />} />
                            </ListItem>
                        </>
                    }
                    {
                        data?.results?.map((item: searchCollections_searchCollections_results, index: number) => {
                            return <ListItemButton key={index} onClick={() => onSelect(item)}>
                                {
                                    {
                                        "Collection": <Collection key={index} item={item} />
                                    }[item.__typename]
                                }
                            </ListItemButton>
                        })
                    }
                </List>
            </Stack>
        </Popover>
    </>)
}