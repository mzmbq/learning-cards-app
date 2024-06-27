import { Box } from "@mantine/core";
import { Link, useNavigate } from "react-router-dom";
import { useUserContext } from "../context/UserContext";

function MainPage() {
  const navigate = useNavigate();
  const [user] = useUserContext();

  if (user.userName !== "") {
    navigate("/decks");
  }

  return (
    <>

    </>
  );
}

export default MainPage;