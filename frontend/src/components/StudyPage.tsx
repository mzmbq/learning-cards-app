import { Button, Container, LoadingOverlay } from "@mantine/core";
import { useEffect, useState } from "react";

import classes from "./StudyPage.module.css";
import { useParams } from "react-router-dom";
import ErrorPage from "./ErrorPage";
import { Card } from "../types";
import CONFIG from "../config";

const frontText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.";
const backText = frontText;

enum Status {
  Again = 0,
  Hard,
  Good,
  Easy,
}

function StudyPage() {
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [backVisible, setbackVisible] = useState(false);
  const [front, setFront] = useState(frontText);
  const [back, setBack] = useState(backText);

  const [card, setCard] = useState<Card | undefined>(undefined);

  const deckIDStr = useParams().deck_id;

  useEffect(
    () => { fetchCard(); },
    []);

  if (error) {
    return <ErrorPage message={error} />;
  }

  const fetchCard = async () => {
    if (deckIDStr === undefined) {
      setError("deck id is undefined");
      return <></>;
    }
    const deckID = parseInt(deckIDStr);
    if (isNaN(deckID)) {
      setError("invalid deck id");
      return <></>;
    }

    try {
      setLoading(true);

      const response = await fetch(`${CONFIG.backendURL}/api/study/get-card/${deckID}`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("failed to fetch a card");
      }

      const data = await response.json();
      setCard(data.card);

    } catch (error: any) {
      console.error(error);
      setError(error.message);
    } finally {
      setLoading(false);
    }
  };


  const submitCard = async (status: Status) => {
    try {
      setLoading(true);

      const response = await fetch(`${CONFIG.backendURL}/api/study/submit/${card?.id}`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({ status: status }),
      });

      if (!response.ok) {
        throw new Error("failed to submit");
      }

    } catch (error: any) {
      console.error(error);
      setError(error.message);
    } finally {
      setLoading(false);
    }
  };


  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>Study</h2>

      <div className={classes.cardContainer}>

        <div className={classes.cardFront}><p>{card?.front}</p></div>
        <div className={classes.cardBack}>{backVisible && <p>{card?.back}</p>}</div>

        <div className={classes.buttons}>
          {!backVisible && <Button onClick={() => setbackVisible(true)}>Show Answwer</Button>}
          {backVisible &&
            <>
              <Button color="red" onClick={() => submitCard(Status.Again)}>Again</Button>
              <Button color="grey" onClick={() => submitCard(Status.Hard)}>Hard</Button>
              <Button color="green" onClick={() => submitCard(Status.Good)}>Good</Button>
              <Button color="blue" onClick={() => submitCard(Status.Easy)}>Easy</Button>
            </>
          }
        </div>
      </div>

    </Container >
  );
}

export default StudyPage;