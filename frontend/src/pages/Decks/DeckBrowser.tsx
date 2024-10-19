import {
  Button,
  Card,
  Container,
  Group,
  LoadingOverlay,
  Stack,
  TextInput,
} from "@mantine/core";
import { useEffect, useState } from "react";
import CONFIG from "../../config";

import { Deck } from "../../types";
import ErrorPage from "../Error/ErrorPage";
import { IconPencilPlus } from "@tabler/icons-react";
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
      const response = await fetch(`${CONFIG.backendURL}/api/deck/list`, {
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
        deckName: formInput,
      };

      const response = await fetch(`${CONFIG.backendURL}/api/deck/create`, {
        method: "POST",
        body: JSON.stringify(body),
        credentials: "include",
      });

      if (!response.ok) {
        const errorText = await response.text();
        throw new Error(
          `Failed to create a deck ${body.deckName}. ${errorText}`
        );
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
      const response = await fetch(
        `${CONFIG.backendURL}/api/deck/delete/${id}`,
        {
          method: "DELETE",
          credentials: "include",
        }
      );

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
      const respoonse = await fetch(
        `${CONFIG.backendURL}/api/deck/rename/${id}`,
        {
          method: "POST",
          credentials: "include",
          body: JSON.stringify({ deckname: deckname }),
        }
      );

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
    <DeckBrowserRow
      key={d.id}
      deck={d}
      deckDelete={deckDelete}
      deckRename={deckRename}
    />
  ));

  if (error) {
    return <ErrorPage message={error} />;
  }

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>My Decks</h2>

      <Group gap="md">
        {rows}

        <Card shadow="sm" padding="lg" radius="md" withBorder>
          <Card.Section></Card.Section>

          <Stack gap="xs">
            <Group justify="space-between" mt="md" mb="xs">
              <TextInput
                placeholder="New deck label"
                value={formInput}
                onChange={(event) => setFormInput(event.currentTarget.value)}
                onKeyDown={(event) => {
                  if (event.key === "Enter") {
                    deckCreate();
                  }
                }}
              />
            </Group>

            <Button leftSection={<IconPencilPlus />} onClick={deckCreate}>
              Create
            </Button>
          </Stack>
        </Card>
      </Group>
    </Container>
  );
}

export default DeckBrowser;
