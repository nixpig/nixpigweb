import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const viewContent = async (
  e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
  id: number
) => {
  e.preventDefault();

  try {
    let res = await axios.get(`https://nixpig.dev/api/content/${id}`);

    let { slug } = res.data.data;

    window.open(`https://nixpig.dev/${slug}`, "_blank");
  } catch (e) {
    alert("Failed to view content");
  }
};

const editContent = async (
  e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
  id: number
) => {
  e.preventDefault();
  console.log(id);
};

const deleteContent = async (
  e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
  id: number
) => {
  e.preventDefault();

  try {
    await axios.delete(`https://nixpig.dev/api/content/${id}`);
  } catch (e) {
    alert("Failed to delete content");
  }
};

export const Content = () => {
  const [content, setContent] = useState<any[]>([]);

  useEffect(() => {
    axios.get("https://nixpig.dev/api/content").then((res) => {
      setContent(res.data.data);
    });
  }, []);

  const navigate = useNavigate();

  return (
    <div>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Type</th>
            <th>Title</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {content ? (
            content.map((content) => (
              <tr key={content.id}>
                <td>{content.id}</td>
                <td>{content.type}</td>
                <td>{content.title}</td>
                <td>
                  <button onClick={(e) => viewContent(e, content.id)}>
                    View
                  </button>
                  |
                  <button onClick={(e) => editContent(e, content.id)}>
                    Edit
                  </button>
                  |
                  <button onClick={(e) => deleteContent(e, content.id)}>
                    Delete
                  </button>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan={4}>No content</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
};
