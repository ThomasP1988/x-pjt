import { useState, MouseEvent, useEffect } from 'react';
import {
    IconButton, Popover, Theme, ThemeProvider, List, ListItemText, ListItemButton, Typography, CircularProgress,
    Badge, Grid, ListItem, Stack
} from '@mui/material';
import CircleNotificationsIcon from '@mui/icons-material/CircleNotifications';
import { OnSubscriptionDataOptions, useQuery, useSubscription, useMutation, ApolloQueryResult } from '@apollo/client';
import { LIST_NOTIFICATIONS } from '../../repositories/notification/queries/list';
import { listNotifications, listNotificationsVariables, listNotifications_listNotifications_notifications } from '../../repositories/notification/queries/__generated__/listNotifications';
import { SnackbarKey, useSnackbar } from 'notistack';
import { onAddedNotification } from '../../repositories/notification/subscriptions/new';
import { notification, notification_notification } from '../../repositories/notification/subscriptions/__generated__/notification';
import { User } from '../../repositories/user/__generated__/User';
import CloseIcon from '@mui/icons-material/Close';
import { SET_LAST_SEEN_NOTIFICATION } from '../../repositories/user/mutations/setLastSeen';
import { setLastSeenNotification } from '../../repositories/user/mutations/__generated__/setLastSeenNotification';
import { LoadingButton } from '@mui/lab';
import MoreTimeIcon from '@mui/icons-material/MoreTime';

export type NotificationButtonArgs = {
    user: User,
    theme: Theme,
}

export const NotificationButton = ({ user, theme }: NotificationButtonArgs) => {
    const [lastSeen, setLastSeen] = useState<Date | undefined>(user?.lastSeenNotification ? new Date(user?.lastSeenNotification) : undefined);
    const [anchorEl, setAnchorEl] = useState<HTMLButtonElement | null>(null);
    const [next, setNext] = useState<string | null>(null);
    const [notifications, setNotifications] = useState<notification_notification[]>([]);
    const [setLastSeenRequest] = useMutation<setLastSeenNotification>(SET_LAST_SEEN_NOTIFICATION);
    const { loading, error, fetchMore } = useQuery<listNotifications, listNotificationsVariables>(LIST_NOTIFICATIONS, {
        variables: {
            limit: 10
        },
        onCompleted: (data: listNotifications): void => {
            if (data.listNotifications) {
                setNotifications(data.listNotifications.notifications as notification_notification[]);
                setNext(data.listNotifications.next)
            }
        }
    });

    const { data: createSubData, error: createSubError } =  useSubscription<notification>(onAddedNotification, {
        shouldResubscribe: true,
        variables: {
            userId: user.id
        },
        onSubscriptionData: ({ subscriptionData: { data } }: OnSubscriptionDataOptions<notification>) => {
            if (data?.notification) {
                const notif: notification_notification = data.notification;

                let key: SnackbarKey | undefined;
                key = enqueueSnackbar(<Grid container>
                    <Grid item container xs={10}>
                        <Grid item xs={12}>
                            <Typography>
                                {notif.type}
                            </Typography>
                        </Grid>
                        <Grid item xs={12}>
                            <Typography variant="subtitle2">
                                {notif.message}
                            </Typography>
                        </Grid>
                    </Grid>
                    <Grid item xs={2} container justifyContent="flex-end">
                        <IconButton
                            onClick={() => closeSnackbar(key)}
                            aria-label="close"
                        >
                            <CloseIcon />
                        </IconButton>
                    </Grid>
                </Grid>)
                const mergedNotifications: notification_notification[] = [notif, ...notifications]
                setNotifications(mergedNotifications);
            }

        }
    });

    console.log("createSubError", createSubError);
    console.log("createSubData", createSubData);

    const { enqueueSnackbar, closeSnackbar } = useSnackbar();

    useEffect(() => {
        if (user) {
            setLastSeen(new Date(user.lastSeenNotification))
        }
    }, [user])


    const handleClose = () => {
        setAnchorEl(null);
    };

    const open = Boolean(anchorEl);
    const id = open ? 'notifications' : undefined;

    const showBadge = (): boolean => {
        return Boolean(lastSeen) && (+(lastSeen as Date) < +new Date(notifications[0]?.createdAt))
    }

    const handleClick = async (event: MouseEvent<HTMLButtonElement>) => {
        setAnchorEl(event.currentTarget);
        if (showBadge()) {
            try {
                await setLastSeenRequest();
            } catch (e) {
                console.log(e);
            }
            setLastSeen(new Date());
        }
    };

    const loadMore = async () => {
        try {
            const result: ApolloQueryResult<listNotifications> = await fetchMore({
                variables: {
                    from: next,
                    limit: 10
                }
            });

            if (result.data.listNotifications) {
                const mergedNotifications: notification_notification[] = [...notifications, ...(result.data.listNotifications?.notifications as notification_notification[] || [])]
                setNotifications(mergedNotifications);
                setNext(result.data.listNotifications.next);
            }
        } catch (e) {
            console.log(e);
        }
    }

    return (
        <>
            <IconButton aria-describedby={id} onClick={handleClick}>
                <Badge color="secondary" variant="dot" invisible={!showBadge()}>
                    <CircleNotificationsIcon />
                </Badge>
            </IconButton>
            <ThemeProvider theme={theme}>
                <Popover
                    id={id}
                    open={open}
                    anchorEl={anchorEl}
                    onClose={handleClose}
                    anchorOrigin={{
                        vertical: 'bottom',
                        horizontal: 'left',
                    }}
                >
                    {
                        error && <Typography color="error">{error.message}</Typography>
                    }
                    {
                        loading && <CircularProgress />
                    }
                    <List sx={{
                        width: 300,
                        height: 400,
                        bgcolor: 'background.paper'
                    }} component="nav">
                        {
                            notifications?.map((item: listNotifications_listNotifications_notifications | null) => {
                                return <ListItemButton key={item?.id}>
                                    <ListItemText primary={item?.type} secondary={item?.message} />
                                </ListItemButton>
                            })
                        }
                        {
                            next && <ListItem>
                                <Stack justifyContent="center">
                                    <LoadingButton startIcon={<MoreTimeIcon />} onClick={loadMore} loading={loading}>Load More</LoadingButton>
                                </Stack>
                            </ListItem>
                        }
                    </List>
                </Popover>
            </ThemeProvider>
        </>
    )
}