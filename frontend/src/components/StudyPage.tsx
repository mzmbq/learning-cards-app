import { Box, Container, LoadingOverlay } from "@mantine/core";
import { useEffect, useState } from "react";
import CONFIG from "../config";



function StudyPage() {
  const [loading, setLoading] = useState(false);
  const [data, setData] = useState<any>(null);

  const fetchUser = async () => {
    try {

      const response = await fetch(`${CONFIG.backendURL}/api/user/whoami`, {
        method: "GET",
        credentials: "include"
      })

      if (!response.ok) {
        throw new Error("lala")
      }

      setData(await response.json())

    } catch (error: any) {
      console.log(error)
    }
  }

  useEffect(() => { fetchUser(); }, [])

  return (
    <Container>
      <LoadingOverlay visible={loading} />

      <p>Study</p>
      {JSON.stringify(data)}


    </Container>
  )
}

export default StudyPage;