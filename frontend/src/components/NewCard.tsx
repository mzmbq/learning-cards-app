import { Button, Container, LoadingOverlay, Textarea } from "@mantine/core";
import classes from "./NewCard.module.css";
import { useEffect, useState } from "react";

import { Card, Deck } from "../types";
import CONFIG from "../config";
import { useParams } from "react-router-dom";
import ErrorPage from "./ErrorPage";

// TODO: don't rerender on each textfield change

function NewCard() {
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [currentDeck, setCurrentDeck] = useState<Deck | null>(null);
  const [front, setFront] = useState("");
  const [back, setBack] = useState("");

  const deckIDStr = useParams().id;

  useEffect(() => {
    fetchCurrentDeck();
  }, []);


  if (error) {
    return <ErrorPage message={error} />;
  }

  let card: Card;
  if (deckIDStr === undefined) {
    setError("deck_id undefined");
  } else {
    const deckID = parseInt(deckIDStr);
    card = {
      front: front,
      back: back,
      deck_id: deckID,
    };
    console.log(card);
  }

  const fetchCurrentDeck = async () => {
    setLoading(true);

    try {
      const response = await fetch(`${CONFIG.backendURL}/api/decks/list`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        if (response.status === 401) {
          throw new Error("unauthorized");
        }
        throw new Error("failed to fetch decks");
      }

      const data = await response.json();
      if (data.decks.length === 0) {
        throw new Error("create a deck first");
      }

      // set current deck
      const decks: Deck[] = data.decks;
      const deck = decks.find(d => d.id!.toString() === deckIDStr);

      if (deck !== undefined) {
        setCurrentDeck(deck);
      } else {
        setError(`invalid deck id "${deckIDStr}"`);
      }

    } catch (error: any) {
      console.error(error);
      setError(error.message);

    } finally {
      setLoading(false);
    }
  };

  const cardCreate = async (card: Card) => {
    setLoading(true);

    try {


      const response = await fetch(`${CONFIG.backendURL}/api/card/create`, {
        method: "POST",
        body: JSON.stringify({ card: card }),
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error(`failed to create a card ${front} : ${back}`);
      }

      setFront("");
      setBack("");
      await fetchCurrentDeck();

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

      <h2>Create a new Card</h2>

      <div className={classes.outerContainer}>

        <p>Current Deck: <b>{currentDeck?.name}</b></p>

        <Textarea
          label="Front side"
          autosize
          minRows={3}
          maxRows={15}
          value={front}
          onChange={event => { setFront(event.target.value); }}
        />

        <Textarea

          label="Back side"
          autosize
          minRows={3}
          maxRows={15}
          value={back}
          onChange={event => { setBack(event.target.value); }}

        />

        <div className={classes.buttons}>
          <Button onClick={() => cardCreate(card)}>Create</Button>
        </div>
      </div>

    </Container>
  );
}

export default NewCard;