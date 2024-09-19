import { Container, Stack } from '@chakra-ui/react'
import TodoForm from './components/TodoForm'
import Navbar from './components/Navbar'
import TodoList from './components/TodoList'

export const BASE_URL = import.meta.env.MODE === "development" ? "http://localhost:5001/api" : "/api";

function App() {

  return (
    <>
      <Stack h="100vh">
        <Navbar />
        <Container>
          <TodoForm />
          <TodoList />
        </Container>
      </Stack>
    </>
  )
}

export default App
