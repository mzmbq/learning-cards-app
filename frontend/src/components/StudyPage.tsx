import { Button, Container, LoadingOverlay } from "@mantine/core";
import { useState } from "react";

import classes from "./StudyPage.module.css"

const frontText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
const backText = frontText;

function StudyPage() {
  const [loading, setLoading] = useState(false);
  const [backVisible, setbackVisible] = useState(false);
  const [front, setFront] = useState(frontText);
  const [back, setBack] = useState(backText);


  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>Study</h2>

      <div className={classes.cardContainer}>

        <div className={classes.cardFront}><p>{front}</p></div>
        <div className={classes.cardBack}>{backVisible && <p>{back}</p>}</div>

        <div className={classes.buttons}>
          {!backVisible && <Button onClick={() => setbackVisible(true)}>Show Answwer</Button>}
          {backVisible &&
            <><Button color="green">Good</Button>
              <Button>Ok</Button>
              <Button color="red">Again</Button></>}
        </div>
      </div>

    </Container >
  )
}

export default StudyPage;