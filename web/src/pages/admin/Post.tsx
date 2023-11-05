import { useEffect, useState } from "react";
import { Post as PostModel, NewPost as NewPostModel } from "../../models/Post";
import { Category as CategoryModel } from "../../models/Category";
import { Status as StatusModel } from "../../models/Status";
import { User as UserModel } from "../../models/User";
import { api } from "../../api";

async function createNewPost(e: any, post: NewPostModel) {
  e.preventDefault();

  try {
    console.log({ post });
    await api.post("/post", post);
  } catch (e: any) {
    console.error(e.message);
  }
}

const Post = () => {
  const [loggedInUser, setLoggedInUser] = useState<UserModel>();
  const [posts, setPosts] = useState<PostModel[]>();
  const [categories, setCategories] = useState<CategoryModel[]>();
  const [statuses, setStatuses] = useState<StatusModel[]>();

  const [newPostTitle, setNewPostTitle] = useState("");
  const [newPostSubtitle, setNewPostSubtitle] = useState("");
  const [newPostBody, setNewPostBody] = useState("");
  const [newPostStatus, setNewPostStatus] = useState<number>(0);
  const [newPostCategoryId, setNewPostCategoryId] = useState<number>(0);

  useEffect(() => {
    const getLoggedInUser = async () => {
      try {
        const loggedInUser = await api.get("/user/me");
        console.log({ loggedInUser });
        setLoggedInUser(loggedInUser.data.data);
      } catch (e: any) {
        console.error(e.message);
      }
    };

    getLoggedInUser();
  }, []);

  useEffect(() => {
    const getPosts = async () => {
      try {
        const posts = await api.get("/post");
        setPosts(posts.data.data);
      } catch (e: any) {
        console.error(e.message);
      }
    };

    getPosts();
  }, []);

  useEffect(() => {
    const getCategories = async () => {
      try {
        const categories = await api.get("/category");
        setCategories(categories.data.data);
        setNewPostCategoryId(categories.data.data[0].id);
      } catch (e: any) {
        console.error(e.message);
      }
    };

    getCategories();
  }, []);

  useEffect(() => {
    const getStatuses = async () => {
      try {
        const statuses = await api.get("/status");
        setStatuses(statuses.data.data);
        setNewPostStatus(statuses.data.data[0].id);
      } catch (e: any) {
        console.error(e.message);
      }
    };

    getStatuses();
  }, []);

  return (
    <>
      <h2>Posts</h2>
      <h3>Create post</h3>
      <div>
        <input type="hidden" value="USER_ID_IN_HERE" />

        <label htmlFor="title">Title</label>
        <input
          id="title"
          type="text"
          value={newPostTitle}
          onChange={(e) => setNewPostTitle(e.target.value)}
        />

        <label htmlFor="subtitle">Subtitle</label>
        <input
          id="subtitle"
          type="text"
          value={newPostSubtitle}
          onChange={(e) => setNewPostSubtitle(e.target.value)}
        />

        <label htmlFor="body">Post content</label>
        <input
          id="body"
          type="text"
          value={newPostBody}
          onChange={(e) => setNewPostBody(e.target.value)}
        />

        <label htmlFor="status">Status</label>
        <select
          id="status"
          onChange={(e) => setNewPostStatus(parseInt(e.target.value))}
          value={newPostStatus}
        >
          {statuses &&
            statuses.map((status, id) => {
              return (
                <option key={id} value={status.id}>
                  {status.name}
                </option>
              );
            })}
        </select>

        <label htmlFor="category">Category</label>
        <select
          id="category"
          onChange={(e) => setNewPostCategoryId(parseInt(e.target.value))}
          value={newPostCategoryId}
        >
          {categories &&
            categories.map((category, id) => {
              return (
                <option key={id} value={category.id}>
                  {category.name}
                </option>
              );
            })}
        </select>

        <button
          onClick={(e) =>
            createNewPost(e, {
              user_id: loggedInUser?.id as number,
              title: newPostTitle,
              subtitle: newPostSubtitle,
              body: newPostBody,
              status: newPostStatus,
              category_id: newPostCategoryId,
            })
          }
        >
          Create
        </button>
      </div>

      <h3>Manage posts</h3>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            <th>Subtitle</th>
            <th>Created</th>
            <th>Updated</th>
            <th>Published</th>
            <th>User ID</th>
            <th>Category ID</th>
          </tr>
        </thead>
        <tbody>
          {posts &&
            posts.map((post) => {
              return (
                <tr>
                  <td>{post.id}</td>
                  <td>{post.title}</td>
                  <td>{post.subtitle}</td>
                  <td>{`${JSON.stringify(post.created_at)}`}</td>
                  <td>{`${JSON.stringify(post.updated_at)}`}</td>
                  <td>{`${JSON.stringify(post.published_at)}`}</td>
                  <td>{post.user_id}</td>
                  <td>{post.category_id}</td>
                </tr>
              );
            })}
        </tbody>
      </table>
    </>
  );
};

export default Post;
