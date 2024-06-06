import { useQuery } from "@tanstack/react-query";

export default function App() {
  const { data, error } = useQuery({
    queryKey: [ 'hello-world' ],
    queryFn: async () => {
      const res = await fetch('http://0.0.0.0:3000');
      console.log(res)
      return res.json();
    }
  });

  return (
    <>
      <p>
        {error?.message}
      </p>
      <h1>Hello from the client</h1>
      <h1>
        {data?.msg}
      </h1>
    </>
  );
}
