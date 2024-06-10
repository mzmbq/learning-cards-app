import { useEffect, useState } from "react";

import CONFIG from "../config";
import { Button, Container, LoadingOverlay, Table } from "@mantine/core";
import { useNavigate, useParams } from "react-router-dom";

import classes from "./CardBrowser.module.css";

type Card = {
  id: number;
  front: string;
  back: string;
  deck_id: number;
};

function CardBrowser() {
  const [cards, setCards] = useState<Card[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  const deckID = Number(useParams().id);

  const rows = cards.map((card: Card) => (
    <Table.Tr key={card.id}>
      <Table.Td>{card.front}</Table.Td>
      <Table.Td>{card.back}</Table.Td>
    </Table.Tr>
  ));

  const fetchCards = async (deckID: number) => {
    setLoading(true);

    try {
      if (isNaN(deckID)) {
        throw new Error("invalid deck id");
      }

      const response = await fetch(`${CONFIG.backendURL}/api/deck/list/${deckID}`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("failed to fetch cards");
      }

      const data = await response.json();
      setCards(data.cards);

    } catch (error: any) {
      console.error(error);
      setError(error.message);

    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchCards(deckID);
  }, []);

  if (error) {
    return <Container><h1 style={{ color: "red" }}>Error: {error}</h1></Container>;
  }

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>Cards</h2>

      <div className={classes.outerContainer}>
        <Table striped highlightOnHover withTableBorder>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>Front</Table.Th>
              <Table.Th>Back</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>
            {rows}
          </Table.Tbody>
        </Table>

        <div className={classes.buttons}>
          <Button onClick={() => { navigate(`/new-card/${deckID}`); }}>Add</Button>
        </div>
      </div>

    </Container>
  );
}

export default CardBrowser;