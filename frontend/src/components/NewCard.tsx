import { Button, Container, Textarea } from "@mantine/core";

import classes from "./NewCard.module.css"

function NewCard() {
  return (
    <Container>
      <h2>Create a new Card</h2>

      <div className={classes.outerContainer}>

        <Textarea
          label="Front side"
          autosize
          minRows={3}
          maxRows={15}
        />

        <Textarea
          label="Back side"
          autosize
          minRows={3}
          maxRows={15}
        />

        <div className={classes.buttons}>
          <Button>Create</Button>
        </div>
      </div>

    </Container>
  )
}

export default NewCard;