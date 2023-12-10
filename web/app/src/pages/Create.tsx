import axios from "axios";
import { useEffect, useState } from "react";

const post = async (
  e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
  title: string,
  subtitle: string,
  type: string,
  content: string
) => {
  e.preventDefault();

  try {
    let res = await axios.post("https://nixpig.dev/api/content", {
      title,
      subtitle,
      type,
      body: content,
    });

    console.log(res.data);
  } catch (e) {
    alert("Post failed");
  }
};

export const Create = ({ id }: { id?: number }) => {
  const [title, setTitle] = useState("");
  const [subtitle, setSubtitle] = useState("");
  const [type, setType] = useState("");
  const [content, setContent] = useState("");

  useEffect(() => {
    if (id) {
      axios.get(`https://nixpig.dev/api/content/${id}`).then((res) => {
        let { title, subtitle, type, body } = res.data.data;
        setTitle(title);
        setSubtitle(subtitle);
        setType(type);
        setContent(body);
      });
    }
  });

  return (
    <div>
      <form>
        <div>
          <label htmlFor="title">Title: </label>
          <input id="title" onChange={(e) => setTitle(e.target.value)} />
        </div>

        <div>
          <label htmlFor="subtitle">Subtitle: </label>
          <input id="subtitle" onChange={(e) => setSubtitle(e.target.value)} />
        </div>

        <div>
          <label htmlFor="type">Type: </label>
          <input id="type" onChange={(e) => setType(e.target.value)} />
        </div>

        <div>
          <label htmlFor="content">Content: </label>
          <textarea id="content" onChange={(e) => setContent(e.target.value)} />
        </div>

        <div>
          <button
            type="submit"
            id="create-btn"
            onClick={(e) => post(e, title, subtitle, type, content)}
          >
            Post
          </button>
        </div>
      </form>
    </div>
  );
};
