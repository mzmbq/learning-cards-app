import { useEffect, useState } from "react";

type Props = {};

export default function Card({}: Props) {
  const [data, setData] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch("http://localhost:8080/api/user/3");
      const result = await response.text();
      setData(result);
    };
    fetchData();
  }, []);

  return (
    <>
      <div>Front Text</div>
      <div>Back Text</div>
      <div>User: {data}</div>
    </>
  );
}
