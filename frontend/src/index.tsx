import React from 'react';
import ReactDOM from 'react-dom/client';

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";

import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';

import App from './components/App';
import Login from './components/Login';
import SignUp from './components/SignUp';
import CardBrowser from './components/CardBrowser';
import DeckBrowser from './components/DeckBrowser';

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/signup",
    element: <SignUp />,
  },
  {
    path: "/decks",
    element: <DeckBrowser />,
  },
  {
    path: "/cards",
    element: <CardBrowser />,
  },
])


const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

root.render(
  <React.StrictMode>
    <MantineProvider>
      <RouterProvider router={router} />
    </MantineProvider>
  </React.StrictMode >
);