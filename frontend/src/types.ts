export type Deck = {
  id?: number;
  name: string;
  user_id: number;
};

export type Card = {
  id?: number;
  front: string;
  back: string;

  deck_id: number;
};

export type DictionaryEntry = {
  word: string;
  definition: string;
  examples?: string[];
};
