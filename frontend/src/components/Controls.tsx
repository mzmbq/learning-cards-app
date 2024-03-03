import { Col, Row, Button } from "react-bootstrap";

type Props = {};

export default function Controls({}: Props) {
  const showAnswer = false;
  return (
    <Row>
      {showAnswer ? (
        <Col>
          <Button variant="primary">Show Answer</Button>
        </Col>
      ) : (
        <Col>
          <Button variant="primary">Good</Button>
          <Button variant="secondary">Ok</Button>
          <Button variant="danger">Again</Button>
        </Col>
      )}
    </Row>
  );
}
