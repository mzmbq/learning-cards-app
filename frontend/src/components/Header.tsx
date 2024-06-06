
import {
  Group,
  Button,
  Box,
  useMantineColorScheme,
} from "@mantine/core";
import classes from "./Header.module.css";
import { Link } from "react-router-dom";

export function Header() {

  const { toggleColorScheme } = useMantineColorScheme();

  return (
    <Box pb={30}>
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
            <Link to="/login">
              <Button variant="default">Log in</Button>
            </Link>
            <Link to="/signup">
              <Button>Sign up</Button>
            </Link>
            <Button onClick={() => toggleColorScheme()}>Dark/Light</Button>
          </Group>

        </Group>
      </header>
    </Box>
  );
}


export default Header;