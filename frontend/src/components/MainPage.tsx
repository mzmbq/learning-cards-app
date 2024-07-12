import { useNavigate } from "react-router-dom";
import { useUserContext } from "../context/UserContext";
import { useEffect } from "react";

function MainPage() {
  const navigate = useNavigate();
  const [user] = useUserContext();

  useEffect(() => {
    if (user.userName !== "") {
      navigate("/decks");
    }
  }, [navigate, user]);

  return (
    <>

    </>
  );
}

export default MainPage;