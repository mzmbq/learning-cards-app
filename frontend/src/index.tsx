import React from 'react';
import ReactDOM from 'react-dom/client';

import {
  createBrowserRouter,
  Outlet,
  RouterProvider,
} from "react-router-dom";

import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';

import App from './components/App';
import CardBrowser from './components/CardBrowser';
import DeckBrowser from './components/DeckBrowser';
import Header from './components/Header';
import Login from './components/Login';
import SignUp from './components/SignUp';
import ErrorPage from './components/ErrorPage';
import StudyPage from './components/StudyPage';

const router = createBrowserRouter([
  {
    element: <><Header /><Outlet /></>,
    errorElement: <><Header /><ErrorPage /></>,
    children: [
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
      {
        path: "/study",
        element: <StudyPage />,
      }
    ],
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