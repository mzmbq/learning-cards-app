import { Button, Container, Group, LoadingOverlay, Table, TextInput } from "@mantine/core";
import { useEffect, useState } from "react";
import CONFIG from "../config";
import { useNavigate } from "react-router-dom";

import { Deck } from "../types";
import ErrorPage from "./ErrorPage";

type DeckCreateReqBody = {
  deckName: string;
};

function DeckBrowser() {
  const [decks, setDecks] = useState<Deck[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [formInput, setFormInput] = useState("");

  const navigate = useNavigate();

  const rows = decks.map((d) => (
    <Table.Tr key={d.id}>
      <Table.Td>{d.name}</Table.Td>
      <Table.Td>
        <Group gap="xs">
          <Button disabled color="green">Study</Button>
          <Button disabled color="blue">Rename</Button>
          <Button color="blue" onClick={() => navigate(`/deck/${d.id}`)}>View</Button>
          <Button color="red" onClick={() => deckDelete(d.id!)}>Delete</Button>
        </Group>
      </Table.Td>
    </Table.Tr>
  ));

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

  if (error) {
    return <ErrorPage message={error} />;
  }

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>Decks</h2>

      <Table striped highlightOnHover withTableBorder>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>Deck</Table.Th>
            <Table.Th></Table.Th>
          </Table.Tr>
        </Table.Thead>

        <Table.Tbody>
          {rows}

          <Table.Tr>

            <Table.Td>
              <TextInput
                placeholder="New deck label"
                value={formInput}
                onChange={(event) => setFormInput(event.currentTarget.value)}
              />
            </Table.Td>
            <Table.Td>
              <Button
                type="submit"
                onClick={() => deckCreate()}>
                Create
              </Button>

            </Table.Td>
          </Table.Tr>

        </Table.Tbody>
      </Table>


    </Container>
  );
}

export default DeckBrowser;
