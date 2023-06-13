import { AppBar, Container, Divider, Drawer, Grid, IconButton, List, ListItem, ListItemIcon, ListItemText, Paper, Toolbar, Typography } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import CollectionsIcon from '@mui/icons-material/Collections';
import LogoutIcon from '@mui/icons-material/Logout';
import { styled, Theme, useTheme, ThemeProvider } from '@mui/material/styles';
import { Authenticator } from '@aws-amplify/ui-react';
import { useState } from 'react';
import './App.css';
import { GetTheme } from './styles/theme';
import { AppRouter } from './Router';
import { Link } from "react-router-dom";

const DrawerHeader = styled('div')(({ theme }: { theme: Theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: 'flex-end',
}));

function App() {
  const theme = useTheme();

  const [menuOpen, setMenuOpen] = useState(false);

  const handleDrawerOpen = () => {
    setMenuOpen(true);
  };

  const handleDrawerClose = () => {
    setMenuOpen(false);
  };


  return (
    <Authenticator hideSignUp={true}>
      {({ signOut, user }) => (
        <>
          <AppBar position="sticky" className="appbar">
            <ThemeProvider theme={GetTheme("dark")}>
              <Toolbar>
                <Grid container>
                  <Grid item sm={6} container direction="row" alignItems="center">
                    <IconButton
                      color="inherit"
                      aria-label="open drawer"
                      onClick={handleDrawerOpen}
                      edge="start"
                      sx={{ mr: 2, ...(menuOpen && { display: 'none' }) }}
                    >
                      <MenuIcon />
                    </IconButton>
                    <Typography variant="h6" noWrap component="div">
                      Admin
                    </Typography>
                  </Grid>
                  <Grid item sm={6} container justifyContent="flex-end" alignItems="center">
                    {user.attributes?.email}
                    <IconButton aria-label="log-out" onClick={signOut} sx={{ marginLeft: 1 }}>
                      <LogoutIcon />
                    </IconButton>
                  </Grid>
                </Grid>
              </Toolbar>
            </ThemeProvider>
          </AppBar>
          <Drawer
            anchor="left"
            open={menuOpen}
            variant="persistent"
            sx={{
              width: 300,
              flexShrink: 0,
              '& .MuiDrawer-paper': {
                width: 300,
                boxSizing: 'border-box',
              },
            }}
          >
            <DrawerHeader>
              <IconButton onClick={handleDrawerClose}>
                {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
              </IconButton>
            </DrawerHeader>
            <Divider />
            <List>
              <ListItem button component={Link} to={"/collections"} >
                <ListItemIcon>
                  <CollectionsIcon />
                </ListItemIcon>
                <ListItemText primary="Collections" />
              </ListItem>
            </List>
          </Drawer>
          <Container sx={{ padding: 3 }}>
            <AppRouter />
          </Container>
        </>
      )}
    </Authenticator>
  );
}

export default App;
