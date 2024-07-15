import { Avatar, LoadingOverlay, Menu, useMantineColorScheme } from "@mantine/core";

import classes from "./UserButton.module.css";
import { useUserContext } from "../context/UserContext";
import CONFIG from "../config";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { IconLogout, IconMoon, IconSun } from "@tabler/icons-react";


function UserButton() {
  const [loading, setLoading] = useState(false);
  const [user, setUser] = useUserContext();
  const navigate = useNavigate();
  const { colorScheme, toggleColorScheme } = useMantineColorScheme();


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
      navigate("/");

    } catch (error: any) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <LoadingOverlay visible={loading} />

      <Menu position="left-start" shadow="md" openDelay={100} closeDelay={400} trigger="click-hover">

        <Menu.Target>
          <div className={classes.userButton}>
            <Avatar src={null} alt="" />
          </div>
        </Menu.Target>

        <Menu.Dropdown>
          <Menu.Label>
            {user.userName}
          </Menu.Label>
          <Menu.Item leftSection={colorScheme === "light" ? <IconMoon /> : <IconSun />} onClick={() => toggleColorScheme()}>
            Toggle Theme
          </Menu.Item>
          <Menu.Item leftSection={<IconLogout />} onClick={() => signOut()}>
            Sign out
          </Menu.Item>
        </Menu.Dropdown>
      </Menu >
    </>
  );
}

export default UserButton;