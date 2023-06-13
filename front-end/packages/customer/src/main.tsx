import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from "react-router-dom";
import { Web3ReactProvider } from '@web3-react/core'
import Web3 from 'web3'
import { ApolloProvider } from '@apollo/client';
import { SnackbarProvider } from 'notistack';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { GetTheme, Colors } from "./styles/theme";
import { ThemeProvider } from '@mui/material/styles';
import "./index.css";
import { provider } from 'web3-core';
import { GetClient } from 'shared/lib/apollo/appsync';
import { ColorsProvider } from 'shared/lib/hooks/colors';
import config from "shared/config/config";

function getLibrary(provider: provider) {
  return new Web3(provider)
}

const defaultClient = GetClient(
  config.appsync.HTTP,
  config.appsync.REGION
);

ReactDOM.render(
  <React.StrictMode>
    <ApolloProvider client={defaultClient}>
      <ThemeProvider theme={GetTheme()}>
        <SnackbarProvider maxSnack={3}>
          <ColorsProvider colors={Colors}>
            <BrowserRouter>
              <Web3ReactProvider getLibrary={getLibrary}>
                <App />
              </Web3ReactProvider>  
            </BrowserRouter>
          </ColorsProvider>
        </SnackbarProvider>
      </ThemeProvider>
    </ApolloProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
