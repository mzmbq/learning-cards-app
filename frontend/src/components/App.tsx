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
import { UserContextProvider } from "../context/UserContext";
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

function App() {
  return (
    <UserContextProvider>
      <MantineProvider defaultColorScheme="auto" theme={theme}>
        <RouterProvider router={router} />
      </MantineProvider>
    </UserContextProvider>
  );
}

export default App;
