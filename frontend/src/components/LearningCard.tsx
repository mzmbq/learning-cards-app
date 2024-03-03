import Card from "react-bootstrap/Card";
import { useState } from "react";

const lorem =
  "Distinctio fugiat magni deserunt ea. Atque dolorem quisquam quibusdam id vero sed pariatur est. Vel commodi porro dolore aliquid consequatur dolores sed rem. Veniam omnis sapiente sed quaerat et nihil. Sunt a nesciunt aut molestias magni animi consequatur. ";

const lorem2 =
  "corporis a veritatis incidunt nihil asperiores dolorem molestias eveniet vero ratione et nostrum possimus officiis excepturi minus illo eum enim ";

function LearningCard() {
  const [front, setFront] = useState<string>("Text Front");
  const [back, setBack] = useState<string>("Text Back");
  const [backVisible, setBackVisible] = useState<boolean>(true);

  const toggleBackVisible = () => {
    setBackVisible(!backVisible);
  };

  return (
    <Card>
      <Card.Body>
        <Card.Title>Question</Card.Title>
        <p>Front Content: {lorem}</p>
      </Card.Body>

      <Card.Footer>{backVisible && <p>Back Content: {lorem2}</p>}</Card.Footer>
    </Card>
  );
}

export default LearningCard;
