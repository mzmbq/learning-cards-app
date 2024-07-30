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

  flashcard?: {
    ease: number;
    interval: number;
    state: number;
    step: number;
    due: Date;
  };
};

export type DictionaryEntry = {
  word: string;
  definition: string;
  examples?: string[];
};
