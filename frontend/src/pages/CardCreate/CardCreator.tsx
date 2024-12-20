import { Button, Container, LoadingOverlay, Textarea } from "@mantine/core";
import classes from "./CardCreator.module.css";
import { useEffect, useRef, useState } from "react";

import { Card, Deck, DictionaryEntry } from "../../types";
import CONFIG from "../../config";
import { Link, useLocation, useParams } from "react-router-dom";
import ErrorPage from "../Error/ErrorPage";
import { useHotkeys } from "@mantine/hooks";

function CardCreator() {
  const location = useLocation();
  const entry: DictionaryEntry = location.state?.entry ?? undefined;

  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [currentDeck, setCurrentDeck] = useState<Deck | null>(null);
  const [front, setFront] = useState(
    entry === undefined
      ? ""
      : entry.definition + "\n" + entry.examples?.join("\n")
  );
  const [back, setBack] = useState(entry === undefined ? "" : entry.word);

  const fronInputRef = useRef<HTMLTextAreaElement | null>(null);

  const deckID = Number(useParams().id);

  useEffect(() => {
    fetchCurrentDeck();
    focusFront();
  }, []);

  let card: Card;
  if (!error) {
    if (isNaN(deckID)) {
      setError("deck_id undefined");
    } else {
      card = {
        front: front,
        back: back,
        deck_id: deckID,
      };
    }
  }

  const focusFront = () => {
    if (fronInputRef.current) {
      fronInputRef.current.focus();
    }
  };

  const fetchCurrentDeck = async () => {
    if (isNaN(deckID)) {
      return;
    }
    setLoading(true);

    try {
      const response = await fetch(`${CONFIG.backendURL}/api/deck/list`, {
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
      const deck = decks.find((d) => d.id! === deckID);

      if (deck !== undefined) {
        setCurrentDeck(deck);
      } else {
        setError(`could not find the deck with id "${deckID}"`);
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

  useHotkeys(
    [
      [
        "ctrl+enter",
        () => {
          cardCreate(card);
          focusFront();
        },
      ],
    ],
    []
  );

  if (error) {
    return <ErrorPage message={error} />;
  }

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>Create a new Card</h2>

      <div className={classes.outerContainer}>
        Current Deck:
        <Link to={`/deck/${currentDeck?.id}`}>{currentDeck?.name}</Link>
        <Textarea
          ref={fronInputRef}
          label="Front side"
          autosize
          minRows={3}
          maxRows={15}
          value={front}
          onChange={(event) => {
            setFront(event.target.value);
          }}
        />
        <Textarea
          label="Back side"
          autosize
          minRows={3}
          maxRows={15}
          value={back}
          onChange={(event) => {
            setBack(event.target.value);
          }}
        />
        <div className={classes.buttons}>
          <Button onClick={() => cardCreate(card)}>Create</Button>
        </div>
      </div>
    </Container>
  );
}

export default CardCreator;
