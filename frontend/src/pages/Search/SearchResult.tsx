import { Button, List } from "@mantine/core";
import { DictionaryEntry } from "../../types";
import classes from "./SearchResult.module.css";

type Props = {
  entry: DictionaryEntry;
  onPress: () => void;
};

function SearchResult({ entry, onPress }: Props) {
  const examples = entry.examples?.map((example, index) => {
    const [beforeTerm, term, postTerm] = example.split("*");
    return (
      <List.Item key={index}>
        {beforeTerm}
        <b>{term}</b>
        {postTerm}
      </List.Item>
    );
  });

  return (
    <>
      <div className={classes.entry}>
        <p>{entry.definition}</p>
        {examples !== undefined && examples.length !== 0 && (
          <>
            <List>{examples}</List>
          </>
        )}
      </div>
      <Button className={classes.button} onClick={onPress}>
        Add to Deck
      </Button>
    </>
  );
}

export default SearchResult;
