import { Container } from "@mantine/core";

type ErrorPageProps = {
  message: string,
};

function ErrorPage({ message }: ErrorPageProps) {
  return (
    <Container><h1 style={{ color: "red" }}>Error: {message}</h1></Container>
  );
}

export default ErrorPage;