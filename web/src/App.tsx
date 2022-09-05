import { useState } from 'react'
import { Box } from '@chakra-ui/react'
// import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <Box bg='tomato' w='100%' p={4} color='white'>
      Translators map
    </Box>
  )
}

export default App
