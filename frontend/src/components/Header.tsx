
import {
  Group,
  Button,
  Box,
  useMantineColorScheme,
  LoadingOverlay,
} from "@mantine/core";
import classes from "./Header.module.css";
import { Link } from "react-router-dom";
import { useUserContext } from "../context/UserContext";
import { useState } from "react";
import CONFIG from "../config";

export function Header() {
  const [loading, setLoading] = useState(false);
  const { toggleColorScheme } = useMantineColorScheme();

  const [user, setUser] = useUserContext();


  const signOut = async () => {
    try {
      setLoading(true);
      const response = await fetch(`${CONFIG.backendURL}/api/user/signout`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        let errorText = await response.text();
        throw new Error("Signout failed: " + errorText);
      }

      setUser({ ...user, userName: "" });


    } catch (error: any) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const buttonsLoginSignup = (<>
    <Link to="/login">
      <Button variant="default">Log in</Button>
    </Link>
    <Link to="/signup">
      <Button>Sign up</Button>
    </Link>
  </>);

  const userButtonSignout = (<>
    <p>Signed in as <b>{user.userName}</b></p>
    <Link to="/signup">
      <Button onClick={() => signOut()}>Sign out</Button>
    </Link>
  </>);

  return (
    <Box pb={30}>
      <LoadingOverlay visible={loading} />

      <header className={classes.header}>
        <Group justify="space-between">

          <Group className={classes.logo}>
            <img alt="logo" src="/logo.svg" />
          </Group>

          <Group h="100%" gap={0}>
            <Link to="/study" className={classes.link}>
              Study
            </Link>

            <Link to="/decks" className={classes.link}>
              My Decks
            </Link>

            <Link to="/cards" className={classes.link}>
              My Cards
            </Link>

            <Link to="/new-card" className={classes.link}>
              New Card
            </Link>
          </Group>

          <Group>
            {
              user.userName == "" ? buttonsLoginSignup : userButtonSignout
            }
            <Button onClick={() => toggleColorScheme()}>Dark/Light</Button>
          </Group>

        </Group>
      </header>
    </Box>
  );
}


export default Header;