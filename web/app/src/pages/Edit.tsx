import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { http } from "../services";

const edit = async (
  e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
  title: string,
  subtitle: string,
  type: string,
  content: string,
  id?: number
) => {
  e.preventDefault();

  try {
    let res = await http.patch(`/api/content/${id}`, {
      title,
      subtitle,
      type,
      id,
      body: content,
    });

    console.log(res.data);
  } catch (e) {
    alert("Post failed");
  }
};

export const Edit = () => {
  const [title, setTitle] = useState("");
  const [subtitle, setSubtitle] = useState("");
  const [type, setType] = useState("");
  const [content, setContent] = useState("");

  const { id } = useParams<{ id: string }>();

  console.log({ id });

  useEffect(() => {
    if (id) {
      console.log("getting...");
      http.get(`/api/content/${id}`).then((res) => {
        const { title, subtitle, type, body } = res.data.data;
        setTitle(title);
        setSubtitle(subtitle);
        setType(type);
        setContent(body);
      });
    }
  }, [id]);

  return (
    <div>
      <form>
        <div>
          <label htmlFor="title">Title: </label>
          <input
            id="title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
        </div>

        <div>
          <label htmlFor="subtitle">Subtitle: </label>
          <input
            id="subtitle"
            value={subtitle}
            onChange={(e) => setSubtitle(e.target.value)}
          />
        </div>

        <div>
          <label htmlFor="type">Type: </label>
          <input
            id="type"
            value={type}
            onChange={(e) => setType(e.target.value)}
          />
        </div>

        <div>
          <label htmlFor="content">Content: </label>
          <textarea
            id="content"
            value={content}
            onChange={(e) => setContent(e.target.value)}
          />
        </div>

        <div>
          <button
            type="submit"
            id="create-btn"
            onClick={(e) =>
              edit(e, title, subtitle, type, content, parseInt(`${id}`))
            }
          >
            Save changes
          </button>
        </div>
      </form>
    </div>
  );
};
