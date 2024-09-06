import { createTheme, MantineProvider } from "@mantine/core";
import { UserContextProvider } from "../context/UserContext";
import Pages from "../pages/Pages";

const theme = createTheme({
  spacing: {
    xs: "0.3rem",
  },
});

function App() {
  return (
    <UserContextProvider>
      <MantineProvider defaultColorScheme="auto" theme={theme}>
        <Pages />
      </MantineProvider>
    </UserContextProvider>
  );
}

export default App;
