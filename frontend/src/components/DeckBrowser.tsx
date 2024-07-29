import { Button, Container, Group, LoadingOverlay, Table, TextInput } from "@mantine/core";
import { useEffect, useState } from "react";
import CONFIG from "../config";

import { Deck } from "../types";
import ErrorPage from "./ErrorPage";
import { IconCirclePlus, } from "@tabler/icons-react";
import DeckBrowserRow from "./DeckBrowserRow";

type DeckCreateReqBody = {
  deckName: string;
};

function DeckBrowser() {
  const [decks, setDecks] = useState<Deck[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [formInput, setFormInput] = useState("");




  const fetchDecks = async () => {
    setLoading(true);
    try {
      const response = await fetch(`${CONFIG.backendURL}/api/decks/list`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        if (response.status === 401) {
          throw new Error("Unauthorized");
        }
        throw new Error("Failed to fetch decks");

      }
      const data = await response.json();

      // Sort decks by id before setting the state
      data.decks.sort((a: Deck, b: Deck) => a.id! - b.id!);
      setDecks(data.decks);

    } catch (error: any) {
      console.error(error);
      setError(error.message);

    } finally {
      setLoading(false);
    }
  };

  const deckCreate = async () => {
    setLoading(true);
    try {
      const body: DeckCreateReqBody = {
        deckName: formInput
      };

      const response = await fetch(`${CONFIG.backendURL}/api/deck/create`, {
        method: "POST",
        body: JSON.stringify(body),
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error(`Failed to create a deck ${formInput}`);
      }
      await fetchDecks();
      setFormInput("");

    } catch (error: any) {
      console.error(error);
      setError(error.message);

    } finally {
      setLoading(false);
    }
  };

  const deckDelete = async (id: number) => {
    setLoading(true);
    try {
      const response = await fetch(`${CONFIG.backendURL}/api/deck/delete/${id}`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error(`Failed to delete a deck. id: ${id}`);
      }
      await fetchDecks();

    } catch (error: any) {
      console.error(error);
      setError(error.message);

    } finally {
      setLoading(false);
    }
  };

  const deckRename = async (id: number, deckname: string) => {
    setLoading(true);
    try {
      const respoonse = await fetch(`${CONFIG.backendURL}/api/deck/rename/${id}`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({ deckname: deckname }),
      });

      if (!respoonse.ok) {
        throw new Error(`Failed to rename a deck. id: ${id}`);
      }
      await fetchDecks();

    } catch (error: any) {
      console.error(error);
      setError(error.message);

    } finally {
      setLoading(false);
    }
  };


  useEffect(() => {
    fetchDecks();
  }, []);

  const rows = decks.map((d) => (
    <DeckBrowserRow key={d.id} deck={d} deckDelete={deckDelete} deckRename={deckRename} />
  ));

  if (error) {
    return <ErrorPage message={error} />;
  }

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>My Decks</h2>

      <Group
        gap="md"
      >
        {rows}

      </Group>



      <h3>Create a new deck</h3>
      <Group >
        <TextInput
          placeholder="New deck label"
          value={formInput}
          onChange={(event) => setFormInput(event.currentTarget.value)}
        />
        <Button
          type="submit"
          leftSection={<IconCirclePlus />}
          onClick={deckCreate}
        >
          Create
        </Button>
      </Group>


    </Container>
  );
}

export default DeckBrowser;
