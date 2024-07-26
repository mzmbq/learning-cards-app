import { ActionIcon, Button, Container, Group, LoadingOverlay, Table, Text, TextInput } from "@mantine/core";
import { useEffect, useState } from "react";
import CONFIG from "../config";
import { Link, useNavigate } from "react-router-dom";

import { Deck } from "../types";
import ErrorPage from "./ErrorPage";
import { IconCards, IconCirclePlus, IconCirclePlus2, IconListSearch, IconPlus, IconTrash } from "@tabler/icons-react";
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


  useEffect(() => {
    fetchDecks();
  }, []);

  const rows = decks.map((d) => (
    <DeckBrowserRow key={d.id} deck={d} deckDelete={deckDelete} />
  ));

  if (error) {
    return <ErrorPage message={error} />;
  }

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>My Decks</h2>
      <Table highlightOnHover withRowBorders={false}>
        <Table.Tbody>

          {rows}


        </Table.Tbody>
      </Table>

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
