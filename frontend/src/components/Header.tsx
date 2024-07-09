import {
  Group,
  Button,
  Box,
} from "@mantine/core";
import classes from "./Header.module.css";
import { Link } from "react-router-dom";
import { useUserContext } from "../context/UserContext";
import UserButton from "./UserButton";

export function Header() {

  const [user, _] = useUserContext();

  const buttonsLoginSignup = (<>
    <Link to="/login">
      <Button variant="default">Log in</Button>
    </Link>
    <Link to="/signup">
      <Button>Sign up</Button>
    </Link>
  </>);

  return (
    <Box pb={30}>

      <header className={classes.header}>
        <Group justify="space-between">

          <Group className={classes.logo}>
            <img alt="logo" src="/logo.svg" />
          </Group>

          <Group h="100%" gap={0}>
            <Link to="/decks" className={classes.link}>
              My Decks
            </Link>
            <Link to="/search" className={classes.link}>
              Search
            </Link>
          </Group>

          <Group>
            {user.userName == "" ? buttonsLoginSignup : <UserButton />}
          </Group>

        </Group>
      </header>
    </Box>
  );
}


export default Header;