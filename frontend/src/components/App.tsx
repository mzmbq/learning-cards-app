import { createBrowserRouter, Outlet, RouterProvider } from "react-router-dom";

import { createTheme, MantineProvider } from "@mantine/core";

import MainPage from "../pages/Home/MainPage";
import CardBrowser from "../pages/Cards/CardBrowser";
import DeckBrowser from "../pages/Decks/DeckBrowser";
import Header from "./Header";
import Login from "../pages/Login/Login";
import SignUp from "../pages/SignUp/SignUp";
import NotFoundPage from "../pages/NotFound/NotFoundPage";
import StudyPage from "../pages/Study/StudyPage";
import CardCreator from "../pages/CardCreate/CardCreator";
import { User, UserContext } from "../context/UserContext";
import { useEffect, useState } from "react";
import CONFIG from "../config";
import SearchPage from "../pages/Search/SearchPage";

const router = createBrowserRouter([
  {
    element: (
      <>
        <Header />
        <Outlet />
      </>
    ),
    errorElement: (
      <>
        <Header />
        <NotFoundPage />
      </>
    ),
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
        path: "/search",
        element: <SearchPage />,
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
        element: <CardCreator />,
      },
    ],
  },
]);

const theme = createTheme({
  spacing: {
    xs: "0.3rem",
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
