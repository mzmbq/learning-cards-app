
import {
  Group,
  Button,
  Box,
} from "@mantine/core";
import classes from "./Header.module.css";
import { Link, redirect } from "react-router-dom";

export function Header() {

  return (
    <Box pb={30}>
      <header className={classes.header}>
        <Group justify="space-between">

          <Group className={classes.logo}>
            <img src="/logo.svg" />
          </Group>

          <Group h="100%" gap={0}>
            <Link to="/study" className={classes.link}>
              Study
            </Link>

            <Link to="/decks" className={classes.link}>
              My Decks
            </Link>

            <Link to="/cards" className={classes.link}>
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
          </Group>

        </Group>
      </header>
    </Box>
  );
}


export default Header;