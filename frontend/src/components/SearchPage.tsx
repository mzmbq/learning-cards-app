import { Autocomplete, Container } from '@mantine/core';
import { useEffect, useState } from 'react';
import ErrorPage from './ErrorPage';
import CONFIG from '../config';


type Suggestions = {
  suggestions: string[];
};

const dict = "wiktionary";

export default function SearchPage() {
  const [formInput, setFormInput] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [dropdownValues, setDropdownValues] = useState<string[]>([]);


  const fetchSuggestions = async () => {
    if (formInput === "") {
      setDropdownValues([]);
      return;
    }
    try {
      const response = await fetch(`${CONFIG.backendURL}/api/search/${dict}/${formInput}`, {
        method: "GET",
        credentials: "include",
      });

      if (!response.ok) {
        console.log(`Failed to fetch suggestions for ${formInput}`);
      }

      const suggestions: Suggestions = await response.json();

      setDropdownValues(suggestions.suggestions);

    } catch (error: any) {
      console.error(error);
      setError(error.message);
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
      <h2>Search</h2>
      <Autocomplete data={dropdownValues} value={formInput} onChange={(s) => setFormInput(s)} />

    </Container>
  );
}