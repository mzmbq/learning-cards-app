import { DictionaryEntry } from "../types";

type Props = {
  entry: DictionaryEntry;
  onPress: () => void;
};

function SearchResult({ entry, onPress }: Props) {
  return (
    <div>
      <p>{entry.word}</p>
      <p>{entry.definition}</p>
      <p>{entry.examples?.join(" | ")}</p>
    </div>
  );
}

export default SearchResult;