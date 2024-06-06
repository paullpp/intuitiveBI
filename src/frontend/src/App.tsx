import { useQuery } from "@tanstack/react-query";

export default function App() {
  const { data, error } = useQuery({
    queryKey: [ 'tables' ],
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
      <h1>Here are your tables</h1>
      <div>
        {data && data.results.map((table: string) => (
          <p>{table}</p>
        ))}
      </div>
    </>
  );
}
