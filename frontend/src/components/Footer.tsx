import { Center, Container, useMantineColorScheme } from "@mantine/core";
import GitHubLogo from "./GitHubLogo";
import classes from "./Footer.module.css";

const Footer = () => {
  const scheme = useMantineColorScheme();
  const isDark = scheme.colorScheme === "dark";

  return (
    <div className={classes.footer}>
      <div className={classes.footerContent}>
        <GitHubLogo
          dark={isDark}
          url={"https://github.com/mzmbq/learning-cards-app"}
        />
      </div>
    </div>
  );
};

export default Footer;
