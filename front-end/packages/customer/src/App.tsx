import React, { useState, useEffect, useCallback, useReducer } from 'react';
import { Amplify } from 'aws-amplify';
import { Link } from "react-router-dom";
import { Auth } from '@aws-amplify/auth';
import '@aws-amplify/ui-react/styles.css';

import config from "./config/config";
import {
  AppBar, Toolbar, Typography, IconButton, ThemeProvider, Theme, Menu, MenuItem, Button, Box, Container, CssBaseline,
  Grid, Paper
} from '@mui/material';
import LogoutIcon from '@mui/icons-material/Logout';
import MenuIcon from '@mui/icons-material/Menu';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';

import { UserContext } from "./context/user";
import { DepositNFTContext } from "./context/deposit";
import { GetTheme } from "./styles/theme";
import { AppRouter } from './Router';
import { SignInDialog } from "shared/components/auth/SignInDialog";
import { SignUpDialog } from "shared/components/auth/SignUpDialog";
import { ConfirmationCodeDialog } from "shared/components/auth/ConfirmationCodeDialog";
import { NotificationButton } from 'shared/components/NotificationButton';
import { MenuSearch } from 'shared/components/search/MenuSearch';
import { User } from 'shared/repositories/user/__generated__/User';
import { useLazyQuery } from '@apollo/client';
import { ME } from 'shared/repositories/user/queries/me';
import { me } from 'shared/repositories/user/queries/__generated__/me';
import { initNFTSelectionState, NFTSelectionReducer } from "shared/reducers/nft-selection/reducer";
Amplify.configure({
  Auth: {
    mandatorySignIn: false,
    region: config.cognito.REGION,
    userPoolId: config.cognito.USER_POOL_ID,
    identityPoolId: config.cognito.IDENTITY_POOL_ID,
    userPoolWebClientId: config.cognito.APP_CLIENT_ID
  },
  Storage: {
    AWSS3: {
      region: config.s3.REGION,
      bucket: config.s3.BUCKET,
    }
  },
});

const AuthStateApp: React.FunctionComponent = () => {

  const [anchorElNav, setAnchorElNav] = useState<null | HTMLElement>(null);
  const [isOpenSignInDialog, setOpenSignInDialog] = useState<boolean>(false);
  const [isOpenSignUpDialog, setOpenSignUpDialog] = useState<boolean>(false);
  const [isOpenConfirmationCodeDialog, setOpenConfirmationCodeDialog] = useState<boolean>(false);
  const [emailConfirmationCode, setEmailConfirmationCode] = useState<string | undefined>();
  const [passwordConfirmationCode, setPasswordConfirmationCode] = useState<string | undefined>();
  const [stateDeposit, dispatchDeposit] = useReducer(NFTSelectionReducer, initNFTSelectionState());

  const theme: Theme = GetTheme("dark");
  let [user, setUser] = useState(null); // cognito user
  let [me, setMe] = useState<User | null>(null); // db user
  // const [getMe, { loading: loadingMe, error: errorME }] = useMutation<User>(ME, {
  //   fetchPolicy: "network-only"
  // });
  const [getMe, { loading: loadingMe, data: dataMe }] = useLazyQuery<me>(ME, {
    fetchPolicy: 'network-only',
  });

  const checkIfUserIsConnected = useCallback(() => {
    Auth.currentAuthenticatedUser().then((user) => {
      setUser(user)
      getMe()
    }).catch((e) => {
      console.log(e);
      setUser(null);
      setMe(null);
    });
  }, [setUser, setMe, getMe])

  useEffect(() => {
    checkIfUserIsConnected()
  }, [checkIfUserIsConnected]);

  useEffect(() => {
    if (user && !loadingMe && dataMe) {
      setMe(dataMe.me)
    }
  }, [user, me, loadingMe, dataMe]);

  const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElNav(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const openSignInDialog = () => {
    setOpenSignInDialog(true);
  }

  const openSignUpDialog = () => {
    setOpenSignUpDialog(true);
  }

  const onSignInSuccess = () => {
    console.log("onSignInSuccess");
    setOpenConfirmationCodeDialog(false);
    setOpenSignInDialog(false);
    console.log("isOpenSignInDialog", isOpenSignInDialog);
    checkIfUserIsConnected();
  }

  const onSignUpSuccess = (email: string, password: string) => {
    setOpenSignUpDialog(false);
    setOpenConfirmationCodeDialog(true);
    setEmailConfirmationCode(email);
    setPasswordConfirmationCode(password);
  }

  const closeConfirmationDialog = () => {
    setOpenConfirmationCodeDialog(false);
    setEmailConfirmationCode(undefined);
    setPasswordConfirmationCode(undefined);
    checkIfUserIsConnected();
  }

  const logOut = () => {
    Auth.signOut();
    setUser(null);
  }

  const pages = [
    {
      title: "Home",
      link: "/"
    },
    {
      title: "Deposit",
      link: "/deposit"
    },
    {
      title: "Withdrawal",
      link: "/withdrawal"
    },
    {
      title: "Discover",
      link: "/discover"
    },
  ]
  return <>
    <CssBaseline />
    <AppBar position="static" className="appbar" sx={{ backgroundImage: "url(/img/bg_75.gif)" }}>
      <ThemeProvider theme={theme}>
        <Container
          sx={{
            "@media (min-width: 600px)": {
              paddingLeft: 0,
              paddingRight: 0,
            }
          }}>
          <Toolbar variant="dense" color="inherit" sx={{
            "@media (min-width: 600px)": {
              paddingLeft: 0,
              paddingRight: 0,
            }
          }}>
            <Grid container>
              <Grid item xs={12} md={6} container alignContent="center">
                <Typography variant="h4">
                  nftquant.io
                </Typography>
              </Grid>
              {
                !user && <Grid item xs={12} md={6} container justifyContent="flex-end" alignContent="center">
                  <Button color="inherit" onClick={openSignInDialog} sx={{ marginRight: 1 }}>
                    Sign in
                  </Button>
                  <Button variant="outlined" color="inherit" onClick={openSignUpDialog}>
                    <Typography variant="subtitle2">
                      <strong>Get Started</strong>
                    </Typography>
                  </Button>
                </Grid>
              }
            </Grid>
          </Toolbar>
          <Toolbar variant="dense" sx={{
            "@media (min-width: 600px)": {
              paddingLeft: 0,
              paddingRight: 0,
            }
          }}>

            {/* MOBILE */}

            <Box sx={{ flexGrow: 1, display: { xs: 'flex', md: 'none' } }}>
              <IconButton
                size="large"
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleOpenNavMenu}
                color="inherit"
              >
                <MenuIcon />
              </IconButton>
              <Menu
                id="menu-appbar"
                anchorEl={anchorElNav}
                anchorOrigin={{
                  vertical: 'bottom',
                  horizontal: 'left',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'left',
                }}
                open={Boolean(anchorElNav)}
                onClose={handleCloseNavMenu}
                sx={{
                  display: { xs: 'block', md: 'none' },
                }}
              >
                {pages.map((page) => (
                  <MenuItem key={page.link} onClick={handleCloseNavMenu} component={Link} to={page.link}>
                    <Typography textAlign="center">{page.title}</Typography>
                  </MenuItem>
                ))}
              </Menu>
            </Box>
            {/* NOT MOBILE */}
            <Box sx={{ flexGrow: 1, display: { xs: 'none', md: 'flex' } }}>
              {pages.map((page) => (
                <Button
                  key={page.link}
                  component={Link} to={page.link}
                  sx={{ my: 2, color: 'white', display: 'block' }}
                >
                  {page.title}
                </Button>
              ))}
            </Box>
            <MenuSearch theme={GetTheme()} />
            {
              user && me && <>
                <IconButton aria-label="account">
                  <AccountCircleIcon />
                </IconButton>
                <NotificationButton user={me} theme={GetTheme()} />
                <IconButton aria-label="log-out" onClick={logOut}>
                  <LogoutIcon />
                </IconButton>
              </>
            }
          </Toolbar>
        </Container>
      </ThemeProvider>
    </AppBar>
    <UserContext.Provider value={{
      user,
      openSignInDialog
    }}>
      <DepositNFTContext.Provider value={{
        stateDeposit,
        dispatchDeposit
      }}>
        <Container sx={{ minHeight: "100vh", marginTop: -35, paddingTop: 5, position: "relative" }} component={Paper}>
          <AppRouter />
        </Container>
      </DepositNFTContext.Provider>
    </UserContext.Provider>
    <SignInDialog open={isOpenSignInDialog} close={() => setOpenSignInDialog(false)} onSuccess={onSignInSuccess} />
    <SignUpDialog open={isOpenSignUpDialog} close={() => setOpenSignUpDialog(false)} onSuccess={onSignUpSuccess} />
    <ConfirmationCodeDialog open={isOpenConfirmationCodeDialog}
      close={closeConfirmationDialog}
      email={emailConfirmationCode}
      password={passwordConfirmationCode}
      onSuccess={onSignInSuccess}
    />
  </>;
}

export default AuthStateApp;
