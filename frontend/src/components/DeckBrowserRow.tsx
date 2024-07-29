import React from 'react';
import { Deck } from '../types';
import { ActionIcon, Button, Card, Group, Menu, rem, Stack, Table, Text, TextInput } from '@mantine/core';
import { IconCards, IconEdit, IconListSearch, IconSettings, IconTrash, IconWheel } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';

type Props = {
  deck: Deck;
  deckDelete: (deckId: number) => void;
  deckRename: (deckId: number, deckname: string) => void;
};

function DeckBrowserRow({ deck, deckDelete, deckRename }: Props) {
  const navigate = useNavigate();
  const [renaming, setRenaming] = React.useState(false);
  const [newName, setNewName] = React.useState(deck.name);

  return (
    // <Table.Tr
    <Card shadow="sm" padding="lg" radius="md" withBorder>
      <Card.Section>
      </Card.Section>

      <Group justify="space-between" mt="md" mb="xs">
        {renaming &&
          <TextInput
            value={newName}
            onChange={(event) => setNewName(event.currentTarget.value)}
            onKeyDown={(event) => {
              if (event.key === 'Enter') {
                deckRename(deck.id!, newName);
                setRenaming(false);
              }
            }}

          />}
        {!renaming && <Text>{deck.name}</Text>}

        <Menu>
          <Menu.Target>
            <ActionIcon variant="subtle" color="gray">
              <IconSettings style={{ width: rem(16), height: rem(16) }} />
            </ActionIcon>
          </Menu.Target>

          <Menu.Dropdown>
            <Menu.Item
              leftSection={<IconEdit />}
              onClick={() => setRenaming(!renaming)}
            >
              Rename
            </Menu.Item>
            <Menu.Item
              leftSection={<IconListSearch />}
              onClick={() => navigate(`/deck/${deck.id}`)}
            >
              Browse
            </Menu.Item>
            <Menu.Divider />
            <Menu.Item
              leftSection={<IconTrash />}
              color="red"
              onClick={() => deckDelete(deck.id!)}
            >
              Delete
            </Menu.Item>
          </Menu.Dropdown>
        </Menu>
      </Group>

      <Button
        size="sm"
        color="lime"
        leftSection={<IconCards />}
        onClick={() => navigate(`/study/${deck.id}`)}>
        Study
      </Button>

    </Card>
  );
}

export default DeckBrowserRow;