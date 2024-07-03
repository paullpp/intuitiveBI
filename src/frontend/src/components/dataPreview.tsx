export default function Preview(props: { data: object[]}) {
  const { data } = props;

  return (
    <table className="table table-xs">
      <thead>
        <tr>
          {Object.keys(data[0]).map((key) => (
            <th key={key}>
              {key}
            </th>
          ))}
        </tr>
      </thead>
      <tbody>
        {data.map((row) => (
          <tr>
            {Object.keys(row).map((key) => (
              <td key={key}>
                {row[key]}
              </td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
}