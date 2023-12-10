import axios from "axios";
import { useEffect, useState } from "react";

export const Content = () => {
  const [content, setContent] = useState<any[]>([]);

  useEffect(() => {
    axios.get("https://nixpig.dev/api/content").then((res) => {
      setContent(res.data.data);
    });
  }, []);

  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            <th>Type</th>
          </tr>
        </thead>
        <tbody>
          {content ? (
            content.map((content) => (
              <tr key={content.id}>
                <td>{content.id}</td>
                <td>{content.title}</td>
                <td>{content.type}</td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={3}>No content</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};
