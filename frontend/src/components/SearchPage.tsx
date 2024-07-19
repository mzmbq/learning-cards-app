import { Autocomplete, Button, Container, Group, LoadingOverlay, Select } from '@mantine/core';
import { useEffect, useState } from 'react';
import ErrorPage from './ErrorPage';
import CONFIG from '../config';
import { IconSearch } from '@tabler/icons-react';
import { DictionaryEntry } from '../types';
import SearchResult from './SearchResult';

type SuggestionsResponse = {
  suggestions: string[];
};

type DefineResponse = {
  definitions: DictionaryEntry[];
};

const dictionalries = [
  "wiktionary",
];

const e: DictionaryEntry = {
  word: "test",
  definition: "test",
  examples: ["test", "test"],
};

export default function SearchPage() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [formInput, setFormInput] = useState("");
  const [suggestions, setDropdownValues] = useState<string[]>([]);
  const [selectedDict, setSelectedDict] = useState<string | null>(dictionalries[0]);
  const [entries, setEntries] = useState<DictionaryEntry[]>([]);
  const [searchPerformed, setSearchPerformed] = useState(false);

  const fetchSuggestions = async () => {
    if (formInput === "") {
      setDropdownValues([]);
      return;
    }
    try {
      const response = await fetch(`${CONFIG.backendURL}/api/search/${selectedDict}/${formInput}`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        console.log(`Failed to fetch suggestions for ${formInput}`);
      }
      const sugg: SuggestionsResponse = await response.json();
      setDropdownValues(sugg.suggestions);

    } catch (error: any) {
      console.error(error);
      setError(error.message);
    }
  };

  const doSearch = async () => {
    if (formInput === "") {
      return;
    }
    try {
      setSearchPerformed(true);
      setLoading(true);
      const response = await fetch(`${CONFIG.backendURL}/api/define/${selectedDict}/${formInput}`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        throw new Error("doSearch failed");
      }

      const data: DefineResponse = await response.json();
      setEntries(data.definitions);

    } catch (error: any) {
      console.error(error);
      setError(error.message);
    } finally {
      setLoading(false);
    }
  };


  useEffect(() => {
    const timer = setTimeout(() => {
      console.log("2 sec passed");
      fetchSuggestions();
    }, 0.5 * 1000);

    return () => clearTimeout(timer);
  }, [formInput]);

  if (error) {
    return <ErrorPage message={error} />;
  }

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <h2>Search</h2>
      <Group
        gap="sm">

        <Autocomplete
          data={suggestions}
          value={formInput}
          onChange={(s) => setFormInput(s)}
          style={{ width: '50%' }}
        />
        <Select
          value={selectedDict}
          onChange={(value) => setSelectedDict(value)}
          data={dictionalries.map((d) => ({ value: d, label: d }))}
        />
        <Button
          leftSection={<IconSearch />}
          onClick={doSearch}
        >
          Search
        </Button>
      </Group>

      {searchPerformed && <p>{entries.length} Results</p>}
      {searchPerformed && !entries && <h3>No Results</h3>}
      {searchPerformed && entries.map((entry, index) => (
        <SearchResult key={index} entry={entry} onPress={() => { }} />
      ))}


    </Container>
  );
}