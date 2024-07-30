import { useEffect, useState } from "react";

import CONFIG from "../config";
import { Button, Container, LoadingOverlay, Table } from "@mantine/core";
import { useNavigate, useParams } from "react-router-dom";

import classes from "./CardBrowser.module.css";
import ErrorPage from "./ErrorPage";
import moment from "moment-timezone";

import { Card } from "../types";

const formatDate = (date: Date) => {
  const m = moment(date).tz(moment.tz.guess());

  if (m.isBefore(moment().add(5, "minute"))) {
    return "Now";
  }
  if (m.day() === moment().day()) {
    return "Today, " + m.format("hh:mm");
  }
  if (m.day() === moment().day() + 1) {
    return "Tomorrow, " + m.format("hh:mm");
  }
  if (m.year() === moment().year()) {
    return m.format("DD MMM, hh:mm");
  }
  return m.format("DD MMM YYYY, hh:mm");
};

function CardBrowser() {
  const [cards, setCards] = useState<Card[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  const deckID = Number(useParams().id);

  // Sort cards by due date
  cards.sort((a: Card, b: Card) => moment(a.flashcard!.due).diff(moment(b.flashcard!.due)));

  const rows = cards.map((card: Card) => (
    <Table.Tr key={card.id}>
      <Table.Td>{card.front}</Table.Td>
      <Table.Td>{card.back}</Table.Td>
      <Table.Td>{formatDate(card.flashcard!.due)}</Table.Td>
    </Table.Tr>
  ));

  const fetchCards = async (deckID: number) => {
    setLoading(true);

    try {
      if (isNaN(deckID)) {
        throw new Error("invalid deck id");
      }

      const response = await fetch(`${CONFIG.backendURL}/api/deck/list-cards/${deckID}`, {
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
  }, [deckID]);

  if (error) {
    return <ErrorPage message={error} />;
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
              <Table.Th>Due Date</Table.Th>
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