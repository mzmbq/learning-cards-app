import { createBrowserRouter, Outlet, RouterProvider } from "react-router-dom";

import MainPage from "../pages/Home/MainPage";
import CardBrowser from "../pages/Cards/CardBrowser";
import DeckBrowser from "../pages/Decks/DeckBrowser";
import Login from "../pages/Login/Login";
import SignUp from "../pages/SignUp/SignUp";
import NotFoundPage from "../pages/NotFound/NotFoundPage";
import StudyPage from "../pages/Study/StudyPage";
import CardCreator from "../pages/CardCreate/CardCreator";
import SearchPage from "../pages/Search/SearchPage";
import Header from "../components/Header";

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

function Pages() {
  return <RouterProvider router={router} />;
}

export default Pages;
