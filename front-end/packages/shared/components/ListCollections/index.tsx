import { useEffect, ReactElement } from 'react';
import { useLazyQuery } from '@apollo/client';
import { LIST_COLLECTIONS } from 'shared/repositories/collection/queries/list';
import { listCollections, listCollectionsVariables, listCollections_listCollections } from 'shared/repositories/collection/queries/__generated__/listCollections';
import { Collection } from "shared/repositories/collection/__generated__/Collection";
import { Alert, CircularProgress, Stack, Table, TableBody, TableCell, TableHead, TableRow, Tooltip, Typography } from '@mui/material';
import { LoadingButton } from '@mui/lab';
import { CollectionStatus } from 'shared/__generated__/globalTypes';

export type Props = {
    status: CollectionStatus | null,
    limit: number,
    actions?: (collectionId: Collection) => ReactElement
}

export const ListCollection = ({ status, limit, actions }: Props) => {

    const [queryCollections, { data, error, loading, fetchMore, refetch }] = useLazyQuery<listCollections, listCollectionsVariables>(LIST_COLLECTIONS);

    useEffect(() => {
        queryCollections({
            variables: {
                status,
                limit
            }
        })
    }, [queryCollections, status, refetch]);


    const loadMore = () => {
        fetchMore({
            variables: {
                status,
                from: data?.listCollections?.next,
                limit
            },
            updateQuery: (previousCollection: listCollections, { fetchMoreResult }): listCollections => {
                return {
                    listCollections: {
                        ...previousCollection.listCollections,
                        ...(fetchMoreResult?.listCollections || {}),
                        collections: [...(previousCollection.listCollections?.collections || []), ...(fetchMoreResult?.listCollections?.collections || [])],
                    } as listCollections_listCollections
                }
            }
        })
    }

    return <>
        {
            loading && <Stack direction="row" justifyContent="center"><CircularProgress /></Stack>
        }
        {
            error && <Alert severity="error">{error.message}</Alert>
        }
        {
            data?.listCollections && <>
                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Symbol</TableCell>
                            <TableCell>Name</TableCell>
                            <TableCell>Chain</TableCell>
                            <TableCell>Description</TableCell>
                            <TableCell>Status</TableCell>
                            <TableCell>Supply</TableCell>
                            {
                                actions && <TableCell></TableCell>
                            }
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {
                            data?.listCollections?.collections?.map((item: Collection, index: number) => {
                                return <TableRow
                                    key={index}
                                    sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                                >
                                    <TableCell component="th" scope="row">
                                        <Tooltip title={item.id} placement="right">
                                            <Typography variant="subtitle2">{item.symbol}</Typography>
                                        </Tooltip>
                                    </TableCell>
                                    <TableCell>
                                        {item.name}
                                    </TableCell>
                                    <TableCell>
                                        {item.chain}
                                    </TableCell>
                                    <TableCell>
                                        {item.description}
                                    </TableCell>
                                    <TableCell>
                                        {item.status}
                                    </TableCell>
                                    <TableCell>
                                        {item.supply}
                                    </TableCell>
                                    {
                                        actions && <TableCell>
                                            {actions(item)}
                                        </TableCell>
                                    }
                                </TableRow>
                            })
                        }
                    </TableBody>
                </Table>
                {
                    data?.listCollections?.next && <Stack direction="row" justifyContent="center">
                        <LoadingButton loading={loading} onClick={loadMore}>Load more</LoadingButton>
                    </Stack>
                }
            </>
        }
    </>
}