import { Box } from '@mantine/core'
import { Link } from 'react-router-dom'

function App() {
  return (
    <>
      <Box m="auto" maw={800}>
        <h1>App</h1>
        <p><Link to="/login">Log in</Link></p>
        <p><Link to="/signup">Sign up</Link></p>
      </Box>
    </>
  )
}

export default App