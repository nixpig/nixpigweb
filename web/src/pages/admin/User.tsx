import { useEffect, useState } from "react";
import { api } from "../../api";
import { User as UserModel, NewUser as NewUserModel } from "../../models/User";

async function createUser(event: any, user: NewUserModel) {
  event.preventDefault();

  try {
    await api.post("/user", user);
  } catch (e: any) {
    console.error(e.message);
  }
}

const User = () => {
  const [users, setUsers] = useState<UserModel[]>([]);
  const [newUserUsername, setNewUserUsername] = useState("");
  const [newUserEmail, setNewUserEmail] = useState("");
  const [newUserPassword, setNewUserPassword] = useState("");

  useEffect(() => {
    const getUsers = async () => {
      try {
        const users = await api.get("/user");
        setUsers(users.data.data);
      } catch (e: any) {
        console.error(e.message);
      }
    };

    getUsers();
  }, []);

  return (
    <div>
      <>
        <h2>Users</h2>
        <h3>Create user</h3>
        <label htmlFor="username">Username</label>
        <input
          id="username"
          type="text"
          value={newUserUsername}
          onChange={(e) => setNewUserUsername(e.target.value)}
        />

        <label htmlFor="email">Email</label>
        <input
          id="email"
          type="text"
          value={newUserEmail}
          onChange={(e) => setNewUserEmail(e.target.value)}
        />

        <label htmlFor="password">Password</label>
        <input
          id="password"
          type="password"
          value={newUserPassword}
          onChange={(e) => setNewUserPassword(e.target.value)}
        />

        <button
          onClick={(e) =>
            createUser(e, {
              username: newUserUsername,
              password: newUserPassword,
              email: newUserEmail,
            })
          }
        >
          Create
        </button>
        <h3>Manage users</h3>
        <table>
          <tr>
            <th>ID</th>
            <th>Username</th>
            <th>Email</th>
            <th>Admin</th>
            <th>Role</th>
            <th>Profile</th>
            <th>Registered at</th>
            <th>Last login</th>
            <th>Manage</th>
          </tr>
          {users.map((user, key) => {
            return (
              <tr>
                <td>{`${JSON.stringify(user.id)}`}</td>
                <td>{`${JSON.stringify(user.username)}`}</td>
                <td>{`${JSON.stringify(user.email)}`}</td>
                <td>{`${JSON.stringify(user.is_admin)}`}</td>
                <td>{`${JSON.stringify(user.role)}`}</td>
                <td>{`${JSON.stringify(user.profile)}`}</td>
                <td>{`${JSON.stringify(user.registered_at)}`}</td>
                <td>{`${JSON.stringify(user.last_login)}`}</td>
                <td>MANAGE</td>
              </tr>
            );
          })}
        </table>
      </>
    </div>
  );
};

export default User;
