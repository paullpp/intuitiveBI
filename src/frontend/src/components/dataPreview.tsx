import { usePreviewStore } from "../zustand/store";

export default function Preview() {
  const data = usePreviewStore((state) => state.data);
  const name = usePreviewStore((state) => state.name);
  console.log("== Data: ", data);

  return (
    <>
      <h3>{name}</h3>
      <table className="table table-xs">
        <thead>
          <tr>
            {data && data.length > 0 && Object.keys(data[0])?.map((key) => (
              <th key={key}>
                {key}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data && data.length > 0 && data.map((row, index) => (
            <tr key={index}>
              {Object.keys(row).map((key, index) => (
                <td key={index}>
                  {row[key]}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </>
  );
}