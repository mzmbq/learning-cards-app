import { Box } from "@mantine/core";
import { Link, useNavigate } from "react-router-dom";
import { useUserContext } from "../context/UserContext";

function MainPage() {
  const navigate = useNavigate();
  const [user] = useUserContext();

  if (user.userName !== "") {
    navigate("/study");
  }

  return (
    <>
      <Box m="auto" maw={800}>
        <h1>App</h1>
        <p><Link to="/login">Log in</Link></p>
        <p><Link to="/signup">Sign up</Link></p>
      </Box>
    </>
  );
}

export default MainPage;