import React from 'react';
import { Deck } from '../types';
import { Button, Group, Table, Text } from '@mantine/core';
import { IconCards, IconListSearch, IconTrash } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';

type Props = {
  deck: Deck;
  deckDelete: (deckId: number) => void;
};

function DeckBrowserRow({ deck, deckDelete }: Props) {
  const navigate = useNavigate();
  const [visible, setVisible] = React.useState(false);

  return (
    <Table.Tr
      key={deck.id}
      onMouseEnter={() => setVisible(true)}
      onMouseLeave={() => setVisible(false)}>
      <Table.Td
        onClick={() => console.log("hello")}>
        {/* <Link to={`/study/${d.id}`}> */}
        <Text size="lg">
          {deck.name}
        </Text>
        {/* </Link> */}
      </Table.Td>
      <Table.Td>
        {visible &&
          <Group justify="right" gap="xs">

            <Button
              size="sm"
              color="lime"
              leftSection={<IconCards />}
              onClick={() => navigate(`/study/${deck.id}`)}>
              Study
            </Button>

            <Button
              size="sm"
              color="yellow"
              leftSection={<IconListSearch />}
              onClick={() => navigate(`/deck/${deck.id}`)}
            >
              Browse
            </ Button>


            <Button
              size="sm"
              color="red"
              leftSection={<IconTrash />}
              onClick={() => deckDelete(deck.id!)}
            >
              Delete
            </ Button>

          </Group>
        }
      </Table.Td>
    </Table.Tr >
  );
}

export default DeckBrowserRow;