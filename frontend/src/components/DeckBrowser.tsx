import { Button, Container, Group, LoadingOverlay, Table, TextInput } from "@mantine/core";
import { useEffect, useState } from "react";
import CONFIG from "../config";

type Deck = {
  id: number;
  name: string;
  user_id: number;
};

type DeckCreateReqBody = {
  deckName: string;
};

function DeckBrowser() {
  const [decks, setDecks] = useState<Deck[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [formInput, setFormInput] = useState("");

  const rows = decks.map((d) => (
    <Table.Tr key={d.id}>
      <Table.Td>{d.name}</Table.Td>
      <Table.Td>
        <Group gap="xs">
          <Button color="green">Study</Button>
          <Button color="blue">Edit</Button>
          <Button color="red">Delete</Button>
        </Group>
      </Table.Td>
    </Table.Tr>
  ));

  const fetchDecks = async () => {
    setLoading(true);

    try {
      const response = await fetch(`${CONFIG.backendURL}/api/deck/list`, {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) {
        throw new Error("Failed to fetch decks");
      }
      const data = await response.json();
      console.log(data);
      setDecks(data.decks);

    } catch (error: any) {
      console.error(error);
      setError(error.message);

    } finally {
      setLoading(false);
    }
  };

  const createNewDeck = async () => {
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

  useEffect(() => {
    fetchDecks();
  }, []);

  if (error) {
    return <p style={{ color: "red" }}>Error: {error}</p>;
  }

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>Decks</h2>

      <Table striped highlightOnHover withTableBorder>
        <Table.Thead>
          <Table.Tr>
            <Table.Th>Table</Table.Th>
            <Table.Th>Actions</Table.Th>
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
                onClick={() => createNewDeck()}>
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
