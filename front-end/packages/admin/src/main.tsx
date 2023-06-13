import React from 'react';
import Amplify from "aws-amplify";
import { AmplifyProvider } from '@aws-amplify/ui-react';
import { BrowserRouter } from "react-router-dom";
import { ThemeProvider } from "@mui/material/styles";
import { GetTheme } from "./styles/theme";
import { createGraphQLClient } from 'shared/lib/apollo/apollo-clients';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { ApolloProvider } from "@apollo/client";
import config from "./config/config";
import globalConfig from "shared/config/config";
import '@aws-amplify/ui-react/styles.css';

Amplify.configure({
  Auth: {
    mandatorySignIn: true,
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

const defaultClient = createGraphQLClient(
  globalConfig.apiGateway.HTTP,
  globalConfig.apiGateway.WS
);

ReactDOM.render(
  <ApolloProvider client={defaultClient}>
    <React.StrictMode>
      <ThemeProvider theme={GetTheme()}>
        <AmplifyProvider>
          <BrowserRouter>
            <App />
          </BrowserRouter>
        </AmplifyProvider>
      </ThemeProvider>
    </React.StrictMode>
  </ApolloProvider>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
