import { ActionIcon, Group } from "@mantine/core";
import { IconBrandGithub } from "@tabler/icons-react";

import classes from "./Footer.module.css";

const Footer = () => {
  return (
    <div className={classes.footer}>
      <div className={classes.inner}>
        <Group gap="xs" justify="flex-end" wrap="nowrap">
          <ActionIcon size="lg" variant="default" radius="xl">
            <a href="https://github.com/mzmbq/learning-cards-app">
              <IconBrandGithub></IconBrandGithub>
            </a>
          </ActionIcon>
        </Group>
      </div>
    </div>
  );
};

export default Footer;
