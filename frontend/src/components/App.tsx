import {
  createBrowserRouter,
  Outlet,
  RouterProvider,
} from "react-router-dom";

import { createTheme, MantineProvider } from '@mantine/core';

import MainPage from './MainPage';
import CardBrowser from './CardBrowser';
import DeckBrowser from './DeckBrowser';
import Header from './Header';
import Login from './Login';
import SignUp from './SignUp';
import NotFoundPage from './NotFoundPage';
import StudyPage from './StudyPage';
import NewCard from './NewCard';
import { User, UserContext } from '../context/UserContext';
import { useEffect, useState } from "react";
import CONFIG from "../config";


const router = createBrowserRouter([
  {
    element: <><Header /><Outlet /></>,
    errorElement: <><Header /><NotFoundPage /></>,
    children: [
      {
        path: "/",
        element: <MainPage />,
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
        path: "/deck/:id",
        element: <CardBrowser />,
      },
      {
        path: "/study/:deck_id",
        element: <StudyPage />,
      },
      {
        path: "/new-card/:id",
        element: <NewCard />,
      }
    ],
  },
]);

const theme = createTheme({
  spacing: {
    xs: '0.3rem',
  },
});

const defaultUser: User = {
  userName: "",
};

function App() {
  const [user, setUser] = useState<User>(defaultUser);

  const fetchUser = async () => {
    try {
      const response = await fetch(`${CONFIG.backendURL}/api/user/whoami`, {
        method: "GET",
        credentials: "include",
      });

      if (response.ok) {
        const data = await response.json();
        setUser({ ...defaultUser, userName: data.email });
      } else {
        setUser({ ...defaultUser, userName: "" });
      }
    } catch (error: any) {
      console.error(error);
    }
  };

  useEffect(() => {
    fetchUser();
  }, []);

  return (
    <UserContext.Provider value={[user, setUser]}>
      <MantineProvider defaultColorScheme="auto" theme={theme}>
        <RouterProvider router={router} />
      </MantineProvider>
    </UserContext.Provider>
  );
}

export default App;